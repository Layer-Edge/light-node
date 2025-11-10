package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Layer-Edge/light-node/clients"
	"github.com/Layer-Edge/light-node/node"
	"github.com/Layer-Edge/light-node/utils"
	"github.com/joho/godotenv"
)

/*
	superProofClient := clients.MerkleTreeClient{}
	superProofContractAddress := utils.GetEnv("SUPER_PROOF_CONTRACT_ADDRESS", "0x7f9ebd5da0829bdb6cdff45bba0d68c0cae39364")
	err := superProofClient.Init(superProofContractAddress)
	if err != nil {
		log.Fatalf("failed to initialize super proof client: %v", err)
	}
	defer superProofClient.Close()
*/

func Worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	merkleTreeClient := clients.MerkleTreeClient{}
	merkleTreeStorageContractAddress := utils.GetEnv("MERKLE_TREE_STORAGE_CONTRACT_ADDRESS", "0x57bEBD1b51dF9F221F763a708591c1E72D2379D6")
	err := merkleTreeClient.Init(merkleTreeStorageContractAddress)
	if err != nil {
		log.Fatalf("failed to initialize cosmos query client: %v", err)
	}
	defer merkleTreeClient.Close()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is shutting down\n", id)
			return
		default:
			fmt.Printf("Worker %d is running...\n", id)
			node.CollectSampleAndVerify(&merkleTreeClient, "base proof")
			time.Sleep(5 * time.Second)
		}
	}
}

func SuperProofWorker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	superProofClient := clients.MerkleTreeClient{}
	superProofContractAddress := utils.GetEnv("SUPER_PROOF_CONTRACT_ADDRESS", "0x7f9ebd5da0829bdb6cdff45bba0d68c0cae39364")
	err := superProofClient.Init(superProofContractAddress)
	if err != nil {
		log.Fatalf("failed to initialize super proof client: %v", err)
	}
	defer superProofClient.Close()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Super Proof Worker %d is shutting down\n", id)
			return
		default:
			fmt.Printf("Super Proof Worker %d is running...\n", id)
			node.CollectSampleAndVerify(&superProofClient, "super proof")
			time.Sleep(5 * time.Second)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pubKey, err := utils.GetCompressedPublicKey()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Printf("Compressed Public Key: %s", pubKey)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGABRT, syscall.SIGTERM)

	log.Printf("Starting 2 worker(s)...")

	wg.Add(2)

	go Worker(ctx, &wg, 1)

	go SuperProofWorker(ctx, &wg, 1)

	<-signalChan
	fmt.Println("\nReceived interrupt signal. Shutting down gracefully...")

	cancel()

	wg.Wait()
	fmt.Println("All workers have shut down. Exiting..")
}
