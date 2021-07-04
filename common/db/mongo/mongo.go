package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoClient struct {
	client *mongo.Client
}

// https://docs.mongodb.com/manual/reference/connection-string/
func ConnectMongoDb(mongoURI string, maxPoolSize uint64) (*MongoClient, error) {
	// 启用优先读取从节点的配置
	opt := options.Client()
	if rpf, err := readpref.New(readpref.SecondaryPreferredMode); err != nil {
		return nil, err
	} else {
		opt.ReadPreference = rpf
	}
	client, err := mongo.Connect(context.Background(),
		opt.ApplyURI(mongoURI).SetMaxPoolSize(maxPoolSize))

	if err != nil {
		log.Printf("connect mongodb fail!")
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		log.Printf("ping mongo fail err(%v)/n", err)
		return nil, err
	}
	return &MongoClient{client: client}, nil
}

func (cli *MongoClient) InsertOne(dbName string, collName string, document interface{}) (interface{}, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Print(err)
		return nil, err
	}
	var result *mongo.InsertOneResult
	result, err = cli.client.Database(dbName).Collection(collName).InsertOne(context.TODO(), document)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return result.InsertedID, nil
}

func (cli *MongoClient) InsertMany(dbName string, collName string, doucments []interface{}) ([]interface{}, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Print(err)
		return nil, err
	}
	var result *mongo.InsertManyResult

	result, err = cli.client.Database(dbName).Collection(collName).InsertMany(context.TODO(), doucments)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return result.InsertedIDs, nil
}

func (cli *MongoClient) Update(dbName string, collName string, filter interface{}, update interface{}, bMany bool) (interface{}, int64, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Print(err)
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
		log.Print(err)
		return nil, 0, err
	}
	return result.UpsertedID, result.ModifiedCount + result.UpsertedCount, nil
}

func (cli *MongoClient) Replace(dbName string, collName string, filter interface{}, replacement interface{}) (interface{}, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Print(err)
		return nil, err
	}
	var result *mongo.UpdateResult
	result, err = cli.client.Database(dbName).Collection(collName).ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return result.UpsertedID, nil
}

func (cli *MongoClient) Delete(dbName string, collName string, filter interface{}, bMany bool) (int64, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Print(err)
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
		log.Print(err)
		return 0, err
	}
	return result.DeletedCount, nil
}

func (cli *MongoClient) Find(dbName string, collName string, filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error) {
	var err error
	if err = cli.client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Print(err)
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
		log.Print("%v", err)
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
		log.Print(err)
		return nil, err
	}
	var cursor *mongo.Cursor

	if cursor, err = cli.client.Database(dbName).Collection(collName).Aggregate(context.TODO(), pipeline); err != nil {
		log.Printf("%v", err)
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
		log.Print(err)
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
		log.Printf("%v", err)
		return 0, err
	}
	return count, err
}
