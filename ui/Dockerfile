# Dockerfile for REmind UI
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build -- --configuration production --output-hashing none --output-mode static

# Production stage - use simple HTTP server
FROM node:20-alpine

# Install http-server globally
RUN npm install -g http-server

WORKDIR /app

# Copy built files
COPY --from=builder /app/dist/ui/browser ./static

# Expose port 8080 (Fly.io expects this)
EXPOSE 8080

# Start http-server
CMD ["http-server", "./static", "-p", "8080", "-a", "0.0.0.0", "--cors"]
