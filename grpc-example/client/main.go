package main

import (
	"context"
	"log"
	"time"

	pb "grpc-example/productpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.ProductRequest{Id: "1"}
	res, err := client.GetProduct(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetProduct: %v", err)
	}

	log.Printf("Product Info:\nID: %s\nName: %s\nDescription: %s\nPrice: $%.2f",
		res.GetId(), res.GetName(), res.GetDescription(), res.GetPrice())
}

/*
OUTPUT:
$ go run main.go
2025/09/13 22:27:19 Product Info:
ID: 1
Name: Wireless Mouse
Description: Bluetooth mouse
Price: $25.99
*/
