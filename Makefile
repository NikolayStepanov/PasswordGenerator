build:
	go build -o  bin/password-server ./cmd/app

run:
	@CGO_ENABLED=0 go run ./cmd/app/main.go
