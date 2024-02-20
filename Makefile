
run:
	@go run cmd/api-debugger/main.go

uidev:
	@npx tailwindcss -i input.css -o static/output.css --watch

style:
	@npx tailwindcss -i input.css -o static/output.css

image:
	@docker build -t api-debugger .

example:
	@docker-compose -f examples/docker-compose.yaml up
