package services

// The services folder can hold complex business logic
// that isn't directly tied to a specific endpoint.
// This helps keep controllers clean and focused.

import (
	"context"
	"log"
	"net/http"
	"time"

	models "ecommerce/models/data"
	db "ecommerce/models/db"
	types "ecommerce/types"
	token "ecommerce/services/token"
	password "ecommerce/services/password"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func ValidateData(c *gin.Context, data any) bool{
	var validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}
	

func CheckUserExistence(ctx context.Context, c *gin.Context, user models.User) bool {
	//if call init function be careful about priority of compiler (global var -> init -> main) 
	userCollection := db.GetMongoClient().CreateCollection(types.UserCollectionName)
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return true
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return true
	}

	count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return true
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone is already in use"})
		return true
	}
	return false
}


func UpdateUserDataToMongo (ctx context.Context, c *gin.Context, user models.User) bool {
	userCollection := db.GetMongoClient().CreateCollection(types.UserCollectionName)
	hashPassword := password.CreateHashPassword(*user.Password)
	token, refreshtoken := token.TokenGenerator(user)
	
	user.Password = &hashPassword
	user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_ID = user.ID.Hex()
	user.Token = &token
	user.Refresh_Token = &refreshtoken
	user.UserCart = make([]models.ProductUser, 0)
	user.Address_Details = make([]models.Address, 0)
	user.Order_Status = make([]models.Order, 0)

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false
	}
	return true
}