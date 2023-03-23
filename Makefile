build:
	go build -o bin/GameEngine main.go
	go build -o bin/perft perft/perft.go

game:
	go build -o bin/Game game/game.go

test:
	go test ./...
test-verbose:
	go test -v ./...
benchmark:
	go test -bench=. ./...
run:
	go run main.go
perft:
	go run perft/perft.go