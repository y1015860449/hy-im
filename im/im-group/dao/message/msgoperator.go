package message

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hymongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hy-im/im/im-common/imbase"
	"sort"
	"sync"
)

type roomMsgOperator struct {
	mCli *hymongodb.HyMongo
	dbName string
	collections map[string]interface{}
	collLock sync.RWMutex
}

func (r *roomMsgOperator) InsertGroupMsg(groupId int64, msg *GroupMsg) error {
	collName := getGroupCollName(groupId)
	if _, err := r.checkAndCreateGroupColl(collName); err != nil {
		return err
	}
	if _, err := r.mCli.InsertOne(r.dbName, collName, getGroupMsgBson(msg)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}
	return nil
}

func (r *roomMsgOperator) InsertGroupMsgList(groupId int64, msgList []GroupMsg) error {
	collName := getGroupCollName(groupId)
	if _, err := r.checkAndCreateGroupColl(collName); err != nil {
		return err
	}
	if _, err := r.mCli.InsertMany(r.dbName, collName, getGroupMsgListBson(msgList)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}
	return nil
}

func (r *roomMsgOperator) FindGroupMsg(groupId, userId int64, clientMsgId string) (*GroupMsg, error) {
	collName := getGroupCollName(groupId)
	if _, err := r.checkAndCreateGroupColl(collName); err != nil {
		return nil, err
	}
	result, err := r.mCli.FindOne(r.dbName, collName, getGroupMsgKey(userId, clientMsgId), nil)
	if err != nil {
		return nil, err
	}
	return parseGroupMsg(result)
}

func (r *roomMsgOperator) FindGroupMsgListByLimit(groupId int64, baseIndex string, limit int64, direction int32) ([]GroupMsg, error) {
	collName := getGroupCollName(groupId)
	if _, err := r.checkAndCreateGroupColl(collName); err != nil {
		return nil, err
	}
	opt := options.Find().SetLimit(limit)
	opt = opt.SetSort(bson.M{"_id": -1})
	if direction == 1 {
		opt = opt.SetSort(bson.M{"_id": 1})
	}
	cursor, err := r.mCli.Find(r.dbName, collName, getGroupMsgByLimitKey(baseIndex, direction), opt)
	if err != nil {
		return nil, err
	}
	return parseGroupMsgList(cursor)
}

// todo 这个方法使用事务才比较合理
func (r *roomMsgOperator) InsertDiffusesGroupMsg(fromId int64, loginType int32, members []int64, msg *DiffusesGroupMsg, groupMsg *GroupMsg) error {
	if err := r.InsertGroupMsg(groupMsg.GroupId, groupMsg); err != nil {
		return err
	}
	dBson := getDiffusesMsgBson(msg)
	for _, v := range members{
		collName := getDiffusesCollName(v)
		if _, err := r.checkAndCreateDiffusesColl(collName); err != nil {
			return err
		}
		if v == fromId {
			if loginType == imbase.LoginApp {
				msg.AppPulled = 1
			} else {
				msg.PcPulled = 1
			}
			if _, err := r.mCli.InsertOne(r.dbName, collName, getDiffusesMsgBson(msg)); err != nil && err.Error() != "ErrorDuplicateKey" {
				return err
			}
		} else {
			if _, err := r.mCli.InsertOne(r.dbName, collName, dBson); err != nil && err.Error() != "ErrorDuplicateKey" {
				return err
			}
		}
	}
	return nil
}

func (r *roomMsgOperator) FindGroupOfflineMsg(userId int64, loginType int32, groupId int64) ([]GroupMsg, error) {
	// 查询扩散数据
	dCollName := getDiffusesCollName(userId)
	if _, err := r.checkAndCreateDiffusesColl(dCollName); err != nil {
		return nil, err
	}
	result, err := r.mCli.Find(r.dbName, dCollName, getOfflineGroupMsgKey(userId, loginType, groupId), nil)
	if err != nil {
		return nil, err
	}
	dMsgList, err := parseDiffusesMsgList(result)
	if err != nil {
		return nil, err
	}

	// 查询原始数据
	gCollName := getGroupCollName(groupId)
	if _, err := r.checkAndCreateGroupColl(gCollName); err != nil {
		return nil, err
	}
	var ids []primitive.ObjectID
	for _, v := range dMsgList {
		ids = append(ids, v.Oid)
	}
	filter := bson.M{"_oid": bson.M{"$in": ids}}
	rst, err := r.mCli.Find(r.dbName, gCollName, filter, nil)
	if err != nil {
		return nil, err
	}
	return parseGroupMsgList(rst)
}

