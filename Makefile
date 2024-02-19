run:
	@go run cmd/api-debugger/main.go

image:
	@docker build -t api-debugger .
