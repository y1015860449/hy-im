package message

import (
	"go.mongodb.org/mongo-driver/bson"
)

func getRoomMsgBson(msg *RoomMsg) interface{} {
	return bson.M{"_id": msg.Oid, "command": msg.Command, "roomId": msg.RoomId,
		"fromId": msg.FromId, "clientMsgId": msg.ClientMsgId, "content": msg.Content, "createTime": msg.CreateTime}
}

func getRoomMsgListBson(msgList []RoomMsg) []interface{} {
	var bsonList []interface{}
	for _, v := range msgList {
		bsonList = append(bsonList, getRoomMsgBson(&v))
	}
	return bsonList
}

func getRoomMsgKey(userId int64, clientMsgId string) interface{} {
	return bson.M{"fromId": userId, "clientMsgId": clientMsgId}
}

func getRoomMsgByLimitKey(baseIndex string, direction int32) interface{} {
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
