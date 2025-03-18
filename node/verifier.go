package node

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Layer-Edge/light-node/clients"
	"github.com/Layer-Edge/light-node/utils"
)

type ZKProverPayload struct {
	Operation    string   `json:"operation"`
	Data         []string `json:"data"`
	ProofRequest *string  `json:"proof_request"` // Using interface{} since it can be null
	Proof        *Proof   `json:"proof"`
}

type ZKProverResponse struct {
	Root          string             `json:"root"`
	Proof         *Proof             `json:"proof"`
	Verified      bool               `json:"verified"`
	Visualization *TreeVisualization `json:"visualization"`
	Receipt       string             `json:"receipt"`
}

type TreeVisualization struct {
	LevelSizes        []uint64   `json:"level_sizes"`
	TreeStructure     [][]string `json:"tree_structure"`
	DataToHashMapping [][]string `json:"data_to_hash_mapping"`
}

type Proof struct {
	LeafValue string          `json:"leaf_value"`
	ProofPath [][]interface{} `json:"proof_path"` // Each element is [string, bool]
}

// TreeState stores the state of each merkle tree
type TreeState struct {
	LastRoot        string    // Last known root hash
	SleepUntil      time.Time // Time until which the tree should sleep
	ConsecutiveSame int       // Counter for consecutive same root occurrences
}

type SubmitProofRequest struct {
	WalletAddress string `json:"walletAddress"`
	Sign          string `json:"sign"`
	Timestamp     string `json:"timestamp"`
	Proof         string `json:"proof"`
	ProofHash     string `json:"proofHash"`
	Receipt       string `json:"receipt"`
}

var zkProverURL = utils.GetEnv("ZK_PROVER_URL", "http://127.0.0.1:3001")
var lightNodePointsAPI = utils.GetEnv("POINTS_API", "http://127.0.0.1:3001")

// Global map to track tree states with mutex for thread safety
var (
	treeStates = make(map[string]*TreeState)
	stateMutex sync.Mutex
)

func proveProof(data []string, proof_request string) (*Proof, error) {
	resp, err := clients.PostRequest[ZKProverPayload, ZKProverResponse](zkProverURL+"/process", ZKProverPayload{
		Operation:    "prove",
		Data:         data,
		ProofRequest: &proof_request,
		Proof:        nil,
	})
	if err != nil {
		return nil, fmt.Errorf("proof verification error: %v", err)
	}
	log.Printf("verification done: %v", resp)
	return resp.Proof, nil
}

