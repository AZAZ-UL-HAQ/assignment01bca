// Azaz-Ul-Haq
// 19I-2000
// Cys-M
// BLockChain A#01
package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Tranaction   string
	Nonce        int
	PreviousHash string
	BlockHash    string
}

type Chain struct {
	ChainBlock []Block
	ChainHash  string
}

func NewBlock(transaction string, Nonce int, PreviousHash string) *Block {
	s := new(Block)
	s.Tranaction = transaction
	s.Nonce = Nonce
	s.PreviousHash = PreviousHash
	s.CalculateHash(s.Tranaction + string(s.Nonce) + s.PreviousHash)
	return s
}

func (b *Block) CalculateHash(stringToHash string) {

	sum := sha256.Sum256([]byte(stringToHash))
	b.BlockHash = hex.EncodeToString(sum[:])

}

func (b *Chain) ListBlocks() {
	//A method to print all the blocks in a nice format showing block data such
	//as transaction, Nonce, previous hash, current block hash

	for i := range b.ChainBlock {
		fmt.Printf(" Tranaction is : %s \n", b.ChainBlock[i].Tranaction)
		fmt.Printf(" Nonce is : %d  \n ", (b.ChainBlock[i].Nonce))
		fmt.Println("\n Previous Hash is :    ", (b.ChainBlock[i].PreviousHash))
		//fmt.Printf("%s :\n ",b.ChainBlock[i].BlockHash)
		fmt.Println("\n  BlockHash is :  ", b.ChainBlock[i].BlockHash)

	}
}
func (b *Chain) ChangeBlock() {
	//	function to change block transaction of the given block ref

	b.ChainBlock[1].Nonce = 4
}

func (b *Chain) VerifyChain() {

	b.ChainBlock[0].CalculateHash(b.ChainBlock[0].Tranaction + string(b.ChainBlock[0].Nonce) + b.ChainBlock[0].PreviousHash)

	for i := 1; i < len(b.ChainBlock); i++ {
		b.ChainBlock[i].CalculateHash(b.ChainBlock[i].Tranaction + string(b.ChainBlock[i].Nonce) + b.ChainBlock[i-1].BlockHash)
	}

	if b.ChainHash == b.ChainBlock[len(b.ChainBlock)-1].BlockHash {
		fmt.Printf("\n Block Chain is verified\n\n")
	} else {
		fmt.Printf("\n Block Chain is modofied\n\n")
	}

}

//	func CalculateHash (stringToHash string) {
//function for calculating hash of a block
