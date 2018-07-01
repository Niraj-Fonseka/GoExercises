package main

/*
Implementing Blockchain
*/

//Blockchain is a database with a certian stucture , an ordered back linked list
//Blocks are stired the insertion order
//Each block is linked to the previos one
//Using the hash we are able to get the previous block quickly

type BlockChain struct {
	blocks []*Block
}

//Add blocks
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]

	newBlock := NewBlock(data, prevBlock.Hash)

	bc.blocks = append(bc.blocks, newBlock)
}

//To add block there has to be at least one block in the block chain
//Genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//Creating a new blockchain
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
