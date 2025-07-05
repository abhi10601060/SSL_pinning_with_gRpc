package grpc

import (
	"context"
	"log"
	"os"

	pb "github.com/abhi1060/proto_example/bookshop_protos"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var BookShopClient pb.BookShopServiceClient
var envs map[string]string

const (
	grpcHostKey string = "GRPC_HOST"
)

func init() {
	grpcHost := os.Getenv("GRPC_HOST")
	if grpcHost == "" {
		envs, err := godotenv.Read(".env")
		if err!=nil {
			log.Println("error in reading .env " + err.Error())
			return
		}
		grpcHost = envs[grpcHostKey]
	}
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
