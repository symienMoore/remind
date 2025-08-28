# Deployment Guide

## GitHub Actions Setup

This repository includes GitHub Actions workflows that automatically deploy to Fly.io on each push to the main branch.

### Prerequisites

1. **Fly.io Account**: Sign up at [fly.io](https://fly.io)
2. **Fly CLI**: Install the Fly CLI locally first
3. **GitHub Repository**: Push your code to GitHub

### Setup Steps

#### 1. Create Fly.io App (First Time Only)

```bash
# Install Fly CLI
curl -L https://fly.io/install.sh | sh

# Login to Fly.io
fly auth login

# Create the app (if not exists)
fly apps create remind
```

#### 2. Generate Fly API Token

```bash
# Generate a new API token
fly auth token

# Copy the token - you'll need it for the next step
```

#### 3. Add GitHub Secret

1. Go to your GitHub repository
2. Navigate to **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Name: `FLY_API_TOKEN`
5. Value: Paste the API token from step 2

#### 4. Push to Main Branch

Once the secret is set up, every push to the main branch will trigger:

1. **Test and Build Job**:
   - Install dependencies
   - Build Angular UI
   - Test Go server
   - Build Go server

2. **Deploy Job**:
   - Deploy to Fly.io
   - Verify deployment status

### Manual Deployment

You can also deploy manually using the local script:

```bash
./deploy.sh
```

### Monitoring

- **View deployment status**: `fly status`
- **View logs**: `fly logs`
- **Open app**: `fly open`

### Troubleshooting

- **Check GitHub Actions**: Go to Actions tab in your repository
- **Verify Fly.io app**: `fly apps list`
- **Check app status**: `fly status --app remind`
- **View recent logs**: `fly logs --app remind`

### Environment Variables

The app runs on port 8080 by default. You can add environment variables in the `fly.toml` file or using:

```bash
flyctl secrets set VARIABLE_NAME=value
```
