build:
	go build -o bin/GameEngine src/GameEngine/main.go
test:
	go test -v ./...

run:
	go run src/GameEngine/main.go