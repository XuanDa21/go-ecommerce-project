package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProductByAdminHandeler(c *gin.Context) {
	c.Status(http.StatusOK)
}