PROJECT_NAME=tinystack
DB_NAME=tinystack

build:
	go build -o $(PROJECT_NAME) cmd/app/main.go

run:
	make build
	./$(PROJECT_NAME)

test: 
	go test -v -cover ./...

verify:
	go mod tidy
	go mod verify
	go vet ./...
	go fmt ./...

check:
	make verify
	make build
	make test
	rm $(PROJECT_NAME)
