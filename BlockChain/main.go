package main

import (
	"fmt"
	"strconv"
)

//In a more complex blockchain you will have to do more complex computations befoer adding a block - Proof of work
//Distributed database without any single decision maker
//The new added blocks must be approved by other participants of the network ( concensus )
func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev Hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s \n", block.Data)
		fmt.Printf("Hash : %x\n ", block.Hash)
		fmt.Println()

		pow := NewProofOfWork(block)
		fmt.Printf("PoW : %s \n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