func NewRoomMsgOperator(dbName string, mCli *hymongodb.HyMongo) GroupMsgDao {
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

func (r *roomMsgOperator) createGroupCollAndIndex(collName string) error {
	mIndex := func() []mongo.IndexModel {
		flag := true
		return []mongo.IndexModel{
			{
				Keys:    bson.D{{"fromId", 1}, {"clientMsgId", 1}},
				Options: &options.IndexOptions{Background: &flag},
			},
			{
				Keys:    bson.D{{"createTime", 1}},
				Options: &options.IndexOptions{Background: &flag},
			},
		}
	}
	return r.createCollAndIndex(collName, mIndex)
}

func (r *roomMsgOperator) createDiffusesCollAndIndex(collName string) error {
	mIndex := func() []mongo.IndexModel {
		flag := true
		return []mongo.IndexModel{
			{
				Keys:    bson.D{{"groupId", 1}},
				Options: &options.IndexOptions{Background: &flag},
			},
		}
	}
	return r.createCollAndIndex(collName, mIndex)
}

func (r *roomMsgOperator) createCollAndIndex(collName string, mIndex func() []mongo.IndexModel) error {
	indexes := mIndex()
	_, err := r.mCli.CreateCollectionIndex(r.dbName, collName, indexes)
	if err == nil {
		r.addCollection(collName)
	}
	return err
}

func (r *roomMsgOperator) checkAndCreateGroupColl(collName string) (bool, error) {
	create := func(collName string) error {
		if err := r.createGroupCollAndIndex(collName); err != nil {
			return err
		}
		return nil
	}
	return r.checkAndCreateColl(collName, create)
}

func (r *roomMsgOperator) checkAndCreateDiffusesColl(collName string) (bool, error) {
	create := func(collName string) error {
		if err := r.createDiffusesCollAndIndex(collName); err != nil {
			return err
		}
		return nil
	}
	return r.checkAndCreateColl(collName, create)
}

func (r *roomMsgOperator) checkAndCreateColl(collName string, create func(collName string) error) (bool, error) {
	find, err := r.checkCollection(collName)
	if err != nil {
		return false, err
	}
	if find {
		return true, nil
	}
	if err := create(collName); err != nil {
		return false, err
	}
	return true, nil
}

func getGroupCollName(groupId int64) string {
	return fmt.Sprintf("group_%d", groupId)
}

func getDiffusesCollName(userId int64) string {
	return fmt.Sprintf("user_%d", userId)
}

func parseGroupMsgList(cursor *mongo.Cursor) ([]GroupMsg, error) {
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	msgList := make([]GroupMsg, 0)
	if err := cursor.All(context.TODO(), &msgList); err != nil {
		return nil, err
	}
	sort.Slice(msgList, func(i, j int) bool {
		return msgList[i].Oid.String() < msgList[j].Oid.String()
	})
	return msgList, nil
}

func parseGroupMsg(result *mongo.SingleResult) (*GroupMsg, error) {
	if result.Err() != nil {
		return nil, result.Err()
	}
	msg := GroupMsg{}
	if err := result.Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

func parseDiffusesMsgList(cursor *mongo.Cursor) ([]DiffusesGroupMsg, error) {
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	msgList := make([]DiffusesGroupMsg, 0)
	if err := cursor.All(context.TODO(), &msgList); err != nil {
		return nil, err
	}
	//sort.Slice(msgList, func(i, j int) bool {
	//	return msgList[i].Oid.String() < msgList[j].Oid.String()
	//})
	return msgList, nil
}

func parseDiffusesMsg(result *mongo.SingleResult) (*DiffusesGroupMsg, error) {
	if result.Err() != nil {
		return nil, result.Err()
	}
	msg := DiffusesGroupMsg{}
	if err := result.Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
