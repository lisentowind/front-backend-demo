#!/bin/bash

# Backend development script
# This script dynamically finds the Air binary and runs it

echo "🚀 Starting backend development server..."

# Function to find Air binary
find_air_binary() {
    # Check if air is in PATH first
    if command -v air &> /dev/null; then
        echo "$(command -v air)"
        return 0
    fi

    # Check common Go installation locations
    local go_path=$(go env GOPATH 2>/dev/null)
    local home_path="$HOME/go"

    # Check GOPATH/bin
    if [[ -n "$go_path" && -f "$go_path/bin/air" ]]; then
        echo "$go_path/bin/air"
        return 0
    fi

    # Check ~/go/bin
    if [[ -f "$home_path/bin/air" ]]; then
        echo "$home_path/bin/air"
        return 0
    fi

    # Check /usr/local/bin (common for global installations)
    if [[ -f "/usr/local/bin/air" ]]; then
        echo "/usr/local/bin/air"
        return 0
    fi

    # Not found
    echo ""
    return 1
}

# Find Air binary
AIR_BINARY=$(find_air_binary)

if [[ -z "$AIR_BINARY" ]]; then
    echo "❌ Air binary not found in common locations"
    echo ""
    echo "Please install Air using one of these methods:"
    echo "  1. go install github.com/cosmtrek/air@latest"
    echo "  2. brew install air (on macOS)"
    echo ""
    echo "Or make sure it's in your PATH"
    exit 1
fi

echo "✅ Found Air at: $AIR_BINARY"
echo "📁 Working directory: $(pwd)"
echo "======================================"

# Run Air
cd backend
exec "$AIR_BINARY"