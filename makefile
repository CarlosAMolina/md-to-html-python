build:
	go build -o cmoli-es-deploy .

dependencies:
	go mod tidy

format:
	go fmt

run:
	go run .
