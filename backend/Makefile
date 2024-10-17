up:
	port=8080 api_secret=secret db_host=localhost db_port=5432 db_user=postgres db_password=postgres db_name=backend sslmode=disable go run cmd/app/main.go
db-up:
	docker-compose up -d db
unit-test: db-up
	go test -coverprofile=c.out -p 1 ./...
open-cov: unit-test
	go tool cover -html="c.out"