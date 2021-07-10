package handler

import (
	"context"
	"github.com/common/base"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hy_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hy-im/im/im-common/imbase"
	"hy-im/im/im-common/objid"
	appPt "hy-im/im/im-common/proto/app"
	innerPt "hy-im/im/im-common/proto/inner"
	mqPt "hy-im/im/im-common/proto/mq"
	"hy-im/im/im-group/dao/cache"
	"hy-im/im/im-group/dao/message"
)

type Handler struct {
	CacheDao    cache.CacheDao
	GroupMsgDao message.GroupMsgDao
}

func (h *Handler) Group(ctx context.Context, req *innerPt.GroupReq, rsp *innerPt.GroupRsp) error {
	if req.UserId <= 0 || req.Retry <= 0 || len(req.Content) <= 0 {
		packageGroupResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), req.Command)
		return nil
	}
	switch req.Command {
	case int32(appPt.ImCmd_cmd_group_msg):
		return h.GroupMsgHandler(req, rsp)
	case int32(appPt.ImCmd_cmd_group_open),
		int32(appPt.ImCmd_cmd_group_join),
		int32(appPt.ImCmd_cmd_group_quit),
		int32(appPt.ImCmd_cmd_group_remove),
		int32(appPt.ImCmd_cmd_group_close):
		return h.GroupOperatorHandler(req, rsp)
	case int32(appPt.ImCmd_cmd_group_msg_deliver_ack),
		int32(appPt.ImCmd_cmd_group_open_deliver_ack),
		int32(appPt.ImCmd_cmd_group_join_deliver_ack),
		int32(appPt.ImCmd_cmd_group_quit_deliver_ack),
		int32(appPt.ImCmd_cmd_group_remove_deliver_ack),
		int32(appPt.ImCmd_cmd_group_close_deliver_ack):
		return h.GroupMsgDeliverAckHandler(req, rsp)
	default:
		packageGroupResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), req.Command)
	}
	return nil
}

func (h *Handler) GroupMsgHandler(req *innerPt.GroupReq, rsp *innerPt.GroupRsp) error {
	// 游客身份不能发送消息
	if req.RoleType == imbase.RoleVisitor {
		packageGroupResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_visitor), int32(appPt.ImErrCode_err_user_visitor), int32(appPt.ImCmd_cmd_group_msg_ack))
		return nil
	}

	var msg appPt.GroupMsg
	if err := proto.Unmarshal(req.Content, &msg); err != nil {
		packageGroupResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), int32(appPt.ImCmd_cmd_group_msg_ack))
		return err
	}

	if msg.GroupType == base.GroupTypeNormal || msg.GroupType == base.GroupTypeDiscussion {
		if msg.FromId != req.UserId {
			packageGroupResponse(req, rsp, msg.ClientMsgId, "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), int32(appPt.ImCmd_cmd_group_msg_ack))
		}
	} else {
		// 聊天室喝直播室
		if msg.FromId != req.UserId || msg.GroupId != req.GroupId {
			packageGroupResponse(req, rsp, msg.ClientMsgId, "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), int32(appPt.ImCmd_cmd_group_msg_ack))
		}
	}

	var members[]int64
	if msg.GroupType == base.GroupTypeNormal || msg.GroupType == base.GroupTypeDiscussion {
		// todo 获取群成员列表喝群成员状态信息
	}

	serverId := ""
	if req.Retry == 1 {
		if rest, err := h.GroupMsgDao.FindGroupMsg(msg.GroupId, msg.FromId, msg.ClientMsgId); err != nil {
			log.Errorf("find group msg err(%+v)", err)
			packageGroupResponse(req, rsp, msg.ClientMsgId, "", int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), int32(appPt.ImCmd_cmd_group_msg_ack))
			return err
		} else {
			serverId = rest.Oid.String()
		}
	}
	// 重传消息不二次保存，直接推送
	var data []byte
	if len(serverId) <= 0 {
		serverId = objid.GetObjectId(imbase.SessionTypeGroup, req.GroupId)
		oid, _ := primitive.ObjectIDFromHex(serverId)
		msg.MsgId = serverId
		msg.MsgTime = hy_utils.GetMillisecond()
		data, _ = proto.Marshal(&msg)
		gMsg := &message.GroupMsg{
			Oid:         oid,
			Command:     int32(appPt.ImCmd_cmd_group_msg_deliver),
			GroupId:     msg.GroupId,
			FromId:      msg.FromId,
			ClientMsgId: msg.ClientMsgId,
			Content:     data,
			CreateTime:  msg.MsgTime,
		}
		if msg.GroupType == base.GroupTypeNormal || msg.GroupType == base.GroupTypeDiscussion {
			dMsg := &message.DiffusesGroupMsg{
				Oid:       oid,
				Command:   int32(appPt.ImCmd_cmd_group_msg_deliver),
				GroupId:   msg.GroupId,
				FromId:    msg.FromId,
			}
			if err := h.GroupMsgDao.InsertDiffusesGroupMsg(msg.FromId, req.LoginType, members, dMsg, gMsg); err != nil {
				log.Errorf("insert group msg err(%+v)", err)
				packageGroupResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), int32(appPt.ImCmd_cmd_room_msg_ack))
				return err
			}
		} else {		// 聊天室、直播
			if err := h.GroupMsgDao.InsertGroupMsg(msg.GroupId, gMsg); err != nil {
				log.Errorf("insert group msg err(%+v)", err)
				packageGroupResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), int32(appPt.ImCmd_cmd_room_msg_ack))
				return err
			}
		}
	}
	// todo 在线转发, 不同的类型发送到不同的topic
	pushMsg := &mqPt.PushGroupMsg{
		Command: int32(appPt.ImCmd_cmd_group_msg_deliver),
		GroupId: msg.GroupId,
		Content: data,
		UserId:  msg.FromId,
	}
	_ = pushMsg

	packageGroupResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_success), int32(appPt.ImErrCode_err_success), int32(appPt.ImCmd_cmd_room_msg_ack))
	return nil
}

