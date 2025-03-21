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

	"github.com/Layer-Edge/light-node/node"
	"github.com/Layer-Edge/light-node/utils"
	"github.com/joho/godotenv"
)

func Worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is shutting down\n", id)
			return
		default:
			fmt.Printf("Worker %d is running...\n", id)
			node.CollectSampleAndVerify()
			time.Sleep(5 * time.Second)
		}
	}
}

func main() {
	log.Println("Starting light-node...")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println(".env loaded successfully")
	}

	pubKey, err := utils.GetCompressedPublicKey()
	if err != nil {
		log.Fatalf("Error getting compressed public key: %v", err)
	}
	log.Printf("Compressed Public Key: %s", pubKey)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)

	wg.Add(1)
	go Worker(ctx, &wg, 1)

	<-signalChan
	fmt.Println("\nReceived interrupt signal. Shutting down gracefully...")

	cancel()
	wg.Wait()
	fmt.Println("Worker has shut down. Exiting..")
}
