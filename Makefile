build:
	go build -o events-notification-bot cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...