func verifyProofs(data []string, proof Proof) (*string, *string, error) {
	resp, err := clients.PostRequest[ZKProverPayload, ZKProverResponse](zkProverURL+"/process", ZKProverPayload{
		Operation:    "verify",
		Data:         data,
		ProofRequest: nil,
		Proof:        &proof,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("proof verification error: %v", err)
	}
	log.Printf("verification done: %v\n", resp)
	return &resp.Receipt, &resp.Root, nil
}

func CollectSampleAndVerify() {
	cosmosQueryClient := clients.CosmosQueryClient{}
	err := cosmosQueryClient.Init()
	if err != nil {
		log.Fatalf("failed to initialize cosmos query client: %v", err)
	}
	defer cosmosQueryClient.Close()

	treeIds, err := cosmosQueryClient.ListMerkleTreeIds()
	if err != nil {
		log.Fatalf("failed to fetch tree ids: %v", err)
	}

	if len(treeIds) == 0 {
		log.Println("No trees available")
		return
	}

	// Try to find an available tree
	var activeTreeFound bool
	for _, treeId := range treeIds {
		// Skip trees that are sleeping
		stateMutex.Lock()
		state, exists := treeStates[treeId]
		if exists && time.Now().Before(state.SleepUntil) {
			log.Printf("Tree %s is sleeping until %s, skipping", treeId, state.SleepUntil.Format(time.RFC3339))
			stateMutex.Unlock()
			continue
		}
		stateMutex.Unlock()

		// Get tree data
		tree, err := cosmosQueryClient.GetMerkleTreeData(treeId)
		if err != nil {
			log.Printf("failed to fetch tree data for %s: %v", treeId, err)
			continue
		}

		// Check if root has changed
		stateMutex.Lock()
		if !exists {
			// First time seeing this tree
			treeStates[treeId] = &TreeState{
				LastRoot:        tree.Root,
				ConsecutiveSame: 0,
			}
			stateMutex.Unlock()
		} else if state.LastRoot == tree.Root {
			// Root hasn't changed - increment counter and potentially sleep
			state.ConsecutiveSame++

			if state.ConsecutiveSame >= 3 {
				// After 3 consecutive same roots, put the tree to sleep for 5 minutes
				sleepDuration := 5 * time.Minute
				state.SleepUntil = time.Now().Add(sleepDuration)
				log.Printf("Tree %s has had the same root %s for %d checks, putting to sleep for %v",
					treeId, tree.Root, state.ConsecutiveSame, sleepDuration)
				stateMutex.Unlock()
				continue
			}
			stateMutex.Unlock()
		} else {
			// Root has changed - reset counter
			state.LastRoot = tree.Root
			state.ConsecutiveSame = 0
			stateMutex.Unlock()
		}

		// Proceed with this tree
		sample := utils.RandomElement[string](tree.Leaves)

		// Track if verification was successful
		verificationSuccessful := false

		proof, err := proveProof(tree.Leaves, sample)
		if err != nil {
			log.Printf("failed to prove sample for tree %s: %v", treeId, err)
			// Continue to the next tree if proving fails
			continue
		}

		receipt, rootHash, err := verifyProofs(tree.Leaves, *proof)
		if err != nil {
			log.Printf("failed to verify sample for tree %s: %v", treeId, err)
			// Continue to the next tree if verification fails
			continue
		}
		if receipt != nil {
			walletAddress, err := utils.GetWalletAddress()
			if err != nil {
				log.Fatalf("failed to get wallet address from private key: %v", err)
			}

			timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
			msg := fmt.Sprintf("Submitting proof verification by %s of %s at %s", *walletAddress, proof.LeafValue, timestamp)
			fmt.Printf("Signing Message %s\n", msg)
			signature, err := utils.SignMessage(msg)
			if err != nil {
				log.Fatalf("failed to sign message: %v", err)
			}

			err = SubmitVerifiedProof(*walletAddress, *signature, *proof, *receipt, timestamp)
			if err != nil {
				log.Printf("Failed to submit verified proof: %v", err)
				// Continue to the next tree if submission fails
				continue
			} else {
				log.Printf("Successfully submitted verified proof for tree %s", treeId)
				verificationSuccessful = true
			}

			// Update the tree state with the verified root
			stateMutex.Lock()
			if state, exists := treeStates[treeId]; exists {
				if state.LastRoot != *rootHash {
					state.LastRoot = *rootHash
					state.ConsecutiveSame = 0
				}
			}
			stateMutex.Unlock()

			log.Printf("Tree %s - Sample Data %v verified with receipt %v\n", treeId, sample, *receipt)
		} else {
			log.Printf("Tree %s - Verification failed: missing receipt or root hash", treeId)
			continue
		}

		// Only mark as complete if verification was successful
		if verificationSuccessful {
			activeTreeFound = true
			break
		}
	}

	if !activeTreeFound {
		log.Println("No active trees available for verification or all verification attempts failed")
	}
}

// Helper function to get sleeping trees
func GetSleepingTrees() []string {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	var sleepingTrees []string
	now := time.Now()

	for treeId, state := range treeStates {
		if now.Before(state.SleepUntil) {
			sleepingTrees = append(sleepingTrees, treeId)
		}
	}

	return sleepingTrees
}

func SubmitVerifiedProof(walletAddress string, signature string, proof Proof, receipt string, timestamp string) error {
	// Create the proof hash (this appears to be required by the API)
	// Note: You may need to adjust how proofHash is calculated based on your requirements
	proofHash := utils.HashString(proof.LeafValue) // Assuming utils.HashString exists

	requestBody := SubmitProofRequest{
		WalletAddress: walletAddress,
		Sign:          signature,
		Timestamp:     timestamp,
		Proof:         proof.LeafValue,
		ProofHash:     proofHash,
		Receipt:       receipt,
	}

	// Make the API request
	// You may need to adjust the URL based on your environment
	resp, err := clients.PostRequest[SubmitProofRequest, map[string]interface{}](
		lightNodePointsAPI+"/api/cli-node/submit-verified-proof",
		requestBody,
	)

	if err != nil {
		return fmt.Errorf("failed to submit verified proof: %v", err)
	}

	log.Printf("Proof submission result: %v", resp)
	return nil
}
