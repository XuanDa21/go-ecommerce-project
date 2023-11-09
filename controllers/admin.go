package controllers

import (
	"log"
	"net/http"

	models "ecommerce/models/data"
	services "ecommerce/services"

	"github.com/gin-gonic/gin"
)

var (
	product models.Product
)

func AddProductByAdminHandeler(c *gin.Context) {
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error BindJson": err.Error()})
		return
	}

	//TODO: CHECK PRODUCT IS EXIST
	isUpdateProduct := services.UpdateProDataToMongo(c, product)

	if isUpdateProduct {
		log.Println("Successfully added our Product Admin!!")
	} else {
		log.Println("Failed to update")
	}
}