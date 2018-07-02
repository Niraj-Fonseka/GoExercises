package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

//need to have a target that takes less than 256 bits in mem
//the difference should be significant enough ( too big -> hard to hash , too small -> easy to hash)
const targetBits = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)                  //Initialize big int with value of 1
	target.Lsh(target, uint(256-targetBits)) //Shift it by left 256 0 target bigs
	//Hex representation of the target = 0x10000000000000000000000000000000000000000000000000000000000

	//Target is the upper boundry , if a number ( hash ) is lower than the boundary its valid.
	//Lowering the boundary will result in fewer valid numbers --> more work required

	pow := &ProofOfWork{b, target}

	return pow
}

// func IntToHex(n int64) []byte {
// 	return []byte(strconv.FormatInt(n, 16))
// }

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

//Prepare data for hashing
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)), //Nonce is the counter
		},
		[]byte{},
	)
	return data
}

//Core POW Algo

/*
 Initializing the variable hashInt is the int representation of hash ,
 Nonce is the counter
 run an infinite loop limited by maxNonce which == math.MaxInt64
 this is done to avoid a possibe overflow of nonce

 -- Inside the loop
 prepare data
 hash with sha-256
 conver the hash to a big integer
 compare the ints with the target
*/
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte

	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n\n")

	return nonce, hash[:]
}

//Validating proof of works
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
