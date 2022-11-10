build:
	go build -o bin/GameEngine main.go
test:
	go test -v ./...

run:
	go run main.go