FROM golang:1.22.0-alpine
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-debugger cmd/api-debugger/main.go
EXPOSE 8080
CMD ["/api-debugger"]
