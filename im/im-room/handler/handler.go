package handler

import (
	"context"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hy_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hy-im/im/im-common/base"
	"hy-im/im/im-common/objid"
	appPt "hy-im/im/im-common/proto/app"
	innerPt "hy-im/im/im-common/proto/inner"
	mqPt "hy-im/im/im-common/proto/mq"
	"hy-im/im/im-room/dao/cache"
	"hy-im/im/im-room/dao/message"
)

type Handler struct {
	CacheDao cache.CacheDao
	MsgDao   message.RoomMsgDao
}

func (h *Handler) Room(ctx context.Context, req *innerPt.RoomReq, rsp *innerPt.RoomRsp) error {
	if req.UserId <= 0 || req.Retry <= 0 || len(req.Content) <= 0 {
		packageRoomResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), req.Command)
		return nil
	}
	switch req.Command {
	case int32(appPt.ImCmd_cmd_room_msg):
		return h.RoomMsgHandler(req, rsp)
	case int32(appPt.ImCmd_cmd_room_open),
		int32(appPt.ImCmd_cmd_room_join),
		int32(appPt.ImCmd_cmd_room_quit),
		int32(appPt.ImCmd_cmd_room_remove),
		int32(appPt.ImCmd_cmd_room_close):
		return h.RoomOperatorHandler(req, rsp)
	case int32(appPt.ImCmd_cmd_room_msg_deliver_ack),
		int32(appPt.ImCmd_cmd_room_open_deliver_ack),
		int32(appPt.ImCmd_cmd_room_join_deliver_ack),
		int32(appPt.ImCmd_cmd_room_quit_deliver_ack),
		int32(appPt.ImCmd_cmd_room_remove_deliver_ack),
		int32(appPt.ImCmd_cmd_room_close_deliver_ack):
		return h.RoomMsgDeliverAckHandler(req, rsp)
	default:
		packageRoomResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), req.Command)
	}
	return nil
}

func (h *Handler) RoomMsgHandler(req *innerPt.RoomReq, rsp *innerPt.RoomRsp) error {
	// 游客身份不能发送消息
	if req.RoleType == base.RoleVisitor {
		packageRoomResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_visitor), int32(appPt.ImErrCode_err_user_visitor), int32(appPt.ImCmd_cmd_room_msg_ack))
		return nil
	}

	var msg appPt.RoomMsg
	if err := proto.Unmarshal(req.Content, &msg); err != nil {
		packageRoomResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), int32(appPt.ImCmd_cmd_room_msg_ack))
		return err
	}
	if msg.FromId != req.UserId || msg.RoomId != req.RoomId {
		packageRoomResponse(req, rsp, msg.ClientMsgId, "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), int32(appPt.ImCmd_cmd_room_msg_ack))
	}

	serverId := ""
	if req.Retry == 1 {
		if rest, err := h.MsgDao.FindRoomMsg(msg.RoomId, msg.FromId, msg.ClientMsgId); err != nil {
			log.Errorf("find room msg err(%+v)", err)
			packageRoomResponse(req, rsp, msg.ClientMsgId, "", int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), int32(appPt.ImCmd_cmd_room_msg_ack))
			return err
		} else {
			serverId = rest.Oid.String()
		}
	}
	var data []byte
	if len(serverId) <= 0 {
		serverId = objid.GetObjectId(base.SessionTypeRoom, req.RoomId)
		oid, _ := primitive.ObjectIDFromHex(serverId)
		msg.ServerMsgId = serverId
		msg.Timestamp = hy_utils.GetMillisecond()
		data, _ = proto.Marshal(&msg)
		if err := h.MsgDao.InsertRoomMsg(msg.RoomId, &message.RoomMsg{
			Oid:         oid,
			Command:     int32(appPt.ImCmd_cmd_room_msg_deliver),
			RoomId:      msg.RoomId,
			FromId:      msg.FromId,
			ClientMsgId: msg.ClientMsgId,
			Content:     data,
			CreateTime:  msg.Timestamp,
		}); err != nil {
			log.Errorf("insert room msg err(%+v)", err)
			packageRoomResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), int32(appPt.ImCmd_cmd_room_msg_ack))
			return err
		}
	}
	// todo 在线转发
	pushMsg := &mqPt.PushRoomMsg{
		Command: int32(appPt.ImCmd_cmd_room_msg_deliver),
		RoomId:  msg.RoomId,
		Content: data,
		UserId:  msg.FromId,
	}
	_ = pushMsg

	packageRoomResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_success), int32(appPt.ImErrCode_err_success), int32(appPt.ImCmd_cmd_room_msg_ack))
	return nil
}

