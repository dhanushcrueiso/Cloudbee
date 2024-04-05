package main

import (
	"log"
	"net"
	"os"

	"Cloudbee/github.com/dhanushcrueiso/blog/protos"
	service "Cloudbee/internal/services"

	"Cloudbee/config"
	db "Cloudbee/internal/database"

	"google.golang.org/grpc"
)

func main() {
	// Create a listener on TCP port 50051env := "dev"
	env := "dev"
	var file *os.File
	var err error

	file, err = os.Open(env + ".json")
	if err != nil {
		log.Println("Unable to open file. Err:", err)
		os.Exit(1)
	}
	//parsing json with the config and passing the dev.json values from here
	var cnf *config.Config
	config.ParseJSON(file, &cnf)
	config.Set(cnf)

	db.Init(&db.Config{
		URL:       cnf.DatabaseURL,
		MaxDBConn: cnf.MaxDBConn,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	postService := service.NewPostService()
	protos.RegisterPostServiceServer(grpcServer, postService)

	log.Println("gRPC server listening on port 50051")
	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
