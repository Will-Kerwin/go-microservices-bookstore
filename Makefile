compose:
	@docker compose up -d

api-service:
	CONSUL_URI=localhost:8500 go run api-service/cmd/main.go --port 8081

books:
	MONGODB_URI=mongodb://localhost:27017 DbName=dbBooks CONSUL_URI=localhost:8500 go run books/cmd/main.go --port 8082

.PHONY: api-service books compose