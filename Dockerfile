FROM golang:1.22.0-alpine3.19
WORKDIR /build
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./api-debugger cmd/api-debugger/main.go

# FROM node:20.11.1-alpine3.19
# WORKDIR /build
# COPY . ./
# RUN npm clean-install
# RUN npx tailwindcss -i input.css -o output.css

# FROM alpine:3.19.1
# WORKDIR /app
# COPY --from=0 /app/api-debugger /app/api-debugger
# COPY --from=1 /ui/style/output.css /ui/node_modules/htmx.org/dist/htmx.min.js /ui/
# EXPOSE 8080
# CMD ["/app/api-debugger"]
