package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/will-kerwin/go-microservice-bookstore/auth/internal/db"
	"github.com/will-kerwin/go-microservice-bookstore/auth/internal/grpc/auth"
	"github.com/will-kerwin/go-microservice-bookstore/gen"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/discovery"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serviceName string = "auth"

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

	registryUri := os.Getenv("CONSUL_URI")
	kafkaUri := os.Getenv("KAFKA_URI")
	registry, err := discovery.NewRegistry(registryUri)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("%s:%d", serviceName, port)); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Println("Failed to report heathly state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)

	log.Printf("Starting the %s service at port: %d\n", serviceName, port)

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

	// load repos
	authRespository := db.NewAuthRepository(client)

	// load handler
	authHandler := auth.New(authRespository, kafkaUri, serviceName)
	authHandler.HandleIngestors(ctx)

	// create grpc listener
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serviceName, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server and listen
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	gen.RegisterUserServiceServer(grpcServer, authHandler)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