func (h *Handler) GroupOperatorHandler(req *innerPt.GroupReq, rsp *innerPt.GroupRsp) error {
	cmdAck := int32(appPt.ImCmd_cmd_group_open_ack)
	cmdDeliver := int32(appPt.ImCmd_cmd_group_open_deliver)
	if req.Command == int32(appPt.ImCmd_cmd_group_join) {
		cmdAck = int32(appPt.ImCmd_cmd_group_join_ack)
		cmdDeliver = int32(appPt.ImCmd_cmd_group_join_deliver)
	} else if req.Command == int32(appPt.ImCmd_cmd_group_quit) {
		cmdAck = int32(appPt.ImCmd_cmd_group_quit_ack)
		cmdDeliver = int32(appPt.ImCmd_cmd_group_quit_deliver)
	} else if req.Command == int32(appPt.ImCmd_cmd_group_remove) {
		cmdAck = int32(appPt.ImCmd_cmd_group_remove_ack)
		cmdDeliver = int32(appPt.ImCmd_cmd_group_remove_deliver)
	} else if req.Command == int32(appPt.ImCmd_cmd_group_close) {
		cmdAck = int32(appPt.ImCmd_cmd_group_close_ack)
		cmdDeliver = int32(appPt.ImCmd_cmd_group_close_deliver)
	}

	var msg appPt.GroupOperator
	if err := proto.Unmarshal(req.Content, &msg); err != nil {
		packageGroupResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), cmdAck)
		return err
	}
	if req.Command == int32(appPt.ImCmd_cmd_group_open) || req.Command == int32(appPt.ImCmd_cmd_room_join) {
		// todo 判断房间是否合法

	}
	serverId := objid.GetObjectId(imbase.SessionTypeRoom, req.GroupId)
	oid, _ := primitive.ObjectIDFromHex(serverId)
	msg.MsgId = serverId
	msg.MsgTime = hy_utils.GetMillisecond()
	data, _ := proto.Marshal(&msg)
	if err := h.GroupMsgDao.InsertGroupMsg(msg.GroupId, &message.GroupMsg{
		Oid:         oid,
		Command:     cmdDeliver,
		GroupId:     msg.GroupId,
		FromId:      msg.UserId,
		ClientMsgId: msg.ClientMsgId,
		Content:     data,
		CreateTime:  msg.MsgTime,
	}); err != nil {
		log.Errorf("insert group msg err(%+v)", err)
		packageGroupResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), cmdAck)
		return err
	}
	// todo 在线转发
	pushMsg := &mqPt.PushGroupMsg{
		Command: int32(appPt.ImCmd_cmd_group_msg_deliver),
		GroupId: msg.GroupId,
		Content: data,
		UserId:  msg.UserId,
		OtherId: msg.OtherId,
	}
	_ = pushMsg

	packageGroupResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_success), int32(appPt.ImErrCode_err_success), cmdAck)
	return nil
}

func (h *Handler) GroupMsgDeliverAckHandler(req *innerPt.GroupReq, rsp *innerPt.GroupRsp) error {
	//var msg appPt.RoomDeliverAck
	//if err := proto.Unmarshal(req.Content, &msg); err != nil {
	//	return err
	//}
	// 不处理，直接成功，此消息客户端可以不发送
	packageGroupResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_success), int32(appPt.ImErrCode_err_success), 0)
	return nil
}

func packageGroupResponse(req *innerPt.GroupReq, rsp *innerPt.GroupRsp, msgId, serverId string, svcCode, appCode, command int32) {
	rsp.Command = command
	rsp.UserId = req.UserId
	rsp.LoginType = req.LoginType
	rsp.RoleType = req.RoleType
	rsp.SvcErr = svcCode
	if command != 0 {
		roomAck := &appPt.GroupAck{
			GroupId:     req.GroupId,
			UserId:      req.UserId,
			ClientMsgId: msgId,
			MsgId:       serverId,
			MsgTime:     hy_utils.GetMillisecond(),
			ErrCode:     appCode,
		}
		content, _ := proto.Marshal(roomAck)
		rsp.Content = content
	}
}
