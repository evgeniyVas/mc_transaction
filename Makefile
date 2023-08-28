

gen-models:
	go run github.com/go-swagger/go-swagger/cmd/swagger@latest generate model -t ./internal/ -f swagger.yml
dev-infra:
	docker-compose -f ./build/docker-compose.yml up -d postgres
dev-migrate:
	goose -dir ./migrations postgres "user=user123 password=transactions dbname=transactions sslmode=disable host=localhost port=5431" up
