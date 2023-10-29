build:
	@go build -o bin/picpay

test:
	@go test -v ./...

# docker run --name picpay-postgres -e POSTGRES_PASSWORD=picpay -p 5432:5432 postgres
run:
	@go run main.go