func (h *Handler) RoomOperatorHandler(req *innerPt.RoomReq, rsp *innerPt.RoomRsp) error {
	cmdAck := int32(appPt.ImCmd_cmd_room_open_ack)
	cmdDeliver := int32(appPt.ImCmd_cmd_room_open_deliver)
	if req.Command == int32(appPt.ImCmd_cmd_room_join) {
		cmdAck = int32(appPt.ImCmd_cmd_room_join_ack)
		cmdDeliver = int32(appPt.ImCmd_cmd_room_join_deliver)
	} else if req.Command == int32(appPt.ImCmd_cmd_room_quit) {
		cmdAck = int32(appPt.ImCmd_cmd_room_quit_ack)
		cmdDeliver = int32(appPt.ImCmd_cmd_room_quit_deliver)
	} else if req.Command == int32(appPt.ImCmd_cmd_room_remove) {
		cmdAck = int32(appPt.ImCmd_cmd_room_remove_ack)
		cmdDeliver = int32(appPt.ImCmd_cmd_room_remove_deliver)
	} else if req.Command == int32(appPt.ImCmd_cmd_room_close) {
		cmdAck = int32(appPt.ImCmd_cmd_room_close_ack)
		cmdDeliver = int32(appPt.ImCmd_cmd_room_close_deliver)
	}

	var msg appPt.RoomOperator
	if err := proto.Unmarshal(req.Content, &msg); err != nil {
		packageRoomResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), cmdAck)
		return err
	}
	if req.Command == int32(appPt.ImCmd_cmd_room_open) || req.Command == int32(appPt.ImCmd_cmd_room_join) {
		// todo 判断房间是否合法

	}
	serverId := objid.GetObjectId(base.SessionTypeRoom, req.RoomId)
	oid, _ := primitive.ObjectIDFromHex(serverId)
	msg.ServerMsgId = serverId
	msg.Timestamp = hy_utils.GetMillisecond()
	data, _ := proto.Marshal(&msg)
	if err := h.MsgDao.InsertRoomMsg(msg.RoomId, &message.RoomMsg{
		Oid:         oid,
		Command:     cmdDeliver,
		RoomId:      msg.RoomId,
		FromId:      msg.UserId,
		ClientMsgId: msg.ClientMsgId,
		Content:     data,
		CreateTime:  msg.Timestamp,
	}); err != nil {
		log.Errorf("insert room msg err(%+v)", err)
		packageRoomResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), cmdAck)
		return err
	}
	// todo 在线转发
	pushMsg := &mqPt.PushRoomMsg{
		Command: int32(appPt.ImCmd_cmd_room_msg_deliver),
		RoomId:  msg.RoomId,
		Content: data,
		UserId:  msg.UserId,
		OtherId: msg.OtherId,
	}
	_ = pushMsg

	packageRoomResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_success), int32(appPt.ImErrCode_err_success), cmdAck)
	return nil
}

func (h *Handler) RoomMsgDeliverAckHandler(req *innerPt.RoomReq, rsp *innerPt.RoomRsp) error {
	//var msg appPt.RoomDeliverAck
	//if err := proto.Unmarshal(req.Content, &msg); err != nil {
	//	return err
	//}
	// 不处理，直接成功，此消息客户端可以不发送
	packageRoomResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_success), int32(appPt.ImErrCode_err_success), 0)
	return nil
}

func packageRoomResponse(req *innerPt.RoomReq, rsp *innerPt.RoomRsp, msgId, serverId string, svcCode, appCode, command int32) {
	rsp.Command = command
	rsp.UserId = req.UserId
	rsp.LoginType = req.LoginType
	rsp.RoleType = req.RoleType
	rsp.SvcErr = svcCode
	if command != 0 {
		roomAck := &appPt.RoomAck{
			RoomId:      req.RoomId,
			UserId:      req.UserId,
			ClientMsgId: msgId,
			ServerMsgId: serverId,
			Timestamp:   hy_utils.GetMillisecond(),
			ErrCode:     appCode,
		}
		content, _ := proto.Marshal(roomAck)
		rsp.Content = content
	}
}
