package test

import (
	"context"
	"testing"

	"Cloudbee/github.com/dhanushcrueiso/blog/protos"
	services "Cloudbee/internal/services"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestPostService_CreatePost(t *testing.T) {

	// Initialize a new instance of PostService with the mock database
	service := services.NewPostService()

	// Define a mock post object to be created
	mockPost := &protos.Post{
		PostId:          "2f4509f4-0f31-4028-b9b8-46b676dfe07c",
		Title:           "Test Post",
		Content:         "This is a test post.",
		Author:          "Test Author",
		PublicationDate: timestamppb.Now(),
		Tags:            []string{"test", "unit test"},
	}

	// Create a request object with the mock post
	req := &protos.CreatePostRequest{Post: mockPost}

	// Call the CreatePost method
	res, err := service.CreatePost(context.Background(), req)
	if err != nil {
		t.Errorf("Failed to create post: %v", err)
	}

	// Verify that the response contains the created post
	if res.Post == nil {
		t.Error("Response does not contain the created post")
	}

	// Verify that the created post matches the mock post

}

// func TestPostService_GetPost(t *testing.T) {
// 	// Initialize a new instance of PostService
// 	service := service.NewPostService()

// 	ctx := context.Background()

// 	// Create a request object
// 	req := protos.GetPostRequest{PostId: "3bc02f95-be3f-4ac1-95aa-ddc9b89079ea"}

// 	// Call the GetPost method
// 	res, err := service.GetPost(ctx, &req)
// 	if err != nil {
// 		t.Errorf("Failed to get post: %v", err)
// 	}

// 	fmt.Println(res)
// }
