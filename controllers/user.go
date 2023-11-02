package controllers

import (
	"context"
	"net/http"
	"time"

	models "ecommerce/models/data"
	services "ecommerce/services"

	"github.com/gin-gonic/gin"
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
	if !isValidate{
		return
	} 

	isExist := services.CheckUserExistence(ctx, c, user)
	if isExist {
		return
	}

	isUpdate := services.UpdateUserDataToMongo(ctx, c, user)
	if !isUpdate {
		return
	}

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
