package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	authorGateway "github.com/will-kerwin/go-microservice-bookstore/api-service/internal/gateway/author"
	bookGateway "github.com/will-kerwin/go-microservice-bookstore/api-service/internal/gateway/book"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/rest/author"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/rest/book"
	_ "github.com/will-kerwin/go-microservice-bookstore/docs" // Import the docs
	"github.com/will-kerwin/go-microservice-bookstore/pkg/discovery"
)

const serviceName = "api"

// @title Go Microservice Bookstore API
// @version 1.0
// @description This is the api for the go bookstore microservices project
// @termsOfService http://swagger.io/terms/

// @contact.name Will Kerwin
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api-service:8080
// @BasePath /
// @schemes http
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

	// setup router
	router := echo.New()

	// setup grpc gateways
	authorGateway := authorGateway.New(*regisrty)
	bookGateway := bookGateway.New(*regisrty)

	// setup handlers
	authorHandler := author.New(authorGateway, kafkaUri)
	bookHandler := book.New(bookGateway, kafkaUri)

	// init handlers
	authorHandler.Register(router)
	bookHandler.Register(router)
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	// middleware

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", port)))

}
