<h1 align="center">Welcome to go-microservice-bookstore 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="https://twitter.com/WillKerwin" target="_blank">
    <img alt="Twitter: WillKerwin" src="https://img.shields.io/twitter/follow/WillKerwin.svg?style=social" />
  </a>
</p>

> Consolidate go web development by creating a microservice book store

## Implementation

### Core Requirements

- [x] Authentication and Authorization
  - I have implemented basic JWT authentication using a mongodb collection to store users.
  - To ensure security i have used bcrypt to hash passwords in the database
  - I have split my route handlers to a protected group to ensure they cannot be access by unauthenticated users. I have done this by using echo middleware
- [x] Synchronous communication
  - RESTful - To complete this i've created a api service in which all front end requests will be made RESTfully. I have also implemented swagger to ensure production standard documentation
  - gRPC - I use gRPC to manage service to service communication this includes communication from the api service to backend services
- [ ] Service Discovery
  - consul - I use consul to create a service registry to use for endpoints and connections.
  - TODO: Map the following to service discovery
    - mongodb
    - kafka
- [x] Asynchronous communication
  - Messaging i use Kafka to handle all asynchronous requests, this includes:
    - create, update, delete requests
- [x] Persistant storage
  - mongodb - i use mongo db to store the data for books and authors
  - redis - I user redis to cache frequently requested data. this inclides invalidating caches when an item is created or updated
- [x] Containerised deployment
  - docker - Containersied the api and book services
  - potential for kubernetes
- [x] API Practises
  - Pagination
  - Cache Aside strategy
  - Swagger Documentation

### How i will implement this

The bookstore will be comprised of three services

- [x] Books service to manage CRUD book and author operations
- [ ] Reviews service to manage reviews and raitings for a book
- [ ] Sales service to manage prices and sales of a book

All of these will be behind a RESTful API which will handle requests.

## Install

```sh
go mod tidy
```

## Usage

```sh
docker compose up -d --build
```

## Author

👤 **Will Kerwin**

- Twitter: [@WillKerwin](https://twitter.com/WillKerwin)
- Github: [@will-kerwin](https://github.com/will-kerwin)
- LinkedIn: [@will-kerwin](https://linkedin.com/in/will-kerwin)

## Credits

I've used a veriety of learning materials here so they are listed below

- go-microservices: [github repository](https://github.com/manavkush/microservices-go)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
