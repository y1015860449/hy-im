package mongo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	client *mongo.Client
}

// https://docs.mongodb.com/manual/reference/connection-string/
func ConnectMongoDb(mongoURI string, maxPoolSize uint64) (*MongoClient, error) {
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(mongoURI).SetMaxPoolSize(maxPoolSize).SetPoolMonitor(&event.PoolMonitor{Event: func(poolEvent *event.PoolEvent) {
			// connectID 可以体现出当前操作使用的哪个connect
			//fmt.Printf("max connectID:%d \n", poolEvent.ConnectionID)
		}}))

	// todo 启用优先读取从节点的配置
	//	var readPref readpref.ReadPref
	//	opts.SetReadPreference(&readPref)
	if err != nil {
		log.Error("connect mongodb fail!")
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Error("mongo ping fail")
		return nil, err
	}
	return &MongoClient{client: client}, nil
}

func (cli *MongoClient) InsertOne(dbName string, collName string, doucment interface{}) (interface{}, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warn(err)
		return nil, err
	}
	var result *mongo.InsertOneResult
	result, err = cli.client.Database(dbName).Collection(collName).InsertOne(context.TODO(), doucment)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result.InsertedID, nil
}

func (cli *MongoClient) InserMany(dbName string, collName string, doucments []interface{}) ([]interface{}, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warn(err)
		return nil, err
	}
	var result *mongo.InsertManyResult

	result, err = cli.client.Database(dbName).Collection(collName).InsertMany(context.TODO(), doucments)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result.InsertedIDs, nil
}

func (cli *MongoClient) Update(dbName string, collName string, filter interface{}, update interface{}, bMany bool) (interface{}, int64, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warn(err)
		return nil, 0, err
	}
	var result *mongo.UpdateResult
	collection := cli.client.Database(dbName).Collection(collName)
	if bMany {
		result, err = collection.UpdateMany(context.TODO(), filter, update)
	} else {
		result, err = collection.UpdateOne(context.TODO(), filter, update)
	}
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}
	return result.UpsertedID, result.ModifiedCount + result.UpsertedCount, nil
}

func (cli *MongoClient) Replace(dbName string, collName string, filter interface{}, replacement interface{}) (interface{}, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warn(err)
		return nil, err
	}
	var result *mongo.UpdateResult
	result, err = cli.client.Database(dbName).Collection(collName).ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result.UpsertedID, nil
}

func (cli *MongoClient) Delete(dbName string, collName string, filter interface{}, bMany bool) (int64, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warn(err)
		return 0, err
	}
	var result *mongo.DeleteResult
	collection := cli.client.Database(dbName).Collection(collName)
	if bMany {
		result, err = collection.DeleteMany(context.TODO(), filter)
	} else {
		result, err = collection.DeleteOne(context.TODO(), filter)
	}
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.DeletedCount, nil
}

func (cli *MongoClient) Find(dbName string, collName string, filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warn(err)
		return nil, err
	}
	var cur *mongo.Cursor
	collection := cli.client.Database(dbName).Collection(collName)
	if opts != nil {
		cur, err = collection.Find(context.TODO(), filter, opts)
	} else {
		cur, err = collection.Find(context.TODO(), filter)
	}
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	if cur.Err() != nil {
		return nil, err
	}
	return cur, err
}

func (cli *MongoClient) Aggregate(dbName string, collName string, pipeline interface{}) (*mongo.Cursor, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warn(err)
		return nil, err
	}
	var cursor *mongo.Cursor

	if cursor, err = cli.client.Database(dbName).Collection(collName).Aggregate(context.TODO(), pipeline); err != nil {
		log.Errorf("%v", err)
		return nil, err
	}

	if cursor.Err() != nil {
		return nil, err
	}
	return cursor, err
}

func (cli *MongoClient) Count(dbName string, collName string, filter interface{}, opts *options.CountOptions) (int64, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warn(err)
		return 0, err
	}
	var count int64
	collection := cli.client.Database(dbName).Collection(collName)
	if opts != nil {
		count, err = collection.CountDocuments(context.TODO(), filter, opts)
	} else {
		count, err = collection.CountDocuments(context.TODO(), filter)
	}
	if err != nil {
		log.Errorf("%v", err)
		return 0, err
	}
	return count, err
}
