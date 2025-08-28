#!/bin/bash

# REmind deployment script for Fly.io (Separate API and UI)
set -e

echo "🚀 Deploying REmind to Fly.io (Separate API and UI)..."

# Check if fly CLI is installed
if ! command -v fly &> /dev/null; then
    echo "❌ Fly CLI not found. Please install it first:"
    echo "   curl -L https://fly.io/install.sh | sh"
    exit 1
fi

# Deploy API Server
echo "🔧 Deploying API Server..."
if ! fly apps list | grep -q "remind-api"; then
    echo "📱 Creating new Fly.io app for API..."
    fly apps create remind-api
fi

echo "🚀 Deploying API server..."
fly deploy --app remind-api

# Deploy UI
echo "🎨 Deploying UI..."
if ! fly apps list | grep -q "remind-ui"; then
    echo "📱 Creating new Fly.io app for UI..."
    fly apps create remind-ui
fi

echo "🚀 Deploying UI..."
fly deploy --app remind-ui

echo "✅ Deployment complete!"
echo ""
echo "🌐 Your services are available at:"
echo "   API Server: https://remind-api.fly.dev"
echo "   UI: https://remind-ui.fly.dev"
echo ""
echo "📊 Monitor your apps:"
echo "   API: fly status --app remind-api"
echo "   UI: fly status --app remind-ui"
echo ""
echo "📝 View logs:"
echo "   API: fly logs --app remind-api"
echo "   UI: fly logs --app remind-ui"
