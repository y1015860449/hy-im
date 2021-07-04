package message

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomMsg struct {
	Oid         primitive.ObjectID `bson:"_id"`
	Command     int32              `bson:"command"`
	RoomId      int64              `bson:"roomId"`
	FromId      int64              `bson:"fromId"`
	ClientMsgId string             `bson:"clientMsgId"`
	Content     []byte             `bson:"content"`
	CreateTime  int64              `bson:"createTime"`
}

type RoomMsgDao interface {
	InsertRoomMsg(roomId int64, msg *RoomMsg) error
	InsertRoomMsgList(roomId int64, msgList []RoomMsg) error
	FindRoomMsg(roomId, userId int64, clientMsgId string) (*RoomMsg, error)
	FindRoomMsgListByLimit(roomId int64, baseIndex string, limit int64, direction int32) ([]RoomMsg, error)
}
