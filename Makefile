build:
	@go build -o bin/carteira-banco-digital

test:
	@go test -v ./...

run:
	@go run main.go


