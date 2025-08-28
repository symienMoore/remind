#!/bin/bash

# REmind deployment script for Fly.io (Separate API and UI)
set -e

echo "🚀 Deploying REmind to Fly.io (Separate API and UI)..."

ENVIRONMENT="${1:-dev}"
case "$ENVIRONMENT" in
  dev)
    API_APP="remind-api-dev"
    UI_APP="remind-ui-dev"
    ;;
  qa)
    API_APP="remind-api-qa"
    UI_APP="remind-ui-qa"
    ;;
  prod|production)
    API_APP="remind-api"
    UI_APP="remind-ui"
    ;;
  *)
    echo "❌ Unknown environment: $ENVIRONMENT (expected dev|qa|prod)"
    exit 1
    ;;
esac

# Check if fly CLI is installed
if ! command -v fly &> /dev/null; then
    echo "❌ Fly CLI not found. Please install it first:"
    echo "   curl -L https://fly.io/install.sh | sh"
    exit 1
fi

# Deploy API Server
echo "🔧 Deploying API Server ($API_APP)..."
if ! fly apps list | grep -q "$API_APP"; then
    echo "📱 Creating new Fly.io app for API ($API_APP)..."
    fly apps create "$API_APP"
fi

echo "🚀 Deploying API server ($API_APP)..."
fly deploy --app "$API_APP"

# Deploy UI
echo "🎨 Deploying UI ($UI_APP)..."
if ! fly apps list | grep -q "$UI_APP"; then
    echo "📱 Creating new Fly.io app for UI ($UI_APP)..."
    fly apps create "$UI_APP"
fi

echo "🚀 Deploying UI ($UI_APP)..."
fly deploy --app "$UI_APP"

echo "✅ Deployment complete!"
echo ""
echo "🌐 Your services are available at:"
echo "   API Server: https://$API_APP.fly.dev"
echo "   UI: https://$UI_APP.fly.dev"
echo ""
echo "📊 Monitor your apps:"
echo "   API: fly status --app $API_APP"
echo "   UI: fly status --app $UI_APP"
echo ""
echo "📝 View logs:"
echo "   API: fly logs --app $API_APP"
echo "   UI: fly logs --app $UI_APP"
