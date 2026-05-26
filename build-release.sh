#!/bin/bash
#
# S-UI Release Build Script
# Builds the s-ui binary and packages it for GitHub Release
#
# Usage:
#   ./build-release.sh                   # Build for current arch (no naiveproxy)
#   ./build-release.sh --naive           # Build with naiveproxy/cronet support
#   ./build-release.sh --all             # Build with naiveproxy + download libcronet
#   ./build-release.sh --docker          # Build for all platforms via Docker
#   VERSION=v1.4.2 ./build-release.sh    # Set version label

set -e
cd "$(dirname "$0")"

# --- config ---
VERSION="${VERSION:-v1.4.1}"
WITH_NAIVE="${WITH_NAIVE:-false}"
FULL_CRONET="${FULL_CRONET:-false}"
CURRENT_ARCH="${CURRENT_ARCH:-}"

# Detect arch
detect_arch() {
    case "$(uname -m)" in
        x86_64|amd64) echo "amd64" ;;
        aarch64|arm64) echo "arm64" ;;
        armv7l|armv7)  echo "armv7" ;;
        armv6l)         echo "armv6" ;;
        i686|i386)      echo "386" ;;
        s390x)          echo "s390x" ;;
        *) echo "unknown"; exit 1 ;;
    esac
}

# --- parse args ---
for arg in "$@"; do
    case "$arg" in
        --naive)   WITH_NAIVE="true" ;;
        --all)     WITH_NAIVE="true"; FULL_CRONET="true" ;;
        --docker)  echo ">> Building via Docker for all platforms..."
                   docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7,linux/arm/v6,linux/386 \
                       -f Dockerfile --progress=plain . 2>&1 | tee docker-build.log
                   echo ">> Done. See docker-build.log"
                   exit 0 ;;
        --help|-h) echo "Usage: $0 [--naive] [--all] [--docker]"
                   echo "  --naive    Include naiveproxy outbound"
                   echo "  --all      Include naiveproxy + download libcronet.so"
                   echo "  --docker   Build multi-platform via Docker"
                   exit 0 ;;
    esac
done

ARCH="${CURRENT_ARCH:-$(detect_arch)}"
PACKAGE="s-ui-linux-${ARCH}"
echo "==> Building S-UI ${VERSION} for linux/${ARCH}"

# --- 1. build frontend ---
echo ""
echo "==> [1/3] Building frontend..."
cd frontend
npm ci --silent 2>/dev/null || npm install --silent
npm run build
cd ..

# --- 2. copy frontend dist to web/html (embedded at build time) ---
echo ""
echo "==> Copying frontend dist to web/html..."
rm -rf web/html
mkdir -p web/html
cp -R frontend/dist/* web/html/

# --- 3. build Go binary ---
echo ""
echo "==> [2/3] Building Go binary..."

BASE_TAGS="with_quic,with_grpc,with_utls,with_acme,with_gvisor,badlinkname,tfogo_checklinkname0,with_tailscale"

if [ "$WITH_NAIVE" = "true" ]; then
    BUILD_TAGS="${BASE_TAGS},with_naive_outbound,with_musl"
    LDFLAGS="-w -s -checklinkname=0"
    echo ">> Tags: ${BUILD_TAGS} (with naiveproxy)"
else
    BUILD_TAGS="${BASE_TAGS}"
    LDFLAGS="-w -s -checklinkname=0"
    echo ">> Tags: ${BUILD_TAGS}"
fi

# Determine if we can link statically (Linux only)
if [ "$(uname -s)" = "Linux" ]; then
    LDFLAGS="${LDFLAGS} -linkmode external -extldflags '-static'"
    echo ">> Linking: static"
fi

go build -ldflags="${LDFLAGS}" -tags "${BUILD_TAGS}" -o sui main.go

echo ">> Binary type: $(file sui | cut -d: -f2)"
ldd sui 2>/dev/null && echo ">> Dynamic binary" || echo ">> Static binary"

# --- 4. download libcronet if requested ---
if [ "$FULL_CRONET" = "true" ]; then
    echo ""
    echo "==> [3/3] Downloading libcronet.so..."
    LIBCRONET_URL="https://github.com/SagerNet/cronet-go/releases/latest/download/libcronet-linux-${ARCH}.so"
    curl -fL -o libcronet.so "$LIBCRONET_URL"
    chmod 755 libcronet.so
    echo ">> Downloaded libcronet.so"
else
    echo ""
    echo "==> [3/3] Skipped libcronet download (use --all to include it)"
fi

# --- 5. package ---
echo ""
echo "==> Packaging ${PACKAGE}.tar.gz..."
rm -rf s-ui/ ${PACKAGE}.tar.gz
mkdir -p s-ui

cp sui s-ui/
cp s-ui.service s-ui/
cp s-ui.sh s-ui/
[ -f libcronet.so ] && cp libcronet.so s-ui/

tar -zcvf "${PACKAGE}.tar.gz" s-ui

# --- 6. summary ---
echo ""
echo "============================================"
echo "  Build complete: ${PACKAGE}.tar.gz"
echo "  Version: ${VERSION}"
echo "  Arch:    ${ARCH}"
echo "  Size:    $(du -h ${PACKAGE}.tar.gz | cut -f1)"
echo ""
echo "  Contents:"
tar -tzf "${PACKAGE}.tar.gz"
echo "============================================"
echo ""
echo "To create a GitHub Release, upload this file to:"
echo "  https://github.com/sayeneBB/s-ui/releases/new?tag=${VERSION}"
echo ""
echo "Or use the gh CLI:"
echo "  gh release create ${VERSION} ${PACKAGE}.tar.gz --title \"s-ui ${VERSION}\" --notes \"Release ${VERSION}\""
