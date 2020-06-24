package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lffranca/suflex-api/api"
	"github.com/lffranca/suflex-api/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Not loading .env file")
	}

	db, errDB := database.GetDBMeta()
	if errDB != nil {
		return nil, errDB
	}

	router := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	routerAPI := router.Group("api")
	{
		api.Init(routerAPI, db)
	}

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
