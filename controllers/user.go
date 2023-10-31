package controllers 

import (
	// "fmt"
	"context"
	"net/http"
	"time"

	models "ecommerce/models/data"
	db "ecommerce/models/db"
	types "ecommerce/types"
	services "ecommerce/services"

	"github.com/gin-gonic/gin"
	
)

var (
	userCollection = db.GetMongoClient().CreateCollection(types.UserCollectionName)
	proCollection = db.GetMongoClient().CreateCollection(types.ProCollectionName)
)

func SignupHandeler(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error BindJson": err.Error()})
		return
	}
	
	isValidate := services.ValidateData(c, user)
	isExist := services.CheckUserExistence(ctx, c, user)
	if !isValidate || !isExist {
		return
	} 
	
	defer cancel()
	c.JSON(http.StatusCreated, "Successfully Signed Up!!")
}

func LoginHandeler(c *gin.Context) {
	c.Status(http.StatusOK)
}

func SearchProductHandeler(c *gin.Context) {
	c.Status(http.StatusOK)
}

func SearchProductByQueryHandeler(c *gin.Context) {
	c.Status(http.StatusOK)
}
