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

type groupMsgOperator struct {
	mCli *hymongodb.HyMongo
	dbName string
	collections map[string]interface{}
	collLock sync.RWMutex
}

func (g *groupMsgOperator) InsertGroupMsg(groupId int64, msg *GroupMsg) error {
	collName := getGroupCollName(groupId)
	if _, err := g.checkAndCreateGroupColl(collName); err != nil {
		return err
	}
	if _, err := g.mCli.InsertOne(g.dbName, collName, getGroupMsgBson(msg)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}
	return nil
}

func (g *groupMsgOperator) InsertGroupMsgList(groupId int64, msgList []GroupMsg) error {
	collName := getGroupCollName(groupId)
	if _, err := g.checkAndCreateGroupColl(collName); err != nil {
		return err
	}
	if _, err := g.mCli.InsertMany(g.dbName, collName, getGroupMsgListBson(msgList)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}
	return nil
}

func (g *groupMsgOperator) FindGroupMsgByClientMsgId(groupId, userId int64, clientMsgId string) (*GroupMsg, error) {
	collName := getGroupCollName(groupId)
	if _, err := g.checkAndCreateGroupColl(collName); err != nil {
		return nil, err
	}
	result, err := g.mCli.FindOne(g.dbName, collName, getGroupMsgKey(userId, clientMsgId), nil)
	if err != nil {
		return nil, err
	}
	return parseGroupMsg(result)
}

func (g *groupMsgOperator) FindGroupMsg(groupId int64, msgId string) (*GroupMsg, error) {
	collName := getGroupCollName(groupId)
	if _, err := g.checkAndCreateGroupColl(collName); err != nil {
		return nil, err
	}
	oid, _ := primitive.ObjectIDFromHex(msgId)
	filter := bson.M{"_id": oid}
	result, err := g.mCli.FindOne(g.dbName, collName, filter, nil)
	if err != nil {
		return nil, err
	}
	return parseGroupMsg(result)
}

func (g *groupMsgOperator) FindGroupMsgListByLimit(groupId int64, baseIndex string, limit int64, direction int32) ([]GroupMsg, error) {
	collName := getGroupCollName(groupId)
	if _, err := g.checkAndCreateGroupColl(collName); err != nil {
		return nil, err
	}
	opt := options.Find().SetLimit(limit)
	opt = opt.SetSort(bson.M{"_id": -1})
	if direction == 1 {
		opt = opt.SetSort(bson.M{"_id": 1})
	}
	cursor, err := g.mCli.Find(g.dbName, collName, getGroupMsgByLimitKey(baseIndex, direction), opt)
	if err != nil {
		return nil, err
	}
	return parseGroupMsgList(cursor)
}

func (g *groupMsgOperator) UpdateGroupMsgCancel(groupId int64, msgId string) error {
	collName := getGroupCollName(groupId)
	if _, err := g.checkAndCreateGroupColl(collName); err != nil {
		return err
	}
	oid, _ := primitive.ObjectIDFromHex(msgId)
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"isCancel": 1}}
	_, _, err := g.mCli.Update(g.dbName, collName, filter, update, false)
	return err
}

func (g *groupMsgOperator) UpdateP2pMsgPulled(userId int64, loginType int32, msgIds []primitive.ObjectID) error {
	dCollName := getDiffusesCollName(userId)
	if _, err := g.checkAndCreateDiffusesColl(dCollName); err != nil {
		return err
	}
	filter := bson.M{"_id": bson.M{"$in": msgIds}}
	update := bson.M{}
	if loginType == imbase.LoginApp {
		update["appPulled"] = bson.M{"$set": 1}
	} else {
		update["pcPulled"] = bson.M{"$set": 1}
	}
	_, _, err := g.mCli.Update(g.dbName, dCollName, filter, update, true)
	return err
}

