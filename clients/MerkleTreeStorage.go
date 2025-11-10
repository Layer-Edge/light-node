package clients

import (
	"fmt"
	"math/big"

	"github.com/Layer-Edge/light-node/contracts"
	"github.com/Layer-Edge/light-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MerkleTreeData struct {
	Root      common.Hash
	Owner     common.Address
	LeafCount *big.Int
	CreatedAt *big.Int
	Leaves    []string
}

type MerkleTreeClient struct {
	rpc      *ethclient.Client
	contract *contracts.MerkleTreeStorage
}

func (mtc *MerkleTreeClient) Init(merkleTreeStorageContractAddress string) error {
	layerEdgeRPC := utils.GetEnv("RPC_URL", "https://rpc.layeredge.io")
	layerEdgeClient, err := ethclient.Dial(layerEdgeRPC)
	if err != nil {
		return fmt.Errorf("error creating layerEdgeClient: %v", err)
	}

	contractAddress := common.HexToAddress(merkleTreeStorageContractAddress)
	merkleTreeStorageContract, err := contracts.NewMerkleTreeStorage(contractAddress, layerEdgeClient)
	if err != nil {
		return fmt.Errorf("error creating merkleTreeStorageContract: %v", err)
	}

	mtc.rpc = layerEdgeClient
	mtc.contract = merkleTreeStorageContract

	return nil
}

func (mtc *MerkleTreeClient) GetAllRoots() ([]common.Hash, error) {
	roots, err := mtc.contract.GetAllRoots(nil)
	if err != nil {
		return nil, fmt.Errorf("error getting all roots: %v", err)
	}

	// Convert [][32]byte to []common.Hash
	hashRoots := make([]common.Hash, len(roots))
	for i, root := range roots {
		hashRoots[i] = common.BytesToHash(root[:])
	}

	return hashRoots, nil
}

func (mtc *MerkleTreeClient) GetTreeInfo(root common.Hash) (*MerkleTreeData, error) {
	treeInfo, err := mtc.contract.GetTreeInfo(nil, root)
	if err != nil {
		return nil, fmt.Errorf("error getting tree info: %v", err)
	}
	// Convert [][]byte to []string
	leaves := make([]string, len(treeInfo.Leaves))
	for i, leaf := range treeInfo.Leaves {
		leaves[i] = string(leaf)
	}

	return &MerkleTreeData{
		Root:      root,
		Owner:     treeInfo.Owner,
		LeafCount: treeInfo.LeafCount,
		CreatedAt: treeInfo.CreatedAt,
		Leaves:    leaves,
	}, nil
}

func (mtc *MerkleTreeClient) Close() {
	if mtc.rpc != nil {
		mtc.rpc.Close()
	}
}
