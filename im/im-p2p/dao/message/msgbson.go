package message

import (
	"go.mongodb.org/mongo-driver/bson"
	"hy-im/im/im-common/imbase"
)

func getP2pMsgBson(msg *P2pMsg) interface{} {
	return bson.M{"_id": msg.Oid, "command": msg.Command, "toId": msg.ToId, "fromId": msg.FromId, "clientMsgId": msg.ClientMsgId,
							"content": msg.Content, "createTime": msg.CreateTime, "isCancel": msg.IsCancel, "appPulled": msg.AppPulled, "pcPulled": msg.PcPulled}
}

func getP2pMsgListBson(msgList []P2pMsg) []interface{} {
	var bsonList []interface{}
	for _, v := range msgList {
		bsonList = append(bsonList, getP2pMsgBson(&v))
	}
	return bsonList
}

func getP2pMsgKey(userId int64, clientMsgId string) interface{} {
	return bson.M{"fromId": userId, "clientMsgId": clientMsgId}
}

func getP2pMsgByLimitKey(baseIndex string, direction int32) interface{} {
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

func getOfflineP2pMsgKey(userId int64, loginType int32) interface{} {
	filter := bson.M{}
	if loginType == imbase.LoginApp {
		filter["appPulled"] = 0
	} else {
		filter["pcPulled"] = 0
	}
	return filter
}
