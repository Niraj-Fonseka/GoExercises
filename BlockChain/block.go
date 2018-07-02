package main

import (
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int //Nonce is saved to the block because its required to veryfy
}

// func (b *Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

// 	//Getting the block fields and concatinating them to create the header and calculate the sha hash on the combination
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)

// 	b.Hash = hash[:]
// }

//Creation of a block
func NewBlock(data string, PrevBlockHash []byte) *Block {
	// block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}}
	// block.SetHash()
	// return block

	block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
