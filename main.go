package main

import (
	"gin-gorm/src/db"
	"gin-gorm/src/route"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/suertle/go-package/logging"
)

func init() {
	// load ENV
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	// connection all databases
	err = db.Init()
	if err != nil {
		panic("failed to connect database")
	}

	// setup logger
	logging.Init("gin-gorm")
}

func main() {
	// create a Gin router
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := route.Init()

	logging.Debug("asdf")

	// start the server
	router.Run(":8080")
}
