package main

// Write tests

import (
	"log"
	"os"

	"github.com/fibreactive/articlelate/templates"

	"github.com/fibreactive/articlelate/handler"
	"github.com/gin-gonic/gin"

	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	MONGO_URI := os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		MONGO_URI = "mongodb://localhost:27017"
	}
	options := options.Client().ApplyURI(MONGO_URI)
	err := mgm.SetDefaultConfig(nil, "articlelate", options)
	if err != nil {
		log.Fatal("Database error: ", err)
	}
}

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.HTMLRender = templates.LoadTemplates("./templates")
	h := handler.NewHandler()
	Routes(r, h)
	//gin.SetMode(gin.ReleaseMode)
	r.Run()
}
