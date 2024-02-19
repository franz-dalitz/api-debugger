run:
	@go run main.go

image:
	@docker build -t api-debugger .
