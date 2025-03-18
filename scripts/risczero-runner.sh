#!/bin/bash
# start-risc0-service.sh
# This script builds and starts the ZK prover service

# Configuration
ZK_PROVER_DIR="risc0-merkle-service"
LOG_FILE="zk-prover.log"

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

# Start ZK prover service
log "Starting ZK prover service..."
cd "$ZK_PROVER_DIR" || error "Failed to navigate to ZK prover directory"

# Start the service
cargo run > "$LOG_FILE" 2>&1 &
ZK_PROVER_PID=$!

# Store PID to a file for other scripts to read
echo "$ZK_PROVER_PID" > "../zk_prover_pid.txt"

# Wait for ZK prover to start
log "Waiting for ZK prover to initialize (5 seconds)..."
sleep 5

# Check if ZK prover is running
if ! ps -p $ZK_PROVER_PID > /dev/null; then
  error "ZK prover failed to start. Check $LOG_FILE for details."
fi

success "ZK prover started successfully with PID: $ZK_PROVER_PID"
success "ZK prover is now running and ready to accept connections"