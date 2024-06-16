package auth

import (
	"context"
	"log"
	"time"

	"github.com/will-kerwin/go-microservice-bookstore/auth/internal/db"
	authModels "github.com/will-kerwin/go-microservice-bookstore/auth/pkg/models"
	"github.com/will-kerwin/go-microservice-bookstore/gen"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/ingester"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	userModels "github.com/will-kerwin/go-microservice-bookstore/pkg/models/user"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedUserServiceServer
	repository         db.AuthRepository
	createUserIngester ingester.Ingester[models.CreateUserEvent]
}

func New(repository db.AuthRepository, addr string, groupID string) *Handler {
	createUserIngester, err := ingester.New[models.CreateUserEvent](addr, groupID, "createUser")
	if err != nil {
		log.Fatalf("Failed to create ingester: %s\n", err)
		createUserIngester = nil
	}

	return &Handler{
		repository:         repository,
		createUserIngester: *createUserIngester,
	}
}

func (h *Handler) HandleIngestors(ctx context.Context) {
	go h.handleCreateUserIngester(ctx)
}

func (h *Handler) handleCreateUserIngester(ctx context.Context) {
	for {
		channel, err := h.createUserIngester.Ingest(ctx)

		if err != nil {
			log.Fatalf("Failed to ingest: %s\n", err)
		}

		for event := range channel {
			log.Println("Processing create author message")
			err := h.CreateUser(ctx, &event)
			if err != nil {
				log.Fatalf("Failed to put rating: %s\n", err)
			}
		}

		// sleep message process every 10 seconds
		time.Sleep(10 * time.Second)
	}
}

func (h *Handler) LoginUser(ctx context.Context, req *gen.LoginUserRequest) (*gen.LoginUserResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "req was nil")
	}
	if req.Username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username was empty")
	}
	if req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password was empty")
	}

	user, err := h.repository.Authenticate(ctx, req.Username, req.Password)

	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case authModels.ErrUnauthenticated:
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}
	}

	if user == nil {
		return nil, status.Errorf(codes.Unauthenticated, authModels.ErrUserNotFound.Error())
	}

	return &gen.LoginUserResponse{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "req or id was empty")
	}

	log.Printf("request get user with id: %s", req.Id)

	user, err := h.repository.GetById(ctx, req.Id)

	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &gen.GetUserResponse{User: userModels.UserToProto(user)}, nil
}

func (h *Handler) CreateUser(ctx context.Context, req *models.CreateUserEvent) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "req was nil")
	}
	if req.Username == "" {
		return status.Errorf(codes.InvalidArgument, "username was empty")
	}
	if req.Password == "" {
		return status.Errorf(codes.InvalidArgument, "password was empty")
	}
	if req.Email == "" {
		return status.Errorf(codes.InvalidArgument, "email was empty")
	}

	log.Printf("Create new user")

	_, err := h.repository.Add(ctx, req)

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (h *Handler) ValidateUsernameUnique(ctx context.Context, req *gen.ValidateUsernameUniqueRequest) (*gen.ValidateUsernameUniqueResponse, error) {
	if req == nil || req.Username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "req or id was empty")
	}

	isValid, err := h.repository.ValidateUsernameUnique(ctx, req.Username)

	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &gen.ValidateUsernameUniqueResponse{IsValid: isValid}, nil
}
