package main

import (
	"log"

	"github.com/HumCoding/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/HumCoding/meu-primeiro-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Aboud to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8085"); err != nil {
		log.Fatal(err)
	}
}
