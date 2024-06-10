package book

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

func New(regisrty discovery.Registry) *Gateway {
	return &Gateway{
		registry: regisrty,
	}
}

func (g *Gateway) Get(ctx context.Context, title string, authorId string, genre string) ([]*models.Book, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "books", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewBookServiceClient(conn)
	resp, err := client.GetBooks(ctx, &gen.GetBooksRequest{})

	if err != nil {
		return nil, err
	}

	return models.ProtosToBooks(resp.Books), err
}

func (g *Gateway) GetById(ctx context.Context, id string) (*models.Book, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "books", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewBookServiceClient(conn)
	resp, err := client.GetBook(ctx, &gen.GetBookRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return models.ProtoToBook(resp.Book), err
}
