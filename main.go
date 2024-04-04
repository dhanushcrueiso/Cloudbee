package main

import (
	"log"
	"net"

	"Cloudbee/github.com/dhanushcrueiso/blog/protos"
	server "Cloudbee/internal"

	"google.golang.org/grpc"
)

func main() {
	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the PostServiceServer with the gRPC server
	protos.RegisterPostServiceServer(grpcServer, &server.PostServiceServer{})

	log.Println("gRPC server listening on port 50051")
	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
