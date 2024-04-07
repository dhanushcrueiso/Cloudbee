package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"Cloudbee/config"
	"Cloudbee/github.com/dhanushcrueiso/blog/protos"
	services "Cloudbee/internal/services"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	db "Cloudbee/internal/database"
)

func TestPostService_CreatePost(t *testing.T) {
	InitialiseDB()
	service := services.NewPostService()
	var tags = []string{"check", "new"}
	// Define a mock post object to be created
	mockPost := protos.Post{
		PostId:          uuid.New().String(),
		Title:           "Test Post",
		Content:         "This is a test post.",
		Author:          "Test Author",
		PublicationDate: timestamppb.Now(),
		Tags:            tags,
	}

	// Create a request object with the mock post
	req := &protos.CreatePostRequest{Post: &mockPost}

	// Call the CreatePost method
	res, err := service.CreatePost(context.Background(), req)
	if err != nil {
		t.Errorf("Failed to create post: %v", err)
	}

	// Verify that the response contains the created post
	if res.Post == nil {
		t.Error("Response does not contain the created post")
	}

	// Verify that the created post matches the data in db

}

func TestPostService_GetPost(t *testing.T) {
	// Initialize a new instance of PostService
	InitialiseDB()
	service := services.NewPostService()

	ctx := context.Background()

	// Create a request object
	req := protos.GetPostRequest{PostId: "2f4509f4-0f31-4028-b9b8-46b676dfe17c"}

	// Call the GetPost method
	res, err := service.GetPost(ctx, &req)
	if err != nil {
		t.Errorf("Failed to get post: %v", err)
	}

	fmt.Println(res)
}

func InitialiseDB() {
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
}

func TestPostService_UpdatePost(t *testing.T) {

	InitialiseDB()
	// Initialize a new instance of PostService with the mock database
	service := services.NewPostService()

	// Define a mock post object with updated details
	mockPost := &protos.Post{
		PostId:          "2f4509f4-0f31-4028-b9b8-46b676dfe07c",
		Title:           "Updated Test Post",
		Content:         "This is an updated test post.",
		Author:          "Updated Test Author",
		PublicationDate: timestamppb.Now(),
		Tags:            []string{"check", "updated"},
	}

	// Create a request object with the mock post and its ID
	req := &protos.UpdatePostRequest{
		PostId:         mockPost.PostId,
		NewPostDetails: mockPost,
	}

	// Call the UpdatePost method
	res, err := service.UpdatePost(context.Background(), req)
	if err != nil {
		t.Errorf("Failed to update post: %v", err)
	}

	// Verify that the response contains the updated post details
	if res.UpdatedPostDetails == nil {
		t.Error("Response does not contain the updated post details")
	}

	fmt.Println(res)
}
