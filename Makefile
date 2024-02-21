
setup:
	cd ui && npm i
	cd app && go mod download

run:
	@go run -C app cmd/api-debugger/main.go

binary:
	@CGO_ENABLED=0 GOOS=linux go build -C app -o ../api-debugger cmd/api-debugger/main.go

style:
	@cd ui && npx tailwindcss -i input.css -o style/output.css

wstyle:
	@cd ui && npx tailwindcss -i input.css -o style/output.css --watch

# image:
# 	@docker build -t api-debugger .

# example:
# 	@docker-compose -f examples/docker-compose.yaml up
