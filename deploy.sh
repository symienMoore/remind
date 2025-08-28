#!/bin/bash

# REmind deployment script for Fly.io
set -e

echo "🚀 Deploying REmind to Fly.io..."

# Check if fly CLI is installed
if ! command -v fly &> /dev/null; then
    echo "❌ Fly CLI not found. Please install it first:"
    echo "   curl -L https://fly.io/install.sh | sh"
    exit 1
fi

# Check if app exists, create if not
if ! fly apps list | grep -q "remind"; then
    echo "📱 Creating new Fly.io app..."
    fly apps create remind
fi

# Deploy the app
echo "🔨 Building and deploying..."
fly deploy

echo "✅ Deployment complete!"
echo "🌐 Your app is available at: https://remind.fly.dev"
echo "📊 Monitor your app: fly status"
echo "📝 View logs: fly logs"
