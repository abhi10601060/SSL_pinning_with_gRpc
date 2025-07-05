package main

import (
	"log"

	"github.com/abhi1060/rpc_client/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.ListAllhandler)
	router.POST("/add", handler.AddBook)

	log.Fatal(router.Run(":8080"))
}
