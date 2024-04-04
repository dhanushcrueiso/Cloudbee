package server

import (
	"Cloudbee/github.com/dhanushcrueiso/blog/protos"
	"context"
	"log"
)

type PostServiceServer struct {
	protos.UnimplementedPostServiceServer
}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *protos.CreatePostRequest) (*protos.CreatePostResponse, error) {
	// In a real implementation, you'd likely store the post in a database
	// For simplicity, let's just echo back the received post
	post := req.GetPost()
	log.Printf("Received request to create post: %+v", post)

	// Simulate generating a unique PostID (in real world, use UUID or similar)
	post.PostId = "123456789"

	return &protos.CreatePostResponse{Post: post}, nil
}
