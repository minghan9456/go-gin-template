package main

import (
	"github.com/joho/godotenv"
	"log"
	"oauth2-interface.pts.org.tw/src/routes"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := routes.NewGinRouter()
	router.Gin.Run(":" + os.Getenv("APP_PORT"))
}
