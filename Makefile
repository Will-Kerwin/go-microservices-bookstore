compose:
	@docker compose up -d

api-service:
	CONSUL_URI=localhost:8500 KAFKA_URI=broker PORT=8080 go run api-service/cmd/main.go

books:
	MONGODB_URI=mongodb://localhost:27017 KAFKA_URI=localhost DbName=dbBooks CONSUL_URI=localhost:8500 go run books/cmd/main.go --port 8082

protobuf:
	protoc -I=api --go_out=. --go-grpc_out=. bookstore.proto         

swag:
	swag init -g ./api-service/cmd/main.go --output ./docs

.PHONY: api-service books compose protobuf swag