package blockchain

import (
	"errors"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// Blockchain represents the Ethereum blockchain client
type Blockchain struct {
	client *rpc.Client
}

// NewBlockchain creates a new instance of Blockchain
func NewBlockchain(endpoint string) (*Blockchain, error) {
	// Connect to the Ethereum client
	client, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, errors.New("failed to connect to Ethereum client: " + err.Error())
	}
	return &Blockchain{client: client}, nil
}

// Close closes the connection to the Ethereum client
//func (b *Blockchain) Close() error {
//if b.client != nil {
//return b.client.Close()
//}
//return nil
//}

// GetBlockCount retrieves the current block count from the blockchain
func (b *Blockchain) GetBlockCount() (uint64, error) {
	var blockCount uint64
	err := b.client.Call(&blockCount, "eth_blockNumber")
	if err != nil {
		return 0, errors.New("failed to retrieve block count: " + err.Error())
	}
	return blockCount, nil
}

// GetBlockByNumber retrieves block information by block number
func (b *Blockchain) GetBlockByNumber(blockNumber uint64) (*types.Block, error) {
	var block types.Block
	hexNumber := strconv.FormatUint(blockNumber, 16) // Convert uint64 to hexadecimal string
	err := b.client.Call(&block, "eth_getBlockByNumber", hexNumber, true)
	if err != nil {
		return nil, errors.New("failed to retrieve block by number: " + err.Error())
	}
	return &block, nil
}
