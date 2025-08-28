# Multi-stage Dockerfile for REmind app
# Stage 1: Build Angular UI
FROM node:18-alpine AS ui-builder
WORKDIR /app/ui
COPY ui/package*.json ./
RUN npm ci --only=production
COPY ui/ ./
RUN npm run build

# Stage 2: Build Go server
FROM golang:1.21-alpine AS server-builder
WORKDIR /app
COPY server/go.mod server/go.sum ./
RUN go mod download
COPY server/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 3: Final runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the built Go binary
COPY --from=server-builder /app/main .

# Copy the built Angular app to serve static files
COPY --from=ui-builder /app/ui/dist/ui/browser /root/ui

# Expose port 8080
EXPOSE 8080

# Run the server
CMD ["./main"]
