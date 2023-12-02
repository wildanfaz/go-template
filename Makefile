install:
	go mod tidy
	go mod download
	go mod vendor

start:
	go run main.go start