package main

import (
	"log"
	
	routes "ecommerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting my ecommerce app")
	log.Println("Initializing HTTP server...")

	router := gin.Default()
	routes.SetupRoutes(router)
	//test 2 trigger webhook
	port := ":5000"
	err := router.Run(port)
	if err != nil {
		log.Println("Cannot start server")
		log.Fatal(err)
	}
}
