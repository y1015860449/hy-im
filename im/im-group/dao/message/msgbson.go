package message

import (
	"go.mongodb.org/mongo-driver/bson"
	"hy-im/im/im-common/imbase"
)

func getGroupMsgBson(msg *GroupMsg) interface{} {
	return bson.M{"_id": msg.Oid, "command": msg.Command, "groupId": msg.GroupId,
		"fromId": msg.FromId, "clientMsgId": msg.ClientMsgId, "content": msg.Content, "createTime": msg.CreateTime}
}

func getGroupMsgListBson(msgList []GroupMsg) []interface{} {
	var bsonList []interface{}
	for _, v := range msgList {
		bsonList = append(bsonList, getGroupMsgBson(&v))
	}
	return bsonList
}

func getGroupMsgKey(userId int64, clientMsgId string) interface{} {
	return bson.M{"fromId": userId, "clientMsgId": clientMsgId}
}

func getGroupMsgByLimitKey(baseIndex string, direction int32) interface{} {
	filter := bson.M{}
	if len(baseIndex) > 0 {
		if direction == 0 {
			filter["_id"] = bson.M{"$lt": baseIndex}
		} else {
			filter["_id"] = bson.M{"$gt": baseIndex}
		}
	}
	return filter
}

func getDiffusesMsgBson(msg *DiffusesGroupMsg) interface{} {
	return bson.M{"_id": msg.Oid, "command": msg.Command, "roomId": msg.GroupId,
		"fromId": msg.FromId, "appPulled": msg.AppPulled, "pcPulled": msg.PcPulled}
}

func getOfflineGroupMsgKey(userId int64, loginType int32, groupId int64) interface{} {
	filter := bson.M{"groupId": groupId}
	if loginType == imbase.LoginApp {
		filter["appPulled"] = 0
	} else {
		filter["pcPulled"] = 0
	}
	return filter
}
