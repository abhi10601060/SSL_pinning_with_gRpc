package handler

import (
	"net/http"
	pb "github.com/abhi1060/proto_example/bookshop_protos"
	"github.com/abhi1060/rpc_client/grpc"
	"github.com/gin-gonic/gin"
)

func ListAllhandler(c *gin.Context) {
	res, err := grpc.ListAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": err.Error(),
			},
		)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, res)
}

func AddBook(ctx *gin.Context) {
	book := pb.Book{}
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

	res, err := grpc.AddBook(&book)
	if err != nil {
		ctx.JSON(
		http.StatusInternalServerError,
		gin.H{
			"message" : "internal server error", 
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, res)
}
