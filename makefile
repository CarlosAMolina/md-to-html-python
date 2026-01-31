# CGO_ENABLED = 0 to avoid these errors in the VPS:
# ```
# ./cmoli-es-deploy: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.34' not found (required by ./cmoli-es-deploy)
# ./cmoli-es-deploy: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.32' not found (required by ./cmoli-es-deploy)
# ```
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmoli-es-deploy

dependencies:
	go mod tidy

format:
	go fmt

run:
	go run .

send:
	scp cmoli-es-deploy dev:~/Software/
