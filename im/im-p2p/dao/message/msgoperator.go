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

type p2pMsgOperator struct {
	mCli *hymongodb.HyMongo
	dbName string
	collections map[string]interface{}
	collLock sync.RWMutex
}

func (p *p2pMsgOperator) InsertP2pMsg(userId int64, loginType int32, msg *P2pMsg) error {
	fromColl := getP2pCollName(userId)
	if _, err := p.checkAndCreateP2pColl(fromColl); err != nil {
		return err
	}

	toColl := getP2pCollName(msg.ToId)
	if _, err := p.checkAndCreateP2pColl(toColl); err != nil {
		return err
	}

	if _, err := p.mCli.InsertOne(p.dbName, toColl, getP2pMsgBson(msg)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}

	if loginType == imbase.LoginApp {
		msg.AppPulled = 1
	} else {
		msg.PcPulled = 1
	}
	if _, err := p.mCli.InsertOne(p.dbName, fromColl, getP2pMsgBson(msg)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}

	return nil
}

func (p *p2pMsgOperator) InsertP2pMsgList(userId int64, msgList []P2pMsg) error {
	collName := getP2pCollName(userId)
	if _, err := p.checkAndCreateP2pColl(collName); err != nil {
		return err
	}
	if _, err := p.mCli.InsertMany(p.dbName, collName, getP2pMsgListBson(msgList)); err != nil && err.Error() != "ErrorDuplicateKey" {
		return err
	}
	return nil
}

func (p *p2pMsgOperator) UpdateP2pMsgPulled(userId int64, loginType int32, msgIds []primitive.ObjectID) error {
	collName := getP2pCollName(userId)
	if _, err := p.checkAndCreateP2pColl(collName); err != nil {
		return err
	}
	filter := bson.M{"_id": bson.M{"$in": msgIds}}
	update := bson.M{}
	if loginType == imbase.LoginApp {
		update["appPulled"] = bson.M{"$set": 1}
	} else {
		update["pcPulled"] = bson.M{"$set": 1}
	}
	_, _, err := p.mCli.Update(p.dbName, collName, filter, update, true)
	return err
}

func (p *p2pMsgOperator) UpdateP2pMsgCancel(userId, toId int64, msgId primitive.ObjectID) error {
	fromColl := getP2pCollName(userId)
	if _, err := p.checkAndCreateP2pColl(fromColl); err != nil {
		return err
	}

	toColl := getP2pCollName(toId)
	if _, err := p.checkAndCreateP2pColl(toColl); err != nil {
		return err
	}

	filter := bson.M{"_id": msgId}
	update := bson.M{"$set": bson.M{"isCancel": 1}}
	_, _, err := p.mCli.Update(p.dbName, fromColl, filter, update, false)
	if err != nil {
		return err
	}
	_, _, err = p.mCli.Update(p.dbName, toColl, filter, update, false)
	return err
}

func (p *p2pMsgOperator) FindP2pMsgByClientMsgId(userId int64, clientMsgId string) (*P2pMsg, error) {
	collName := getP2pCollName(userId)
	if _, err := p.checkAndCreateP2pColl(collName); err != nil {
		return nil, err
	}
	result, err := p.mCli.FindOne(p.dbName, collName, getP2pMsgKey(userId, clientMsgId), nil)
	if err != nil {
		return nil, err
	}
	return parseGroupMsg(result)
}

func (p *p2pMsgOperator) FindP2pMsg(userId int64, msgId string) (*P2pMsg, error) {
	collName := getP2pCollName(userId)
	if _, err := p.checkAndCreateP2pColl(collName); err != nil {
		return nil, err
	}
	oid, _ := primitive.ObjectIDFromHex(msgId)
	filter := bson.M{"_id": oid}
	result, err := p.mCli.FindOne(p.dbName, collName, filter, nil)
	if err != nil {
		return nil, err
	}
	return parseGroupMsg(result)
}

func (p *p2pMsgOperator) FindP2pMsgListByLimit(userId int64, baseIndex string, limit int64, direction int32) ([]P2pMsg, error) {
	collName := getP2pCollName(userId)
	if _, err := p.checkAndCreateP2pColl(collName); err != nil {
		return nil, err
	}
	opt := options.Find().SetLimit(limit)
	opt = opt.SetSort(bson.M{"_id": -1})
	if direction == 1 {
		opt = opt.SetSort(bson.M{"_id": 1})
	}
	cursor, err := p.mCli.Find(p.dbName, collName, getP2pMsgByLimitKey(baseIndex, direction), opt)
	if err != nil {
		return nil, err
	}
	return parseP2pMsgList(cursor)
}

func (p *p2pMsgOperator) FindP2pOfflineMsg(userId int64, loginType int32) ([]P2pMsg, error) {
	// 查询扩散数据
	dCollName := getP2pCollName(userId)
	if _, err := p.checkAndCreateP2pColl(dCollName); err != nil {
		return nil, err
	}
	result, err := p.mCli.Find(p.dbName, dCollName, getOfflineP2pMsgKey(userId, loginType), nil)
	if err != nil {
		return nil, err
	}
	return parseP2pMsgList(result)
}

func NewP2pMsgOperator(dbName string, mCli *hymongodb.HyMongo) P2pMsgDao {
	opr := p2pMsgOperator{}
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

func (r *p2pMsgOperator) checkCollection(collName string) (bool, error) {
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

func (r *p2pMsgOperator) addCollection(collName string) {
	r.collLock.Lock()
	r.collections[collName] = 1
	r.collLock.Unlock()
}

func (r *p2pMsgOperator) createP2pCollAndIndex(collName string) error {
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

func (r *p2pMsgOperator) createCollAndIndex(collName string, mIndex func() []mongo.IndexModel) error {
	indexes := mIndex()
	_, err := r.mCli.CreateCollectionIndex(r.dbName, collName, indexes)
	if err == nil {
		r.addCollection(collName)
	}
	return err
}

func (r *p2pMsgOperator) checkAndCreateP2pColl(collName string) (bool, error) {
	create := func(collName string) error {
		if err := r.createP2pCollAndIndex(collName); err != nil {
			return err
		}
		return nil
	}
	return r.checkAndCreateColl(collName, create)
}

func (r *p2pMsgOperator) checkAndCreateColl(collName string, create func(collName string) error) (bool, error) {
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

func getP2pCollName(userId int64) string {
	return fmt.Sprintf("p2p_%d", userId)
}

func parseP2pMsgList(cursor *mongo.Cursor) ([]P2pMsg, error) {
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	msgList := make([]P2pMsg, 0)
	if err := cursor.All(context.TODO(), &msgList); err != nil {
		return nil, err
	}
	sort.Slice(msgList, func(i, j int) bool {
		return msgList[i].Oid.String() < msgList[j].Oid.String()
	})
	return msgList, nil
}

func parseGroupMsg(result *mongo.SingleResult) (*P2pMsg, error) {
	if result.Err() != nil {
		return nil, result.Err()
	}
	msg := P2pMsg{}
	if err := result.Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
