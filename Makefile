tidy:
	go mod tidy

run:
	go run cmd/fingo/main.go

.PHONY: test
test:
	go test --cover ./...

refresh:
	go run cmd/migrator/main.go

tester:
	go run cmd/tester/main.go
