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

- [ ] Authentication and Authorization
- [ ] Synchronous communication
  - RESTful
  - gRPC
- [ ] Service Discovery
  - consul
- [ ] Asynchronous communication
  - Messaging (kafka / rabbitmq)
- [ ] Persistant storage
  - mongodb
- [ ] Containerised deployment
  - docker
  - potential for kubernetes

### How i will implement this

The bookstore will be comprised of three services

- Books service to manage CRUD book operations and data
- Reviews service to manage reviews and raitings for a book
- Sales service to manage prices and sales of a book

All of these will be behind a RESTful API which will handle requests.

## Install

```sh
make build-all
```

## Usage

```sh
make dev
```

## Run tests

```sh
make test
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
