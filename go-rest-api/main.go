package main

import (
	"log"
	"os"

	"example/rest-api/db"
	"example/rest-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	PORT := os.Getenv("PORT")

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":" + PORT)

}
