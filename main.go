package main

import (
	"fmt"
	"goBlockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 10 BTC to Alice")
	bc.AddBlock("Send 5 BTC to Bob")
	//s := "hjikei"
	//bc.Blocks[1].Data = s

	for _, block := range bc.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Data:", block.Data)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("PrevHash:", block.PrevHash)
		fmt.Println("----------------------")
	}
	fmt.Println("Blockchain valid?", bc.IsValid())

}
