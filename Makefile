
style:
	@npx tailwindcss -i input.css -o web/assets/output.css

setup:
	@go mod download
	@npm ci
	@cp node_modules/htmx.org/dist/htmx.min.js web/assets
	@make style

run:
	@go run cmd/api-debugger/main.go --config-file=config.yaml --log-level=DEBUG

bin:
	@CGO_ENABLED=0 GOOS=linux go build -o ./api-debugger cmd/api-debugger/main.go

image:
	@docker build -t api-debugger .

container:
	@docker-compose down
	@docker-compose up
