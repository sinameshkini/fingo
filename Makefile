tidy:
	go mod tidy

run:
	go run cmd/fingo/main.go

.PHONY: test
test:
	go test --cover ./...