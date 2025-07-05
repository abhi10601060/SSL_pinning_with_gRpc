package handler

import (
	"log"
	"net/http"

	"github.com/abhi1060/rpc_client/model"
	"github.com/gin-gonic/gin"
)

func ListAllhandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		model.ListAllBookResponse{Count: 2, Books: []model.Book{{Id: 1, Name: "Harry"}, {Id: 2, Name: "Potter"}}})
}

func AddBook(ctx *gin.Context) {
	book := model.Book{}
	err := ctx.BindJSON(&book)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "not valid body",
			},
		)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})

	log.Println(book)
}
