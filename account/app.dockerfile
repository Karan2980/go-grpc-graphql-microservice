# Build stage
FROM golang:1.21-alpine3.18 AS build

# Install necessary packages including protobuf compiler
RUN apk --no-cache add gcc g++ make ca-certificates git protobuf protobuf-dev

# Set working directory
WORKDIR /go/src/github.com/Karan2980/go-grpc-graphql-microservice
# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install protoc plugins for Go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copy vendor directory if you're using vendor mode
COPY vendor vendor

# Copy account service source code
COPY account account

# Generate protobuf code (if needed)
WORKDIR /go/src/github.com/Karan2980/go-grpc-graphql-microservice/account
RUN if [ -f "account.proto" ]; then \
    protoc --go_out=. --go-grpc_out=. account.proto; \
    fi

# Build the application
WORKDIR /go/src/github.com/Karan2980/go-grpc-graphql-microservice
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./account/cmd/account

# Runtime stage
FROM alpine:3.18

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user for security
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /usr/bin

# Copy binary from build stage
COPY --from=build /go/bin/app .

# Change ownership to non-root user
RUN chown appuser:appgroup app

# Switch to non-root user
USER appuser

# Expose port (adjust if your service uses a different port)
EXPOSE 8080

# # Health check (optional)
# HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
#     CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./app"]
