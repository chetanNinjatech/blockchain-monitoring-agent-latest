package analysis

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

// Define Transaction struct with added fields (adjust based on your data format)
type Transaction struct {
	From      string    `json:"from"`                // Sender address
	To        string    `json:"to"`                  // Recipient address
	Amount    int64     `json:"amount"`              // Transaction amount
	Hash      string    `json:"hash"`                // Added Hash field for transaction identification
	Fee       int64     `json:"fee,omitempty"`       // Fee associated with the transaction (if applicable)
	Timestamp time.Time `json:"timestamp,omitempty"` // Timestamp of the transaction (if applicable)
	// Add more fields as needed (e.g., signature)
}

type Block struct {
	Hash         string        `json:"hash"`         // Block hash
	Transactions []Transaction `json:"transactions"` // Transactions within the block
}

func DetectVulnerabilities(block *Block) {
	if block != nil {
		// Double-spending detection (example)
		transactionMap := make(map[string]bool) // Track seen transaction IDs
		for _, tx := range block.Transactions {
			if transactionMap[tx.Hash] {
				log.Println("Potential double-spending detected for transaction:", tx.Hash)
			} else {
				transactionMap[tx.Hash] = true
			}
		}

		// Additional checks:

		// Reentrancy attack (basic check) - Looks for transactions with the same sender and recipient within a short time window
		checkReentrancy(block)

		// 51% attack (basic check) - Looks for inconsistencies in the block hash chain (requires previous block hash)
		previousBlockHash := "" // Replace with actual logic to get the previous block hash
		check51PercentAttack(block, previousBlockHash)
	}
}

func checkReentrancy(block *Block) {
	senderMap := make(map[string]time.Time) // Track recent transactions by sender
	for _, tx := range block.Transactions {
		sender := tx.From
		previousTxTime, ok := senderMap[sender]
		if ok && time.Since(previousTxTime) < 5*time.Second { // Adjustable time window
			log.Printf("Transaction %s from %s might be part of a reentrancy attack (recent transaction)", tx.Hash, sender)
		}
		senderMap[sender] = tx.Timestamp
	}
}

func check51PercentAttack(block *Block, previousBlockHash string) {
	if previousBlockHash == "" {
		log.Println("51% attack check: Unable to determine previous block hash")
		return
	}

	// Calculate the expected hash based on the previous block hash and current block data
	expectedHash := calculateBlockHash(block, previousBlockHash)

	if expectedHash != block.Hash {
		log.Println("51% attack check: Potential inconsistency in block hash chain")
	}
}

func calculateBlockHash(block *Block, previousHash string) string {
	// This is a simplified example. Actual hash calculation might involve more fields and specific hashing algorithms used by the blockchain.
	data := fmt.Sprintf("%s-%s-%d", previousHash, block.Hash, len(block.Transactions))
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
