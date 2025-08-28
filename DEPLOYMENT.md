# Deployment Guide - Separate API and UI

This repository is configured for separate deployments of the API server and UI to Fly.io.

## Architecture

- **API Server** (`remind-api`): Go server handling `/reminders/*` endpoints
- **UI** (`remind-ui`): Angular app served as static files with Nginx

## Prerequisites

1. **Fly.io Account**: Sign up at [fly.io](https://fly.io)
2. **Fly CLI**: Install the Fly CLI locally first
3. **GitHub Repository**: Push your code to GitHub

## Setup Steps

### 1. Create Fly.io Apps (First Time Only)

```bash
# Install Fly CLI
curl -L https://fly.io/install.sh | sh

# Login to Fly.io
fly auth login

# Create API app
fly apps create remind-api

# Create UI app
fly apps create remind-ui
```

### 2. Generate Fly API Token

```bash
# Generate a new API token
fly auth token

# Copy the token - you'll need it for the next step
```

### 3. Add GitHub Secret

1. Go to your GitHub repository
2. Navigate to **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Name: `FLY_API_TOKEN`
5. Value: Paste the API token from step 2

### 4. Push to Main Branch

Once the secret is set up, every push to the main branch will trigger:

1. **Test and Build Job**:
   - Install dependencies
   - Build Angular UI
   - Test Go server
   - Build Go server

2. **Deploy API Job**:
   - Deploy Go server to `remind-api.fly.dev`

3. **Deploy UI Job**:
   - Deploy Angular app to `remind-ui.fly.dev`

## Manual Deployment

You can also deploy manually using the local script:

```bash
./deploy.sh
```

## Service URLs

- **API Server**: https://remind-api.fly.dev
- **UI**: https://remind-ui.fly.dev

## API Endpoints

- **Health Check**: `GET /ping`
- **API Info**: `GET /`
- **Reminders**: `GET /reminders`
- **Create Reminder**: `POST /reminders`
- **Get Reminder**: `GET /reminders/:id`
- **Update Reminder**: `PUT /reminders/:id`
- **Delete Reminder**: `DELETE /reminders/:id`

## Monitoring

- **API Status**: `fly status --app remind-api`
- **UI Status**: `fly status --app remind-ui`
- **API Logs**: `fly logs --app remind-api`
- **UI Logs**: `fly logs --app remind-ui`

## Troubleshooting

- **Check GitHub Actions**: Go to Actions tab in your repository
- **Verify Fly.io apps**: `fly apps list`
- **Check app status**: `fly status --app remind-api` or `fly status --app remind-ui`
- **View recent logs**: `fly logs --app remind-api` or `fly logs --app remind-ui`

## Environment Variables

### API Server
- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin framework mode (default: release)

### UI
- `PORT`: Nginx port (default: 80)

## Benefits of Separate Deployment

1. **Independent Scaling**: Scale API and UI separately
2. **Better Performance**: UI served by optimized Nginx
3. **Cleaner Architecture**: Clear separation of concerns
4. **Easier Maintenance**: Update API or UI independently
5. **Better Caching**: Static assets cached by CDN
6. **Cost Optimization**: Different resource allocation for each service
