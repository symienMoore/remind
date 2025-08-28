# Dockerfile for REmind API Server serving Angular static files

# --- UI build stage ---
FROM node:20-alpine AS ui-builder

WORKDIR /ui

# Install dependencies
COPY ui/package*.json ./
RUN npm ci

# Copy UI source and build
COPY ui/ .
RUN npm run build -- --configuration production --output-hashing none --output-mode static

# --- Go build stage ---
FROM golang:1.25-alpine AS builder

# Install git and ca-certificates
RUN apk --no-cache add git ca-certificates

WORKDIR /app

# Copy go mod files
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy source code
COPY server/ ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# --- Final stage ---
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy built Angular static files
COPY --from=ui-builder /ui/dist/ui/browser ./static

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./main"]
