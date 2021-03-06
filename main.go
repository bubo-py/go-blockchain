package main

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.Hash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Genesis block
func FirstBlock() *Block {
	return CreateBlock("A Genesis block", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{FirstBlock()}}
}

func main() {
	bc := InitBlockchain()

	bc.AddBlock("Block 1 after Genesis")
	bc.AddBlock("Block 2 after Genesis")

}
