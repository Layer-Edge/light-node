#!/bin/bash
# run-all.sh
# This script builds and runs both the ZK prover service and the light node client

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

# Handle exit - cleanup
cleanup() {
  log "Shutting down services..."
  
  # Check if PID file exists
  if [ -f "zk_prover_pid.txt" ]; then
    ZK_PROVER_PID=$(cat zk_prover_pid.txt)
    kill $ZK_PROVER_PID 2>/dev/null
    log "ZK prover service stopped (PID: $ZK_PROVER_PID)"
    rm zk_prover_pid.txt
  fi
  
  log "Cleanup complete"
}

# Set trap to catch script termination
trap cleanup EXIT INT TERM

# Build RISC0 Merkle Service
log "Building RISC0 Merkle Service..."
scripts/build-risczero.sh
if [ $? -ne 0 ]; then
  error "Failed to build RISC0 Merkle Service"
fi

# Build Light Node
log "Building Light Node Client..."
scripts/build-light-node.sh
if [ $? -ne 0 ]; then
  error "Failed to build Light Node Client"
fi

# Start RISC0 Merkle Service
log "Starting RISC0 Merkle Service..."
scripts/risczero-runner.sh
if [ $? -ne 0 ]; then
  error "Failed to start RISC0 Merkle Service"
fi

# Start Light Node
log "Starting Light Node Client..."
scripts/light-node-runner.sh

success "All services have been shut down"
exit 0