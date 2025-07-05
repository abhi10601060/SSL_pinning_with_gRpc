package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/abhi1060/proto_example/bookshop_protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type bookShopServer struct {
	pb.UnimplementedBookShopServiceServer
}

func (s *bookShopServer) ListAll(ctx context.Context, req *pb.Void) (*pb.ListAllResponse, error) {
	// TODO: Implement your logic here
	return &pb.ListAllResponse{
		Books: []*pb.Book{
			{Name: "Harry Potter", Id: 0},
			{Name: "Harry Potter1", Id: 1},
			{Name: "Harry Potter2", Id: 2},
			{Name: "Harry Potter3", Id: 3},
			{Name: "Harry Potter4", Id: 4},
		},
		Count: 5,
	}, nil
}

func (s *bookShopServer) AddBook(ctx context.Context, req *pb.Book) (*pb.AddBookResponse, error) {
	// TODO: Implement your logic here
	if req.Id == 0 {
		return nil, errors.New("not valid book")
	}
	return &pb.AddBookResponse{
		IsSuccessful: true,
		Message:      "Book added successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookShopServiceServer(grpcServer, &bookShopServer{})

	// Enable reflection for grpcurl
	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}
