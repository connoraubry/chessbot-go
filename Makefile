build:
	go build -o bin/GameEngine main.go
test:
	go test -v ./...
benchmark:
	go test -bench=. ./...
run:
	go run main.go