compose:
	@docker compose up -d

api-service:
	CONSUL_URI=localhost:8500 KAFKA_URI=localhost go run api-service/cmd/main.go --port 8081

books:
	MONGODB_URI=mongodb://localhost:27017 KAFKA_URI=localhost DbName=dbBooks CONSUL_URI=localhost:8500 go run books/cmd/main.go --port 8082

protobuf:
	protoc -I=api --go_out=. --go-grpc_out=. bookstore.proto         

.PHONY: api-service books compose protobuf