package message

import "go.mongodb.org/mongo-driver/bson/primitive"

type P2pMsg struct {
	Oid         primitive.ObjectID `bson:"_id"`
	Command     int32              `bson:"command"`
	ToId        int64              `bson:"toId"`
	FromId      int64              `bson:"fromId"`
	ClientMsgId string             `bson:"clientMsgId"`
	Content     []byte             `bson:"content"`
	CreateTime  int64              `bson:"createTime"`
	IsCancel    int32              `bson:"isCancel"`
	AppPulled   int32              `bson:"appPulled"`
	PcPulled    int32              `bson:"pcPulled"`
}

type P2pMsgDao interface {
	InsertP2pMsg(userId int64, loginType int32, msg *P2pMsg) error
	InsertP2pMsgList(userId int64, msgList []P2pMsg) error
	UpdateP2pMsgPulled(userId int64, loginType int32, msgIds []primitive.ObjectID) error
	UpdateP2pMsgCancel(userId, toId int64, msgId primitive.ObjectID) error
	FindP2pMsgByClientMsgId(userId int64, clientMsgId string) (*P2pMsg, error)
	FindP2pMsg(userId int64, msgId string) (*P2pMsg, error)
	FindP2pMsgListByLimit(userId int64, baseIndex string, limit int64, direction int32) ([]P2pMsg, error)
	FindP2pOfflineMsg(userId int64, loginType int32) ([]P2pMsg, error)
}
