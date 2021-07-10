package message

import "go.mongodb.org/mongo-driver/bson/primitive"

type GroupMsg struct {
	Oid         primitive.ObjectID `bson:"_id"`
	Command     int32              `bson:"command"`
	GroupId     int64              `bson:"groupId"`
	FromId      int64              `bson:"fromId"`
	ClientMsgId string             `bson:"clientMsgId"`
	Content     []byte             `bson:"content"`
	CreateTime  int64              `bson:"createTime"`
}

type DiffusesGroupMsg struct {
	Oid       primitive.ObjectID `bson:"_id"`
	Command   int32              `bson:"command"`
	GroupId   int64              `bson:"groupId"`
	FromId    int64              `bson:"fromId"`
	AppPulled int32              `bson:"appPulled"`
	PcPulled  int32              `bson:"pcPulled"`
}

type GroupMsgDao interface {
	InsertGroupMsg(groupId int64, msg *GroupMsg) error
	InsertGroupMsgList(groupId int64, msgList []GroupMsg) error
	FindGroupMsg(groupId, userId int64, clientMsgId string) (*GroupMsg, error)
	FindGroupMsgListByLimit(groupId int64, baseIndex string, limit int64, direction int32) ([]GroupMsg, error)

	InsertDiffusesGroupMsg(fromId int64, loginType int32, members []int64, msg *DiffusesGroupMsg, groupMsg *GroupMsg) error
	FindGroupOfflineMsg(userId int64, loginType int32, groupId int64) ([]GroupMsg, error)
}
