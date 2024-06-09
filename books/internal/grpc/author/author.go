package author

import (
	"context"
	"log"

	"github.com/will-kerwin/go-microservice-bookstore/books/internal/db"
	"github.com/will-kerwin/go-microservice-bookstore/gen"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedAuthorServiceServer
	repository db.AuthorRepository
}

func New(repository db.AuthorRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) GetAuthors(ctx context.Context, req *gen.GetAuthorsRequest) (*gen.GetAuthorsResponse, error) {

	log.Println("Request Get authors")
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "req was nil")
	}

	authors, err := h.repository.Get(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.GetAuthorsResponse{Authors: models.AuthorsToProtos(authors)}, nil

}

func (h *Handler) GetAuthor(ctx context.Context, req *gen.GetAuthorRequest) (*gen.GetAuthorResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "req was nil, or id was empty")
	}

	log.Printf("request get author with id: %s", req.Id)

	// Get data logic refer back in sec
	author, err := h.repository.GetById(ctx, req.Id)

	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &gen.GetAuthorResponse{Author: models.AuthorToProto(author)}, nil
}
