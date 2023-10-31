package controllers

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductViewerAdminHandeler(c *gin.Context) {
	c.Status(http.StatusOK)
}