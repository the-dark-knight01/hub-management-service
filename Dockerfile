# Step 1: Use the Golang image to build the app
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Step 2: Copy go mod and sum files to download dependencies
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Step 3: Copy the entire project into the container's /app directory
COPY . .

# Step 4: Set the working directory to where your main.go is located
WORKDIR /app/cmd/app

# Step 5: Build the Go application
RUN go build -o main .

# Step 6: Create a new stage with a minimal image to run the app
FROM alpine:latest

# Install ca-certificates, necessary tools, and the `migrate` tool
RUN apk --no-cache add ca-certificates curl bash \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar -xz -C /usr/local/bin \
    && chmod +x /usr/local/bin/migrate

# Set the working directory for the app inside the container
WORKDIR /root/

# Step 7: Copy the compiled Go binary and migration files from the builder image
COPY --from=builder /app/cmd/app/main .
COPY --from=builder /app/migrations /migrations
COPY --from=builder /app/.env .env


# Step 8: Expose the application port
EXPOSE 8080

# Step 9: Define the entrypoint to run migrations and start the application
CMD ["sh", "-c", "migrate -path /migrations -database \"postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable\" up && ./main"]