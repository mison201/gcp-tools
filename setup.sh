#!/bin/bash

# GCP Port Forwarding Tool Setup Script
# This script helps you set up the development environment

set -e

echo "🚀 Setting up GCP Port Forwarding Tool..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.23+ first."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 16+ first."
    echo "   Visit: https://nodejs.org/"
    exit 1
fi

# Check if npm is installed
if ! command -v npm &> /dev/null; then
    echo "❌ npm is not installed. Please install npm first."
    exit 1
fi

# Check if gcloud is installed
if ! command -v gcloud &> /dev/null; then
    echo "⚠️  Google Cloud SDK (gcloud) is not installed."
    echo "   Please install it from: https://cloud.google.com/sdk/docs/install"
    echo "   This is required for the application to work."
fi

# Check if Wails is installed
if ! command -v wails &> /dev/null; then
    echo "📦 Installing Wails CLI..."
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
fi

echo "📦 Installing Go dependencies..."
go mod tidy

echo "📦 Installing frontend dependencies..."
cd frontend
npm install
cd ..

echo "🔨 Building application..."
wails build

echo "✅ Setup complete!"
echo ""
echo "🎯 Next steps:"
echo "1. Make sure you have gcloud CLI installed and authenticated"
echo "2. Run 'gcloud auth login' if you haven't already"
echo "3. Launch the application from build/bin/gcp-tools.app"
echo "4. Or run 'wails dev' for development mode"
echo ""
echo "📚 For more information, see README.md" 