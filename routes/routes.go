package routes

import (
	// controllers "ecommerce/controllers"
	"ecommerce/controllers"
	authentication "ecommerce/middlewares"

	"github.com/gin-gonic/gin"
)


func SetupRoutes(route *gin.Engine) {
	
	route.SetTrustedProxies(nil)

	//apis for user
	userGroupRouter := route.Group("/user")
	userGroupRouter.POST("/signup", controllers.SignupHandeler)
	userGroupRouter.POST("/login", controllers.LoginHandeler)
	userGroupRouter.GET("/productview", controllers.ProductViewHandler)
	userGroupRouter.GET("/search", controllers.SearchProductByQueryHandeler)
	
	//apis for admin
	adminGroupRouter := route.Group("/admin")
	adminGroupRouter.POST("/addproduction", controllers.AddProductByAdminHandeler)
	

	//handle authentication middleware before running any handle after to keep security
	route.Use(authentication.Authentication)
	

	//apis for cart 
	cartGroupRouter := route.Group("/cart")
	cartGroupRouter.PUT("/add-product-to-cart", controllers.AddProductToCartHandler)
	cartGroupRouter.GET("/delete-product-from-cart", controllers.DeleteProductFromCartHandler)
}


