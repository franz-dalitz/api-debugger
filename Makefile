
setup:
	@npm i
	@go mod download

run:
	@go run cmd/api-debugger/main.go -config=example/config.yaml

bin:
	@CGO_ENABLED=0 GOOS=linux go build -o ./api-debugger cmd/api-debugger/main.go

style:
	@npx tailwindcss -i input.css -o output.css

wstyle:
	@npx tailwindcss -i input.css -o output.css --watch

image:
	@docker build -t api-debugger .

# example:
# 	@docker-compose -f examples/docker-compose.yaml up
