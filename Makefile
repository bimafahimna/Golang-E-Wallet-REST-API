run:
	go run main.go

run_nodemon:
	nodemon --exec go run main.go --signal SIGTERM

test:
	go test -cover -v ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm -rf coverage.out

mock:
	mockery --all --case snake --output .internal/mocks

reset_sql:
	go run .

.PHONY: cli test coverage mock
