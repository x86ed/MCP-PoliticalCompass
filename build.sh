#!/bin/bash

# Build script for cross-compiling MCP-PoliticalCompass to multiple architectures
# Usage: ./build.sh [version]

set -e

VERSION=${1:-$(cat VERSION 2>/dev/null || echo "dev")}
BINARY_NAME="mcp-political-compass"
BUILD_DIR="dist"

echo "Building MCP-PoliticalCompass v${VERSION} for multiple architectures..."

# Clean previous builds
rm -rf ${BUILD_DIR}
mkdir -p ${BUILD_DIR}

echo "Created build directory: ${BUILD_DIR}"

# Define target platforms
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64" 
    "darwin/arm64"
    "windows/amd64"
)

echo "Building for ${#PLATFORMS[@]} platforms..."

# Build for each platform
for PLATFORM in "${PLATFORMS[@]}"; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    
    OUTPUT_NAME="${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}"
    
    # Add .exe extension for Windows
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME="${OUTPUT_NAME}.exe"
    fi
    
    echo "Building for ${GOOS}/${GOARCH}..."
    
    # Build with ldflags to embed version info
    env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build \
        -ldflags "-s -w -X main.Version=${VERSION}" \
        -o "${BUILD_DIR}/${OUTPUT_NAME}" \
        .
    
    if [ -f "${BUILD_DIR}/${OUTPUT_NAME}" ]; then
        echo "✓ Built ${OUTPUT_NAME}"
        
        # Create compressed archive
        if [ "$GOOS" = "windows" ]; then
            # Create zip for Windows
            (cd ${BUILD_DIR} && zip -q "${OUTPUT_NAME%.exe}.zip" "${OUTPUT_NAME}" && rm "${OUTPUT_NAME}")
        else
            # Create tar.gz for Unix-like systems
            (cd ${BUILD_DIR} && tar -czf "${OUTPUT_NAME}.tar.gz" "${OUTPUT_NAME}" && rm "${OUTPUT_NAME}")
        fi
    else
        echo "✗ Failed to build ${OUTPUT_NAME}"
    fi
done

# Create checksums
echo "Generating checksums..."
if command -v sha256sum >/dev/null 2>&1; then
    (cd ${BUILD_DIR} && sha256sum * > checksums.txt)
else
    (cd ${BUILD_DIR} && shasum -a 256 * > checksums.txt)
fi

echo ""
echo "Build complete! Artifacts available in ${BUILD_DIR}/"
echo "Checksums generated in ${BUILD_DIR}/checksums.txt"
echo ""
ls -la ${BUILD_DIR}/
