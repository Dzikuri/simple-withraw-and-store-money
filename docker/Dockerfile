# Build stage
FROM golang:1.23.0-alpine3.20 as build

# Install build dependencies
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code and build
COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/main.go

# Final stage
FROM alpine:latest

# Install certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /main

# Copy only the compiled binary and .env file
COPY --from=build /app/server /main/server
COPY .env /.env

# Set the timezone (optional)
ENV TZ=Asia/Jakarta

# Expose the application port
EXPOSE 8089

# Run the compiled binary
ENTRYPOINT ["/main/server"]