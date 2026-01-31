package blockchain

import "goBlockchain/block"

//import "github.com/Amlan2006/go-blockchain/block"

type Blockchain struct {
	Blocks []block.Block
}

// create blockchain with genesis block
func NewBlockchain() *Blockchain {
	genesis := block.NewGenesisBlock()
	return &Blockchain{
		Blocks: []block.Block{genesis},
	}
}

// add a block
func (bc *Blockchain) AddBlock(data string) bool {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(lastBlock, data)
	if block.ValidateBlock(newBlock, lastBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
		return true
	}
	return false
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if !block.ValidateBlock(bc.Blocks[i], bc.Blocks[i-1]) {
			return false
		}
	}
	return true
}
