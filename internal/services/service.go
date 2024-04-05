package server

import (
	"Cloudbee/github.com/dhanushcrueiso/blog/protos"
	"context"
	"log"
	"sync"
)

type PostServiceServer struct {
	protos.UnimplementedPostServiceServer
	posts map[string]*protos.Post
	mu    sync.Mutex
}

// NewPostService creates a new instance of PostService
func NewPostService() *PostServiceServer {
	return &PostServiceServer{
		posts: make(map[string]*protos.Post),
	}
}
func (s *PostServiceServer) CreatePost(ctx context.Context, req *protos.CreatePostRequest) (*protos.CreatePostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// In a real implementation, you'd likely store the post in a database
	// For now, let's store it in memory
	post := req.GetPost()
	log.Printf("Received request to create post: %+v", post)

	// Simulate generating a unique PostID (in real world, use UUID or similar)
	post.PostId = "123456789"

	s.posts[post.PostId] = post

	return &protos.CreatePostResponse{Post: post}, nil
}
