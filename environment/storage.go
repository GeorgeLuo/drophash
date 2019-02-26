package environment

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"time"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

var DbClient *mongo.Client 
var DbCollection *mongo.Collection
var UpsertOption *options.ReplaceOptions
var DatabaseName string
var CollectionName string

func InitializeDatabaseClient(database_path string, 
	database_name string, collection_name string) error {
	if database_name != "" {
		DatabaseName = database_name
	} else {
		fmt.Println("database name undefined, using default: hash-to-messages")
		DatabaseName = "hash-to-messages"
	}
	if collection_name != "" {
		CollectionName = collection_name
	} else {
		fmt.Println("collection name undefined, using default: data")
		CollectionName = "data"
	}
	if database_path != "" {
		var err error
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		DbClient, err = mongo.Connect(ctx, options.Client().ApplyURI(database_path))
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("connected to %s", database_path)
		DbCollection = DbClient.Database(DatabaseName).Collection(CollectionName)
		UpsertOption = options.Replace()
		UpsertOption.SetUpsert(true)
		return nil
	} else {
		fmt.Println("database path undefined, using default: mongodb://localhost:27017")
		return InitializeDatabaseClient("mongodb://localhost:27017", DatabaseName, CollectionName)
	}
}

func DatabaseGetValue(hashKey string, messageValue *string) bool {
	filter := bson.M{"hashKey" : hashKey}
	var messageResponse DbMessageResponse
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := DbCollection.FindOne(ctx, filter).Decode(&messageResponse)
	if err != nil {
		fmt.Printf("DatabaseGetValue: %s\n", err)
	    return false
	} else {
		*messageValue = messageResponse.MessageValue
		return true
	}

	return false
}

func DatabaseInsertKV(hashKey string, messageValue string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// res, err := DbCollection.InsertOne(ctx, bson.M{"hashKey": hashKey, "messageValue": messageValue})
	// fmt.Printf("insert operation returned InsertID %s", res.InsertedID)

	filter := bson.M{"hashKey" : hashKey}
	replace := bson.M{"hashKey" : hashKey, "messageValue": messageValue}
	_, err := DbCollection.ReplaceOne(ctx, filter, replace, UpsertOption)

	if err != nil {
		fmt.Printf("DatabaseInsertKV: %s\n", err)
		return false
	}
	return true
}