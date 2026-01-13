# Development container for automation-hub
FROM golang:1.25.5

# Set working directory
WORKDIR /workspace

# Copy go mod files first for better caching
COPY go.mod ./

# Download dependencies
RUN go mod download

# Install Go development tools
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Expose the application port
EXPOSE 8000

# Keep container running
CMD ["sleep", "infinity"]
