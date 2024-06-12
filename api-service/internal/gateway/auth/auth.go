package auth

import (
	"context"

	"github.com/will-kerwin/go-microservice-bookstore/gen"
	"github.com/will-kerwin/go-microservice-bookstore/internal/grpcutil"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/discovery"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{
		registry: registry,
	}
}

func (g *Gateway) LoginUser(ctx context.Context, username string, password string) (*gen.LoginUserResponse, error) {

	conn, err := grpcutil.ServiceConnection(ctx, "auth", g.registry)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewUserServiceClient(conn)

	resp, err := client.LoginUser(ctx, &gen.LoginUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (g *Gateway) GetUser(ctx context.Context, id string) (*models.User, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "auth", g.registry)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewUserServiceClient(conn)

	resp, err := client.GetUser(ctx, &gen.GetUserRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return models.ProtoToUser(resp.User), err
}

func (g *Gateway) ValidateUsernameUnique(ctx context.Context, username string) (bool, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "auth", g.registry)
	if err != nil {
		return false, err
	}

	defer conn.Close()
	client := gen.NewUserServiceClient(conn)

	resp, err := client.ValidateUsernameUnique(ctx, &gen.ValidateUsernameUniqueRequest{
		Username: username,
	})

	if err != nil {
		return false, err
	}

	return resp.IsValid, err
}
