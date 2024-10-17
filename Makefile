db-up:
	docker stop test_db
	docker rm test_db
	docker run -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=test_db -p 5432:5432 --name test_db -d postgres
unit-test: db-up
	cd backend && go test -coverprofile=c.out -p 1 ./...
	docker stop test_db
	docker rm test_db
open-cov: unit-test
	go tool cover -html="c.out"