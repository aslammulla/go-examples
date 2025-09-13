package main

import (
	"context"
	"log"
	"net"

	pb "grpc-example/productpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productServer struct {
	pb.UnimplementedProductServiceServer
}

func (s *productServer) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductReply, error) {
	products := map[string]*pb.ProductReply{
		"1": {Id: "1", Name: "Wireless Mouse", Description: "Bluetooth mouse", Price: 25.99},
	}

	product, exists := products[req.GetId()]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Product with ID %s not found", req.GetId())
	}

	return product, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &productServer{})

	log.Println("gRPC server is running on port 50051...")
	grpcServer.Serve(lis)
}

/*
OUTPUT:
$ go run main.go
2025/09/13 22:26:45 gRPC server is running on port 50051...
*/
