// Middlewares handle tasks like authentication, validation, and error handling.
// keeping them separate improves code readability
package authentication

import (
	"log"
	"net/http"
	"strings"

	token "ecommerce/services/token"

	"github.com/gin-gonic/gin"
)

// this function will handle for middleware before run other routes
func Authentication(c *gin.Context) {
	cookieHeader := c.GetHeader("Cookie")
	var accessToken string

	if strings.HasPrefix(cookieHeader, "accessToken=") {
		accessToken = strings.TrimPrefix(cookieHeader, "accessToken=")
	}

	if accessToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missed token key"})
		log.Println("Missed token key")
		c.Abort()
		return
	}

	claims, err := token.ValidateToken(accessToken)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Println(err)
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Set("uid", claims.Uid)
	c.Next()
}