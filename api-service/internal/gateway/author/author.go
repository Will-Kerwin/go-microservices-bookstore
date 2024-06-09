package author

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

func (g *Gateway) Get(ctx context.Context) ([]*models.Author, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "books", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewAuthorServiceClient(conn)
	resp, err := client.GetAuthors(ctx, &gen.GetAuthorsRequest{})

	if err != nil {
		return nil, err
	}

	return models.ProtosToAuthors(resp.Authors), err
}

func (g *Gateway) GetById(ctx context.Context, id string) (*models.Author, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "books", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewAuthorServiceClient(conn)
	resp, err := client.GetAuthor(ctx, &gen.GetAuthorRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return models.ProtoToAuthor(resp.Author), err
}

func (g *Gateway) DeleteAuthor(ctx context.Context, id string) error {
	conn, err := grpcutil.ServiceConnection(ctx, "books", g.registry)

	if err != nil {
		return err
	}

	defer conn.Close()
	client := gen.NewAuthorServiceClient(conn)
	_, err = client.DeleteAuthor(ctx, &gen.DeleteAuthorRequest{
		Id: id,
	})

	if err != nil {
		return err
	}

	return nil
}

func (g *Gateway) CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "books", g.registry)

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := gen.NewAuthorServiceClient(conn)
	resp, err := client.CreateAuthor(ctx, &gen.CreateAuthorRequest{
		Name:        author.Name,
		DateOfBirth: author.DateOfBirth.Unix(),
	})

	if err != nil {
		return nil, err
	}

	return models.ProtoToAuthor(resp.Author), err
}
