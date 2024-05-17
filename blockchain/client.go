package blockchain

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

// CustomBlockType defines a structure for the desired block data
type CustomBlockType struct {
	Hash   common.Hash `json:"hash"`
	Number uint64      `json:"number"`
}

// GetLatestBlock retrieves the latest block from the blockchain
func (b *Blockchain) GetLatestBlock() (*CustomBlockType, error) {
	blockNumber, err := b.GetBlockCount()
	if err != nil {
		return nil, errors.New("failed to get latest block: " + err.Error())
	}

	latestBlock, err := b.GetBlockByNumber(blockNumber)
	if err != nil {
		return nil, errors.New("failed to get block by number: " + err.Error())
	}

	// Convert relevant data from types.Block to CustomBlockType
	return &CustomBlockType{
		Hash:   latestBlock.Hash(),
		Number: latestBlock.NumberU64(),
	}, nil
}
