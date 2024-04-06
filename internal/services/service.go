package server

import (
	"Cloudbee/github.com/dhanushcrueiso/blog/protos"
	"Cloudbee/internal/database/daos"
	"Cloudbee/internal/database/models"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

type PostServiceServer struct {
	protos.UnimplementedPostServiceServer
}

// NewPostService creates a new instance of PostService
func NewPostService() *PostServiceServer {
	return &PostServiceServer{}
}
func (s *PostServiceServer) CreatePost(ctx context.Context, req *protos.CreatePostRequest) (*protos.CreatePostResponse, error) {

	post := req.Post
	log.Printf("Received request to create post: %+v", post)
	post.PostId = uuid.New().String()

	daomodel := s.DtosToDao(post)
	err := daos.SaveCustomer(&daomodel)
	if err != nil {
		return nil, err
	}

	return &protos.CreatePostResponse{Post: post}, nil
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *protos.GetPostRequest) (*protos.GetPostResponse, error) {

	postID := req.PostId
	post, err := daos.GetPost(postID)
	if err != nil {
		return nil, err
	}
	Post := s.DaoToDtos(post)
	return &protos.GetPostResponse{
		Post: &Post,
	}, nil
}

func (s *PostServiceServer) GetAllPosts(ctx context.Context, emp *emptypb.Empty) (*protos.GetAllPostResponse, error) {

	res, err := daos.GetAllPosts()
	if err != nil {
		return nil, err
	}
	allRes := protos.GetAllPostResponse{
		Posts: make([]*protos.Post, 0, len(res)),
	}
	Posts := []*protos.Post{}
	for _, val := range res {
		Post := protos.Post{
			PostId:          val.Id.String(),
			Title:           val.Title,
			Content:         val.Content,
			Author:          val.Author,
			PublicationDate: ConvertTimetoTimestamppb(val.PublicationDate),
			Tags:            val.Tags,
		}
		// allRes.Post = append(allRes.Post, Post)
		Posts = append(Posts, &Post)

	}
	allRes.Posts = Posts
	fmt.Println("allres:", &allRes)
	return &allRes, nil
}

func (s *PostServiceServer) DtosToDao(req *protos.Post) models.Post {

	id, _ := uuid.Parse(req.PostId)
	date := ConvertTimestamppbToTime(req.PublicationDate)
	return models.Post{
		Id:              id,
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: date,
		Tags:            req.Tags,
	}
}

func ConvertTimestamppbToTime(ts *timestamp.Timestamp) time.Time {
	return time.Unix(ts.GetSeconds(), int64(ts.GetNanos())).UTC()
}

func (s *PostServiceServer) DaoToDtos(req *models.Post) protos.Post {

	id := req.Id.String()
	date := ConvertTimetoTimestamppb(req.PublicationDate)
	return protos.Post{
		PostId:          id,
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: date,
		Tags:            req.Tags,
	}
}

func ConvertTimetoTimestamppb(req time.Time) *timestamp.Timestamp {
	timestampProto := &timestamp.Timestamp{
		Seconds: int64(req.Unix()),
		Nanos:   int32(req.Nanosecond()),
	}
	return timestampProto
}
