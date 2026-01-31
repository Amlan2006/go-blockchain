package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index      int
	Timestamp  string
	Hash       string
	PrevHash   string
	Data       string
	Nonce      int
	Difficulty int
}

func calculateHash(block Block) string {
	record := fmt.Sprintf(
		"%d%s%s%s%d",
		block.Index,
		block.Timestamp,
		block.Data,
		block.PrevHash,
		block.Nonce,
	)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func NewGenesisBlock() Block {
	block := Block{
		Index:      0,
		Timestamp:  time.Now().String(),
		Data:       "Genesis Block",
		PrevHash:   "",
		Difficulty: 3,
	}
	block.Hash = calculateHash(block)
	MineBlock(&block)
	return block
}

func NewBlock(prevBlock Block, data string) Block {
	block := Block{
		Index:      prevBlock.Index + 1,
		Timestamp:  time.Now().String(),
		Data:       data,
		PrevHash:   prevBlock.Hash,
		Difficulty: prevBlock.Difficulty,
	}
	block.Hash = calculateHash(block)
	MineBlock(&block)
	return block
}

func ValidateBlock(newBlock Block, prevBlock Block) bool {
	if newBlock.Index != prevBlock.Index+1 {
		return false
	}

	// 2. Previous hash must match
	if newBlock.PrevHash != prevBlock.Hash {
		return false
	}

	target := strings.Repeat("0", newBlock.Difficulty)

	if !strings.HasPrefix(newBlock.Hash, target) {
		return false
	}

	// 3. Hash must be correct
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func MineBlock(b *Block) {
	target := strings.Repeat("0", b.Difficulty)

	for {
		hash := calculateHash(*b)

		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		}

		b.Nonce++
	}
}
