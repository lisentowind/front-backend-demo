#!/bin/bash

# Platform detection script for Air configuration
# This script helps detect the current platform and set up the appropriate Air configuration

echo "🔧 Platform Detection and Setup Script"
echo "======================================"

# Detect the platform
PLATFORM=$(uname -s)
ARCH=$(uname -m)

echo "Detected Platform: $PLATFORM"
echo "Detected Architecture: $ARCH"

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

# Define the appropriate Air configuration file and binary path
if [[ "$PLATFORM" == "Darwin" ]]; then
    # macOS
    echo "🍎 macOS detected"
    AIR_CONFIG=".air.toml"
    BINARY_NAME="tmp/main"
    BUILD_CMD="go build -o ./tmp/main ./cmd/server/main.go"

    # Find Air binary
    AIR_BINARY=$(find_air_binary)
    if [[ -z "$AIR_BINARY" ]]; then
        echo "❌ Air binary not found in common locations"
        echo "Please install Air using: go install github.com/cosmtrek/air@latest"
        echo "Or make sure it's in your PATH"
        exit 1
    fi

elif [[ "$PLATFORM" == "Linux" ]]; then
    # Linux
    echo "🐧 Linux detected"
    AIR_CONFIG=".air.toml"
    BINARY_NAME="tmp/main"
    BUILD_CMD="go build -o ./tmp/main ./cmd/server/main.go"

    # Find Air binary
    AIR_BINARY=$(find_air_binary)
    if [[ -z "$AIR_BINARY" ]]; then
        echo "❌ Air binary not found in common locations"
        echo "Please install Air using: go install github.com/cosmtrek/air@latest"
        echo "Or make sure it's in your PATH"
        exit 1
    fi

elif [[ "$PLATFORM" == "MINGW"* ]] || [[ "$PLATFORM" == "MSYS"* ]] || [[ "$PLATFORM" == "CYGWIN"* ]]; then
    # Windows
    echo "🪟 Windows detected"
    AIR_CONFIG=".air.toml.windows"
    BINARY_NAME="tmp\\main.exe"
    BUILD_CMD="go build -o ./tmp/main.exe ./cmd/server/main.go"
    AIR_BINARY="air"  # Assume air is in PATH on Windows

else
    echo "❓ Unknown platform: $PLATFORM"
    echo "Using default Unix-like configuration"
    AIR_CONFIG=".air.toml"
    BINARY_NAME="tmp/main"
    BUILD_CMD="go build -o ./tmp/main ./cmd/server/main.go"

    # Find Air binary
    AIR_BINARY=$(find_air_binary)
    if [[ -z "$AIR_BINARY" ]]; then
        echo "❌ Air binary not found in common locations"
        echo "Please install Air using: go install github.com/cosmtrek/air@latest"
        echo "Or make sure it's in your PATH"
        exit 1
    fi
fi

echo ""
echo "Configuration Summary:"
echo "  Air Config: $AIR_CONFIG"
echo "  Binary Name: $BINARY_NAME"
echo "  Build Command: $BUILD_CMD"
echo "  Air Binary: $AIR_BINARY"
echo ""

# Check if the selected config file exists
if [[ -f "backend/$AIR_CONFIG" ]]; then
    echo "✅ Configuration file found: backend/$AIR_CONFIG"
else
    echo "❌ Configuration file not found: backend/$AIR_CONFIG"
    echo "Please ensure the configuration files exist."
    exit 1
fi

# Check if Air binary exists
if [[ -f "$AIR_BINARY" ]] || [[ "$AIR_BINARY" == "air" ]]; then
    echo "✅ Air binary found: $AIR_BINARY"
else
    echo "❌ Air binary not found at: $AIR_BINARY"
    echo "Please install Air using: go install github.com/cosmtrek/air@latest"
    echo "Or add it to your PATH if it's installed elsewhere"
    exit 1
fi

# Create a symlink or copy the appropriate config
echo ""
echo "Setting up Air configuration..."
if [[ "$PLATFORM" == "Darwin" ]] || [[ "$PLATFORM" == "Linux" ]]; then
    # For Unix-like systems, ensure the main config is correct
    if [[ -f "backend/.air.toml" ]]; then
        echo "✅ Main Air configuration already exists"
    else
        echo "❌ Main Air configuration missing"
        exit 1
    fi
else
    # For Windows, copy the Windows config to the main config
    cp "backend/$AIR_CONFIG" "backend/.air.toml"
    echo "✅ Copied Windows configuration to main config"
fi

echo ""
echo "🎉 Setup complete! You can now run:"
if [[ "$PLATFORM" == "Darwin" ]] || [[ "$PLATFORM" == "Linux" ]]; then
    echo "   cd backend && $AIR_BINARY"
    echo ""
    echo "Or from the root directory:"
    echo "   pnpm backend-dev"
else
    echo "   cd backend && air"
    echo ""
    echo "Or from the root directory:"
    echo "   pnpm backend-dev"
fi