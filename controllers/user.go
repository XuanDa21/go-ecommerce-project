package controllers

import (
	"net/http"

	models "ecommerce/models/data"
	db "ecommerce/models/db"
	services "ecommerce/services"
	token "ecommerce/services/token"
	password "ecommerce/services/password"


	"github.com/gin-gonic/gin"
)

var (
	user models.User
)


func SignupHandeler(c *gin.Context) {

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error BindJson": err.Error()})
		return
	}

	isValidate := services.ValidateData(c, user)
	if !isValidate{
		return
	} 

	isExist := services.CheckUserExistence(c, user)
	if isExist {
		return
	}

	isUpdate := services.UpdateUserDataToMongo(c, user)
	if !isUpdate {
		return
	}

	c.JSON(http.StatusCreated, "Successfully Signed Up!!")

}

func LoginHandeler(c *gin.Context) {

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error BindJson": err.Error()})
		return
	}

	fouldUser, err := db.GetMongoClient().SearchUserByField(*user.Email, "email")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
		return
	}

	isVerify := password.VerifyPassword(*user.Password, *fouldUser.Password)
	if !isVerify {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
		return
	}

	currentToken, refreshToken:= token.TokenGenerator(fouldUser)
	token.UpdateAllTokens(currentToken, refreshToken, fouldUser.User_ID)

	c.JSON(http.StatusFound, fouldUser)

}

func SearchProductHandeler(c *gin.Context) {
	c.Status(http.StatusOK)
}

func SearchProductByQueryHandeler(c *gin.Context) {
	c.Status(http.StatusOK)
}
