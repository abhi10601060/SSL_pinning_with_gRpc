package main

import (
	"log"
	"net/http"

	"github.com/abhi1060/rpc_client/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", pong)
	router.GET("/", handler.ListAllhandler)
	router.POST("/add", handler.AddBook)

	log.Fatal(router.Run(":8080"))
}


func pong(c *gin.Context){
	c.JSON(
		http.StatusOK,
		gin.H{
			"message" : "Pong",
		},
	)
}