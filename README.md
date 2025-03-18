# Layer Edge Light Node

## Introduction

Layer Edge Light Node is a client that connects to the Layer Edge network to verify Merkle trees by collecting random samples from available trees and verifying their integrity. The light node performs Zero-Knowledge proof verification operations through a local ZK prover service, and submits verified proofs to the network.

Key features:
- Automatically discovers available Merkle trees from the network
- Collects random samples from trees for verification
- Generates and verifies Zero-Knowledge proofs
- Implements intelligent sleep mechanisms to avoid redundant work on unchanged trees
- Submits verified proofs to earn rewards

## Building and Running

### Prerequisites

- Go 1.18 or higher
- Rust 1.81.0 or higher
- Access to a Layer Edge gRPC endpoint
- The 'risc0' toolchain could not be found.
  To install the risc0 toolchain, use rzup.
  For example:
    ```bash
    curl -L https://risczero.com/install | bash
    rzup install
    ```

### Build Instructions

Configure environment variables or in .env file
```env
GRPC_URL=34.31.74.109:9090
CONTRACT_ADDR=cosmos1ufs3tlq4umljk0qfe8k5ya0x6hpavn897u2cnf9k0en9jr7qarqqt56709
ZK_PROVER_URL=http://127.0.0.1:3001
API_REQUEST_TIMEOUT=100
POINTS_API=http://127.0.0.1:8080
PRIVATE_KEY='cli-node-private-key'
```

Grant execute permission and run the script
```bash
chmod +x scripts/*
scripts/runner.sh
```

## Logging and Monitoring

The light node provides detailed logging about its operations. You can monitor the log output to track:
- Tree discovery
- Proof generation and verification
- Submission of verified proofs
- Sleep state of trees

## Troubleshooting

If you encounter issues:

1. Check your gRPC connection to the Layer Edge network
2. Ensure the ZK prover service is running and accessible
3. Verify your wallet address and signature format
4. Check logs for specific error messages

## License

This project is licensed under the MIT License - see the LICENSE file for details.