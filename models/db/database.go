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


// func (c MongoDB) SearchByEmail(ctx context.Context, user models.User) (foundUser *models.User, err error) {
// 	userCollection := c.CreateCollection(types.UserCollectionName)
// 	err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
// 	return foundUser, err
// }

func (c MongoDB) SearchUserByField(value any, field string) (result models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	collection := c.CreateCollection(types.UserCollectionName)
	err = collection.FindOne(ctx, bson.M{field: value}).Decode(&result)
	return result, err
}


