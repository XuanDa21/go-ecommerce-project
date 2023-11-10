package controllers

import (
	"net/http"

	services "ecommerce/services"

	"github.com/gin-gonic/gin"
)


func AddProductToCartHandler(c *gin.Context) {
	userId := c.Query("userId")
	productId := c.Query("productId")
	if userId == "" || productId == "" {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusNotFound, gin.H{"error": "Missed ID"})
		c.Abort()
		return
	}

	isAdded := services.AddProductToCart(userId, productId)
	if isAdded {
		c.JSON(http.StatusOK, gin.H{"Status": "Added product to card"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Error": "Falsed to add product to card"})
	}

}

func DeleteProductFromCartHandler(c *gin.Context) {
	userId := c.Query("userId")
	productId := c.Query("productId")
	isDelete := services.DeleteProductFromCart(userId, productId)
	if isDelete {
		c.JSON(http.StatusOK, gin.H{"Status": "Deleted product successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Error": "Falsed to delete"})
	}
}