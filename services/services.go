package services
// The services folder can hold complex business logic
// that isn't directly tied to a specific endpoint. 
// This helps keep controllers clean and focused.

import (
	"log"
	"context"
	"net/http"

	db "ecommerce/models/db"
	types "ecommerce/types"
	models "ecommerce/models/data"
			
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	userCollection = db.GetMongoClient().CreateCollection(types.UserCollectionName)
	proCollection = db.GetMongoClient().CreateCollection(types.ProCollectionName)
)

var validate = validator.New()

func ValidateData(c *gin.Context, data any) bool{
	err := validate.Struct(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

func CheckUserExistence(ctx context.Context, c *gin.Context, user models.User) bool {
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