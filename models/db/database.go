// Configuration files can be placed in the config folder.
// This includes environment variables, database configurations, and more.

package db

import (
	"context"
	"log"
	"sync"
	"time"

	models "ecommerce/models/data"
	types "ecommerce/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//here are global variables
var (
	initOnce = sync.Once{}
	mongoDB *MongoDB
)

type mongoConfig struct {
	mongoUrl string
	dbName   string
}

type MongoDB struct {
	mongoClient *mongo.Client
	config      *mongoConfig
}

type Collection struct {
	UserCollection *mongo.Collection
	ProductCollection *mongo.Collection
}

func getMongoConfig() *mongoConfig {
	return &mongoConfig{
		mongoUrl: types.MongoUrl,
		dbName:   types.DBName,
	}
}


func GetMongoClient() *MongoDB {
	initOnce.Do(func() {
		mongoConfig := getMongoConfig()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConfig.mongoUrl))
		if err != nil {
			log.Fatal(err)
		}
		//check if a MongoDB server has been found and connected to
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Println("failed to connect to mongodb")
		}
		mongoDB = &MongoDB{
			mongoClient: client,
			config:      mongoConfig,
		}
	})

	return mongoDB
}


func (c MongoDB) CreateCollection(collectionName string) *mongo.Collection {
	collection := c.mongoClient.Database(types.DBName).Collection(collectionName)
	if collection == nil {
		log.Println("Failed to create collection")
	}
	return collection
}

func (c MongoDB) GetCollection() (Collection) {
	collection := Collection {
		UserCollection:	c.CreateCollection(types.UserCollectionName),
		ProductCollection: c.CreateCollection(types.ProCollectionName),
	}
	return collection
}

// TODO: implement for multiple search users
func (c MongoDB) SearchUserByField(value any, field string) (result models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	collection := c.GetCollection().UserCollection
	err = collection.FindOne(ctx, bson.M{field: value}).Decode(&result)
	return result, err
}


func (c MongoDB) SearchProductByFiled(value any, field string, isMultipleSearch bool) (results []models.Product, result models.Product, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	collection := c.GetCollection().ProductCollection
	var filter any
	
	if field != "_id" {
		//Search Product by regex function
		filter = bson.M{field: bson.M{"$regex": value}}
	} else {
		filter = bson.M{field: value}
	}

	if isMultipleSearch {
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			log.Println(err.Error())
		}

		err = cursor.All(ctx, &results)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			log.Println(err.Error())
		}
	}

	return results, result, err
}


