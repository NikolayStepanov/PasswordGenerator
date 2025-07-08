build:
	go build -o  bin/cart-server ./cmd/app

run:
	@CGO_ENABLED=0 go run ./cmd/app/main.go