// todo 这个方法使用事务才比较合理
func (g *groupMsgOperator) InsertDiffusesGroupMsg(fromId int64, loginType int32, members []int64, msg *DiffusesGroupMsg, groupMsg *GroupMsg) error {
	if err := g.InsertGroupMsg(groupMsg.GroupId, groupMsg); err != nil {
		return err
	}
	dBson := getDiffusesMsgBson(msg)
	for _, v := range members{
		collName := getDiffusesCollName(v)
		if _, err := g.checkAndCreateDiffusesColl(collName); err != nil {
			return err
		}
		if v == fromId {
			if loginType == imbase.LoginApp {
				msg.AppPulled = 1
			} else {
				msg.PcPulled = 1
			}
			if _, err := g.mCli.InsertOne(g.dbName, collName, getDiffusesMsgBson(msg)); err != nil && err.Error() != "ErrorDuplicateKey" {
				return err
			}
		} else {
			if _, err := g.mCli.InsertOne(g.dbName, collName, dBson); err != nil && err.Error() != "ErrorDuplicateKey" {
				return err
			}
		}
	}
	return nil
}

func (g *groupMsgOperator) FindGroupOfflineMsg(userId int64, loginType int32, groupId int64) ([]GroupMsg, error) {
	// 查询扩散数据
	dCollName := getDiffusesCollName(userId)
	if _, err := g.checkAndCreateDiffusesColl(dCollName); err != nil {
		return nil, err
	}
	result, err := g.mCli.Find(g.dbName, dCollName, getOfflineGroupMsgKey(userId, loginType, groupId), nil)
	if err != nil {
		return nil, err
	}
	dMsgList, err := parseDiffusesMsgList(result)
	if err != nil {
		return nil, err
	}

	// 查询原始数据
	gCollName := getGroupCollName(groupId)
	if _, err := g.checkAndCreateGroupColl(gCollName); err != nil {
		return nil, err
	}
	var ids []primitive.ObjectID
	for _, v := range dMsgList {
		ids = append(ids, v.Oid)
	}
	filter := bson.M{"_oid": bson.M{"$in": ids}}
	rst, err := g.mCli.Find(g.dbName, gCollName, filter, nil)
	if err != nil {
		return nil, err
	}
	return parseGroupMsgList(rst)
}

func NewGroupMsgOperator(dbName string, mCli *hymongodb.HyMongo) GroupMsgDao {
	opr := groupMsgOperator{}
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

func (g *groupMsgOperator) checkCollection(collName string) (bool, error) {
	g.collLock.RLock()
	if _, ok := g.collections[collName]; ok {
		g.collLock.RUnlock()
		return true, nil
	}
	g.collLock.RUnlock()

	has, err := g.mCli.HasCollection(g.dbName, collName)
	if err != nil {
		log.Errorf("mongodb is broken")
		return false, err
	}
	if has {
		g.addCollection(collName)
		return true, nil
	}
	return false, nil
}

func (g *groupMsgOperator) addCollection(collName string) {
	g.collLock.Lock()
	g.collections[collName] = 1
	g.collLock.Unlock()
}

func (g *groupMsgOperator) createGroupCollAndIndex(collName string) error {
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
	return g.createCollAndIndex(collName, mIndex)
}

func (g *groupMsgOperator) createDiffusesCollAndIndex(collName string) error {
	mIndex := func() []mongo.IndexModel {
		flag := true
		return []mongo.IndexModel{
			{
				Keys:    bson.D{{"groupId", 1}},
				Options: &options.IndexOptions{Background: &flag},
			},
		}
	}
	return g.createCollAndIndex(collName, mIndex)
}

func (g *groupMsgOperator) createCollAndIndex(collName string, mIndex func() []mongo.IndexModel) error {
	indexes := mIndex()
	_, err := g.mCli.CreateCollectionIndex(g.dbName, collName, indexes)
	if err == nil {
		g.addCollection(collName)
	}
	return err
}

func (g *groupMsgOperator) checkAndCreateGroupColl(collName string) (bool, error) {
	create := func(collName string) error {
		if err := g.createGroupCollAndIndex(collName); err != nil {
			return err
		}
		return nil
	}
	return g.checkAndCreateColl(collName, create)
}

func (g *groupMsgOperator) checkAndCreateDiffusesColl(collName string) (bool, error) {
	create := func(collName string) error {
		if err := g.createDiffusesCollAndIndex(collName); err != nil {
			return err
		}
		return nil
	}
	return g.checkAndCreateColl(collName, create)
}

func (g *groupMsgOperator) checkAndCreateColl(collName string, create func(collName string) error) (bool, error) {
	find, err := g.checkCollection(collName)
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
