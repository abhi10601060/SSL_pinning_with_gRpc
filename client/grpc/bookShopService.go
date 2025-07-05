package grpc

import (
	"context"
	"log"

	pb "github.com/abhi1060/proto_example/bookshop_protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var BookShopClient pb.BookShopServiceClient

func init() {
	// Create gRPC server connection
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	// Create client
	BookShopClient = pb.NewBookShopServiceClient(conn)
}

func ListAll() (*pb.ListAllResponse, error) {
	res, err := BookShopClient.ListAll(context.Background(), &pb.Void{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddBook(book *pb.Book) (*pb.AddBookResponse,error){
	res, err := BookShopClient.AddBook(context.Background(), book)
	if err != nil {
		return nil, err
	}
	return res, nil
}
