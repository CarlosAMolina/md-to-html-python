compile:
	go build -o cmoli-es-deploy main.go

dependencies:
	go mod tidy

format:
	go fmt

run:
	go run main.go
