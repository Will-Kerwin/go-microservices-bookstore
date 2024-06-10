package author

import (
	"context"
	"log"
	"time"

	"github.com/will-kerwin/go-microservice-bookstore/books/internal/db"
	"github.com/will-kerwin/go-microservice-bookstore/books/internal/ingester"
	"github.com/will-kerwin/go-microservice-bookstore/gen"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedAuthorServiceServer
	repository           db.AuthorRepository
	createAuthorIngester ingester.Ingester[models.CreateAuthorEvent]
	deleteAuthorIngester ingester.Ingester[models.DeleteAuthorEvent]
}

func New(repository db.AuthorRepository, addr string, groupID string) *Handler {

	createAuthorIngester, err := ingester.New[models.CreateAuthorEvent](addr, groupID, "createAuthor")
	if err != nil {
		log.Fatalf("Failed to create ingester: %s\n", err)
		createAuthorIngester = nil
	}
	deleteAuthorIngester, err := ingester.New[models.DeleteAuthorEvent](addr, groupID, "deleteAuthor")
	if err != nil {
		log.Fatalf("Failed to create ingester: %s\n", err)
		deleteAuthorIngester = nil
	}

	return &Handler{
		repository:           repository,
		createAuthorIngester: *createAuthorIngester,
		deleteAuthorIngester: *deleteAuthorIngester,
	}
}

func (h *Handler) HandleIngestors(ctx context.Context) {

	go h.handleCreateAuthorIngester(ctx)
	go h.handleDeleteAuthorIngester(ctx)

}

func (h *Handler) handleCreateAuthorIngester(ctx context.Context) {

	for {
		channel, err := h.createAuthorIngester.Ingest(ctx)

		if err != nil {
			log.Fatalf("Failed to ingest: %s\n", err)
		}

		for event := range channel {
			log.Println("Processing create author message")
			err := h.CreateAuthor(ctx, &event)
			if err != nil {
				log.Fatalf("Failed to put rating: %s\n", err)
			}
		}

		// sleep message process every 10 seconds
		time.Sleep(10 * time.Second)
	}

}

func (h *Handler) handleDeleteAuthorIngester(ctx context.Context) {
	for {
		channel, err := h.deleteAuthorIngester.Ingest(ctx)

		if err != nil {
			log.Fatalf("Failed to ingest: %s\n", err)
		}

		for event := range channel {
			log.Println("Processing delete author message")
			err := h.DeleteAuthor(ctx, &event)
			if err != nil {
				log.Fatalf("Failed to put rating: %s\n", err)
			}
		}

		// sleep message process every 10 seconds
		time.Sleep(10 * time.Second)
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

func (h *Handler) CreateAuthor(ctx context.Context, req *models.CreateAuthorEvent) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "req was nil")
	}
	if req.DateOfBirth == time.Now() {
		return status.Errorf(codes.InvalidArgument, "date of birth was 0")
	}
	if req.Name == "" {
		return status.Errorf(codes.InvalidArgument, "name was empty")
	}

	log.Println("Create new author")

	dob := req.DateOfBirth

	author := &models.Author{
		Name:        req.Name,
		DateOfBirth: dob,
	}

	_, err := h.repository.Add(ctx, author)

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (h *Handler) DeleteAuthor(ctx context.Context, req *models.DeleteAuthorEvent) error {
	if req == nil || req.ID == "" {
		return status.Errorf(codes.InvalidArgument, "req was nil, or id was empty")
	}

	log.Printf("request get author with id: %s", req.ID)

	err := h.repository.Delete(ctx, req.ID)

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}
