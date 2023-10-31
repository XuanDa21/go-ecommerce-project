package main

import (
	"log"
	// "os"
	// "fmt"

	// db "ecommerce/models/db"
	routes "ecommerce/routes"

	// controllers "ecommerce/controllers"
	// middlewares "ecommerce/middlewares"
	// models "ecommerce/models"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting my ecommerce app")
	log.Println("Initializing HTTP server...")

	router := gin.Default()
	routes.SetupRoutes(router)

	port := ":5000"
	err := router.Run(port)
	if err != nil {
		log.Println("cannot start server: %s", err)
		log.Fatal(err)
	}
}