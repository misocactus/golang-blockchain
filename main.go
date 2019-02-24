package main

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:] //you can't just write hash, need to specify
	//because we already defined it as []byte which is different
}

func CreateBlock(data string, prevHash []byte) *Block { //* to say use this block
	block := &Block{[]byte{}}, []byte(data), prevHash}
	block.DeriveHash
	return block
}

func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.block[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main(){
	chain := InitBlockChain

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.block {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %S\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}

}
