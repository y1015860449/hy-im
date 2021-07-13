package handler

import (
	"context"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hy_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hy-im/im/im-common/imbase"
	"hy-im/im/im-common/objid"
	appPt "hy-im/im/im-common/proto/app"
	innerPt "hy-im/im/im-common/proto/inner"
	mqPt "hy-im/im/im-common/proto/mq"
	"hy-im/im/im-common/utils"
	"hy-im/im/im-p2p/dao/cache"
	"hy-im/im/im-p2p/dao/message"
)

type Handler struct {
	CacheDao    cache.CacheDao
	GroupMsgDao message.P2pMsgDao
}

func (h *Handler) P2p(ctx context.Context, req *innerPt.P2PReq, rsp *innerPt.P2PRsp) error {
	if req.UserId <= 0 || req.Retry <= 0 || len(req.Content) <= 0 {
		packageP2pResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), req.Command)
		return nil
	}
	switch req.Command {
	case int32(appPt.ImCmd_cmd_p2p_msg):
		return h.P2pMsgHandler(req, rsp)
	case int32(appPt.ImCmd_cmd_p2p_msg_deliver_ack):
		return h.P2pMsgDeliverAckHandler(req, rsp)
	default:
		packageP2pResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), req.Command)
	}
	return nil
}

func (h *Handler) P2pMsgHandler(req *innerPt.P2PReq, rsp *innerPt.P2PRsp) error {
	var msg appPt.P2PMsg
	if err := proto.Unmarshal(req.Content, &msg); err != nil {
		packageP2pResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), int32(appPt.ImCmd_cmd_p2p_msg_ack))
		return err
	}

	if msg.FromId != req.UserId {
		packageP2pResponse(req, rsp, msg.ClientMsgId, "", int32(innerPt.SrvErr_srv_err_param), int32(appPt.ImErrCode_err_param_except), int32(appPt.ImCmd_cmd_p2p_msg_ack))
	}

	// todo 好友状态判断

	serverId := ""
	if req.Retry == 1 {
		if rest, err := h.GroupMsgDao.FindP2pMsg(msg.FromId, msg.ClientMsgId); err != nil {
			log.Errorf("find p2p msg err(%+v)", err)
			packageP2pResponse(req, rsp, msg.ClientMsgId, "", int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), int32(appPt.ImCmd_cmd_p2p_msg_ack))
			return err
		} else {
			serverId = rest.Oid.String()
		}
	}
	// 重传消息不二次保存，直接推送
	var data []byte
	if len(serverId) <= 0 {
		chatId := utils.GetChatId(msg.FromId, msg.ToId)
		serverId = objid.GetObjectId(imbase.SessionTypeP2p, chatId)
		oid, _ := primitive.ObjectIDFromHex(serverId)
		msg.MsgId = serverId
		msg.MsgTime = hy_utils.GetMillisecond()
		data, _ = proto.Marshal(&msg)
		p2pMsg := &message.P2pMsg{
			Oid:         oid,
			Command:     int32(appPt.ImCmd_cmd_p2p_msg_deliver),
			ToId:        msg.ToId,
			FromId:      msg.FromId,
			ClientMsgId: msg.ClientMsgId,
			Content:     data,
			CreateTime:  msg.MsgTime,
		}

		if err := h.GroupMsgDao.InsertP2pMsg(req.UserId, req.LoginType, p2pMsg); err != nil {
			log.Errorf("insert p2p msg err(%+v)", err)
			packageP2pResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_mongo), int32(appPt.ImErrCode_err_server_except), int32(appPt.ImCmd_cmd_p2p_msg_ack))
			return err
		}
	}
	// todo 在线转发, 不同的类型发送到不同的topic
	pushMsg := &mqPt.PushP2PMsg{
		Command: int32(appPt.ImCmd_cmd_group_msg_deliver),
		ToId: msg.ToId,
		Content: data,
		UserId:  msg.FromId,
	}
	_ = pushMsg

	packageP2pResponse(req, rsp, msg.ClientMsgId, serverId, int32(innerPt.SrvErr_srv_err_success), int32(appPt.ImErrCode_err_success), int32(appPt.ImCmd_cmd_p2p_msg_ack))
	return nil
}

func (h *Handler) P2pMsgDeliverAckHandler(req *innerPt.P2PReq, rsp *innerPt.P2PRsp) error {
	//var msg appPt.RoomDeliverAck
	//if err := proto.Unmarshal(req.Content, &msg); err != nil {
	//	return err
	//}
	packageP2pResponse(req, rsp, "", "", int32(innerPt.SrvErr_srv_err_success), int32(appPt.ImErrCode_err_success), 0)
	return nil
}

func packageP2pResponse(req *innerPt.P2PReq, rsp *innerPt.P2PRsp, msgId, serverId string, svcCode, appCode, command int32) {
	rsp.Command = command
	rsp.UserId = req.UserId
	rsp.LoginType = req.LoginType
	rsp.SvcErr = svcCode
	if command != 0 {
		roomAck := &appPt.P2PAck{
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
