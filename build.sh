#!/bin/bash

# ================= 配置 =================
# 你的 Go 代码所在的子目录名
BACKEND_DIR="backend"
# 你的入口文件 (相对于 backend 目录)
ENTRY_FILE="./cmd/server"
# 生成的文件名前缀
APP_NAME="backend-server"
# 输出目录 (在根目录下)
OUTPUT_DIR="$(pwd)/bin"
# =======================================

echo "🚀 开始构建..."

# 1. 检查并进入 backend 目录
if [ -d "$BACKEND_DIR" ]; then
    cd "$BACKEND_DIR" || exit
    echo "📂 已进入 $BACKEND_DIR 目录"
else
    echo "❌ 错误：找不到 $BACKEND_DIR 目录，请确保你在项目根目录运行此脚本。"
    exit 1
fi

# 2. 确保依赖完整 (防止 go mod 报错)
if [ ! -f "go.mod" ]; then
     echo "❌ 错误：在当前目录下找不到 go.mod 文件"
     exit 1
fi

# 3. 创建输出目录
mkdir -p "$OUTPUT_DIR"

# --- 开始编译 ---

# 1. Windows
echo "📦 Building for Windows (amd64)..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "$OUTPUT_DIR/${APP_NAME}-windows.exe" $ENTRY_FILE

# 2. Linux
echo "🐧 Building for Linux (amd64)..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "$OUTPUT_DIR/${APP_NAME}-linux" $ENTRY_FILE

# 3. macOS (Intel)
echo "🍎 Building for macOS (Intel)..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o "$OUTPUT_DIR/${APP_NAME}-darwin-amd64" $ENTRY_FILE

# 4. macOS (Apple Silicon)
echo "🍎 Building for macOS (Apple Silicon)..."
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o "$OUTPUT_DIR/${APP_NAME}-darwin-arm64" $ENTRY_FILE

# 返回根目录
cd ..

echo "-----------------------------------"
echo "✅ 构建完成！文件已生成在 bin 目录："
ls -lh "$OUTPUT_DIR"