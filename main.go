package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// STRUCTURE
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

// FUNCTIONS
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block`s relevant info with the new blocks
	hash := sha256.Sum256(info)
	// This performs the actual hashing algorithm
	b.Hash = hash[:]
	// If this ^ doesn`t make sense, you can look up slice defaults
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash}
	// If this is gibberish to you, look at pointer sintax
	// in Go

	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// All of our functions are relying on a prevHash so, in order to start our
// chain we need to generate our first block by creating a Genesis block and
// initializing it
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// Create a main function to initilize the block and print our blocks
func main() {

	chain := InitBlockchain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
	}
}
