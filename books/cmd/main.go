package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/will-kerwin/go-microservice-bookstore/books/internal/db"
	"github.com/will-kerwin/go-microservice-bookstore/books/internal/grpc/author"
	"github.com/will-kerwin/go-microservice-bookstore/books/internal/grpc/book"
	"github.com/will-kerwin/go-microservice-bookstore/gen"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/discovery"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serviceName string = "books"

func main() {

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	// register with consul
	registryUri := os.Getenv("CONSUL_URI")
	kafkaUri := os.Getenv("KAFKA_URI")
	regisrty, err := discovery.NewRegistry(registryUri)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := regisrty.Register(ctx, instanceID, serviceName, fmt.Sprintf("%s:%d", serviceName, port)); err != nil {
		panic(err)
	}

	// Setup health check
	go func() {
		for {
			if err := regisrty.HealthCheck(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// deregister on close
	defer regisrty.Deregister(ctx, instanceID, serviceName)

	log.Printf("Starting the %s service at port: %d\n", serviceName, port)

	// load mongodb connection

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}

	log.Println("Connected to mongodb")

	// load repo and handler
	authorRepository := db.NewAuthorRepository(client)
	bookRepository := db.NewBookRepository(client)

	authorHandler := author.New(authorRepository, kafkaUri, serviceName)
	bookHandler := book.New(bookRepository, kafkaUri, serviceName)

	// handle ingestors
	authorHandler.HandleIngestors(ctx)
	bookHandler.HandleIngestors(ctx)

	// create grpc listener
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serviceName, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server and listen
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	gen.RegisterAuthorServiceServer(grpcServer, authorHandler)
	gen.RegisterBookServiceServer(grpcServer, bookHandler)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
