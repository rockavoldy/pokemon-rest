package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rockavoldy/pokemon-rest/controller"
	"github.com/rockavoldy/pokemon-rest/model"
	"log"
	"os"
)
import "net/http"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	//time.Sleep(2000 * time.Millisecond) // need to make sure init db created first before connect to database
	model.Connect(dbUser, dbPass, dbHost, dbPort, dbName)

	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	controller.RouteType(api)
	controller.RoutePokemon(api)

	err = r.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
