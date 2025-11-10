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
    curl -L https://risczero.com/install | bash && rzup install
    ```

### Build Instructions

Configure environment variables or in .env file
```env
RPC_URL=https://rpc.layeredge.io
MERKLE_TREE_STORAGE_CONTRACT_ADDRESS=0x57bEBD1b51dF9F221F763a708591c1E72D2379D6
SUPER_PROOF_CONTRACT_ADDRESS=0x7f9EBd5DA0829bdB6CDFF45BbA0D68c0cae39364
ZK_PROVER_URL=http://127.0.0.1:3000
API_REQUEST_TIMEOUT=100
POINTS_API=http://127.0.0.1:8080
PRIVATE_KEY='cli-node-private-key'
```

Please make sure the ZK Prover URL is the same URL as that of the server where the merkle service is running

## Run both the servers manually

```bash
cd risc0-merkle-service
cargo build && cargo run
```
Make sure risc0 server is up first
In a separate terminal (navigate to the root and run it):
```bash
go build
./light-node
```

Make sure both the servers are running independently

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

Need more help? Visit the FAQ & Troubleshooting Guide: https://docs.layeredge.io/introduction/developer-guide/run-a-node/faq-and-troubleshooting-guide

## License

This project is licensed under the MIT License - see the LICENSE file for details.
