

gen-models:
	go run github.com/go-swagger/go-swagger/cmd/swagger@latest generate model -t ./internal/ -f swagger.yml
dev-infra:
	docker-compose -f ./build/docker-compose.yml up -d postgres wiremock-wwe-server
dev-migrate:
	goose -dir ./migrations postgres "user=user123 password=transactions dbname=transactions sslmode=disable host=localhost port=5429" up
gen-api-models:
	swagger generate model -t ./internal/http_client/paysystem/ -f ./internal/http_client/paysystem/swagger.yml --model-package models
gen-api-client-paysystem:
	swagger generate client -t ./internal/http_client/paysystem/ -f ./internal/http_client/paysystem/swagger.yml -c payclient