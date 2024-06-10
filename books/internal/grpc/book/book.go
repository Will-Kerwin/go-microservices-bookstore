package book

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
	gen.UnimplementedBookServiceServer
	repository         db.BookRepository
	createBookIngester ingester.Ingester[models.CreateBookEvent]
	deleteBookIngester ingester.Ingester[models.DeleteBookEvent]
	updateBookIngester ingester.Ingester[models.UpdateBookEvent]
}

func New(repository db.BookRepository, addr string, groupID string) *Handler {

	createBookIngester, err := ingester.New[models.CreateBookEvent](addr, groupID, "createBook")
	if err != nil {
		log.Fatalf("Failed to create ingester: %s\n", err)
		createBookIngester = nil
	}

	deleteBookIngester, err := ingester.New[models.DeleteBookEvent](addr, groupID, "deleteBook")
	if err != nil {
		log.Fatalf("Failed to create ingester: %s\n", err)
		deleteBookIngester = nil
	}

	updateBookIngester, err := ingester.New[models.UpdateBookEvent](addr, groupID, "updateBook")
	if err != nil {
		log.Fatalf("Failed to create ingester: %s\n", err)
		updateBookIngester = nil
	}

	return &Handler{
		repository:         repository,
		createBookIngester: *createBookIngester,
		updateBookIngester: *updateBookIngester,
		deleteBookIngester: *deleteBookIngester,
	}

}

func (h *Handler) HandleIngestors(ctx context.Context) {
	go h.handleCreateBookIngestor(ctx)
	go h.handleDeleteBookIngestor(ctx)
	go h.handleUpdateBookIngestor(ctx)
}

func (h *Handler) handleCreateBookIngestor(ctx context.Context) {
	for {
		channel, err := h.createBookIngester.Ingest(ctx)

		if err != nil {
			log.Fatalf("Failed to ingest: %s\n", err)
		}

		for event := range channel {
			log.Println("Processing create book message")
			err := h.CreateBook(ctx, &event)
			if err != nil {
				log.Fatalf("Failed to put rating: %s\n", err)
			}
		}

		// sleep message process every 10 seconds
		time.Sleep(10 * time.Second)
	}
}

func (h *Handler) handleUpdateBookIngestor(ctx context.Context) {

	for {
		channel, err := h.updateBookIngester.Ingest(ctx)

		if err != nil {
			log.Fatalf("Failed to ingest: %s\n", err)
		}

		for event := range channel {
			log.Println("Processing update book message")
			err := h.UpdateBook(ctx, &event)
			if err != nil {
				log.Fatalf("Failed to put rating: %s\n", err)
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func (h *Handler) handleDeleteBookIngestor(ctx context.Context) {
	for {
		channel, err := h.deleteBookIngester.Ingest(ctx)

		if err != nil {
			log.Fatalf("Failed to ingest: %s\n", err)
		}

		for event := range channel {
			log.Println("Processing delete book message")
			err := h.DeleteBook(ctx, &event)
			if err != nil {
				log.Fatalf("Failed to put rating: %s\n", err)
			}
		}

		// sleep message process every 10 seconds
		time.Sleep(10 * time.Second)
	}
}

func (h *Handler) GetBooks(ctx context.Context, req *gen.GetBooksRequest) (*gen.GetBooksResponse, error) {
	log.Println("Request Get Books")

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "req was nil")
	}

	books, err := h.repository.Get(ctx, req.Title, req.AuthorId, req.Genre)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &gen.GetBooksResponse{Books: models.BooksToProtos(books)}, nil
}

func (h *Handler) GetBook(ctx context.Context, req *gen.GetBookRequest) (*gen.GetBookResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "req was nil, or id was empty")
	}

	log.Printf("request get author with id: %s", req.Id)

	// Get data logic refer back in sec
	book, err := h.repository.GetById(ctx, req.Id)

	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &gen.GetBookResponse{Book: models.BookToProto(book)}, nil
}

func (h *Handler) CreateBook(ctx context.Context, req *models.CreateBookEvent) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "req was nil")
	}

	if req.Title == "" {
		return status.Errorf(codes.InvalidArgument, "date of birth was 0")
	}

	if req.Genre == "" {
		return status.Errorf(codes.InvalidArgument, "name was empty")
	}

	if req.AuthorId == "" {
		return status.Errorf(codes.InvalidArgument, "name was empty")
	}

	if req.Synopsis == "" {
		return status.Errorf(codes.InvalidArgument, "name was empty")
	}

	log.Println("Create new book")

	book := &models.Book{
		Title:    req.Title,
		AuthorId: req.AuthorId,
		Synopsis: req.Synopsis,
		ImageUrl: req.ImageUrl,
		Genre:    req.Genre,
	}

	_, err := h.repository.Add(ctx, book)

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (h *Handler) UpdateBook(ctx context.Context, req *models.UpdateBookEvent) error {

	err := h.repository.Update(ctx, req.ID, &req.Data)

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (h *Handler) DeleteBook(ctx context.Context, req *models.DeleteBookEvent) error {
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
