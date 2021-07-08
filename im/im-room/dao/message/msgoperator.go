package message

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hymongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sort"
	"sync"
)

type roomMsgOperator struct {
	mCli *hymongodb.HyMongo
	dbName string
	collections map[string]interface{}
	collLock sync.RWMutex
}

func (r *roomMsgOperator) InsertRoomMsg(roomId int64, msg *RoomMsg) error {
	collName := getCollectionName(roomId)
	if _, err := r.checkAndCreateCollection(collName); err != nil {
		return err
	}
	if _, err := r.mCli.InsertOne(r.dbName, collName, getRoomMsgBson(msg)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}
	return nil
}

func (r *roomMsgOperator) InsertRoomMsgList(roomId int64, msgList []RoomMsg) error {
	collName := getCollectionName(roomId)
	if _, err := r.checkAndCreateCollection(collName); err != nil {
		return err
	}
	if _, err := r.mCli.InsertMany(r.dbName, collName, getRoomMsgListBson(msgList)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}
	return nil
}

func (r *roomMsgOperator) FindRoomMsg(roomId, userId int64, clientMsgId string) (*RoomMsg, error) {
	collName := getCollectionName(roomId)
	if _, err := r.checkAndCreateCollection(collName); err != nil {
		return nil, err
	}
	result, err := r.mCli.FindOne(r.dbName, collName, getRoomMsgKey(userId, clientMsgId), nil)
	if err != nil {
		return nil, err
	}
	return parseRoomMsg(result)
}

func (r *roomMsgOperator) FindRoomMsgListByLimit(roomId int64, baseIndex string, limit int64, direction int32) ([]RoomMsg, error) {
	collName := getCollectionName(roomId)
	if _, err := r.checkAndCreateCollection(collName); err != nil {
		return nil, err
	}
	opt := options.Find().SetLimit(limit)
	opt = opt.SetSort(bson.M{"_id": -1})
	if direction == 1 {
		opt = opt.SetSort(bson.M{"_id": 1})
	}
	cursor, err := r.mCli.Find(r.dbName, collName, getRoomMsgByLimitKey(baseIndex, direction), opt)
	if err != nil {
		return nil, err
	}
	return parseRoomMsgList(cursor)
}

func NewRoomMsgOperator(dbName string, mCli *hymongodb.HyMongo) RoomMsgDao {
	opr := roomMsgOperator{}
	opr.dbName = dbName
	opr.mCli = mCli
	opr.collections = make(map[string]interface{}, 0)
	names, _ := mCli.GetCollectionNames(dbName)
	opr.collLock.Lock()
	for _, v := range names{
		opr.collections[v] = 0
	}
	opr.collLock.Unlock()
	return &opr
}


func (r *roomMsgOperator) checkCollection(collName string) (bool, error) {
	r.collLock.RLock()
	if _, ok := r.collections[collName]; ok {
		r.collLock.RUnlock()
		return true, nil
	}
	r.collLock.RUnlock()

	has, err := r.mCli.HasCollection(r.dbName, collName)
	if err != nil {
		log.Errorf("mongodb is broken")
		return false, err
	}
	if has {
		r.addCollection(collName)
		return true, nil
	}
	return false, nil
}

func (r *roomMsgOperator) addCollection(collName string) {
	r.collLock.Lock()
	r.collections[collName] = 1
	r.collLock.Unlock()
}

func (r *roomMsgOperator) createCollectionAndIndex(collName string) error {
	flag := true
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{"fromId", 1}, {"clientMsgId", 1}},
			Options: &options.IndexOptions{Background: &flag},
		},
		{
			Keys:    bson.D{{"createTime", 1}},
			Options: &options.IndexOptions{Background: &flag},
		},
	}
	_, err := r.mCli.CreateCollectionIndex(r.dbName, collName, indexes)
	if err == nil {
		r.addCollection(collName)
	}
	return err
}

func (r *roomMsgOperator) checkAndCreateCollection(collName string) (bool, error) {
	find, err := r.checkCollection(collName)
	if err != nil {
		return false, err
	}
	if find {
		return true, nil
	}
	if err := r.createCollectionAndIndex(collName); err != nil {
		return false, err
	}
	return true, nil

}

func getCollectionName(roomId int64) string {
	return fmt.Sprintf("room_%d", roomId)
}

func parseRoomMsgList(cursor *mongo.Cursor) ([]RoomMsg, error) {
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	msgList := make([]RoomMsg, 0)
	if err := cursor.All(context.TODO(), &msgList); err != nil {
		return nil, err
	}
	sort.Slice(msgList, func(i, j int) bool {
		return msgList[i].Oid.String() < msgList[j].Oid.String()
	})
	return msgList, nil
}

func parseRoomMsg(result *mongo.SingleResult) (*RoomMsg, error) {
	if result.Err() != nil {
		return nil, result.Err()
	}
	msg := RoomMsg{}
	if err := result.Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
