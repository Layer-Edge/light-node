#!/bin/bash
# build-risc0-service.sh
# This script builds the ZK prover service

# Configuration
ZK_PROVER_DIR="risc0-merkle-service"

# Color codes
BLUE='\033[0;34m'
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Log function
log() {
  echo -e "${BLUE}[$(date '+%Y-%m-%d %H:%M:%S')] $1${NC}"
}

error() {
  echo -e "${RED}[$(date '+%Y-%m-%d %H:%M:%S')] ERROR: $1${NC}"
  exit 1
}

success() {
  echo -e "${GREEN}[$(date '+%Y-%m-%d %H:%M:%S')] $1${NC}"
}

# Check if zk-prover directory exists
if [ ! -d "$ZK_PROVER_DIR" ]; then
  error "ZK prover directory not found. Please run this script from the project root."
fi

# Check if risc0 toolchain is installed
if ! command -v rzup &> /dev/null; then
  error "The 'risc0' toolchain could not be found. Install it using: curl -L https://risczero.com/install | bash && rzup install"
fi

# Build the service
log "Building RISC0 Merkle service..."
cd "$ZK_PROVER_DIR" || error "Failed to navigate to ZK prover directory"
cargo build --release || error "Failed to build RISC0 Merkle service"

success "RISC0 Merkle service built successfully"