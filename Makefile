
setup:
	@cd web/dependencies && npm i
	@go mod download

run:
	@go run cmd/api-debugger/main.go -config=example/config.yaml

bin:
	@CGO_ENABLED=0 GOOS=linux go build -o ./api-debugger cmd/api-debugger/main.go

style:
	@cd web/dependencies && npx tailwindcss -i input.css -o output.css

image:
	@docker build -t api-debugger .

# example:
# 	@docker-compose -f examples/docker-compose.yaml up
