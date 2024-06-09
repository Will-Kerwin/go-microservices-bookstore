package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	authorGateway "github.com/will-kerwin/go-microservice-bookstore/api-service/internal/gateway/author"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/rest/author"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/rest/book"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/discovery"
)

const serviceName = "api"

func main() {

	var port int
	flag.IntVar(&port, "port", 8081, "Port to listen on")
	flag.Parse()

	// register with consul
	registryUri := os.Getenv("CONSUL_URI")
	regisrty, err := discovery.NewRegistry(registryUri)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := regisrty.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
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

	// setup router
	router := echo.New()

	// setup grpc gateways
	authorGateway := authorGateway.New(*regisrty)

	// setup handlers
	authorHandler := author.New(authorGateway)
	bookHandler := book.New()

	// init handlers
	authorHandler.Register(router)
	bookHandler.Register(router)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", port)))

}
