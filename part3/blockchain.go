package blockchain

import (
	"fmt"
	"log"

	collection "mycoin/database/collections"

	"go.mongodb.org/mongo-driver/bson"
)

const collectionName = "blocks"

var blockCollection = collection.Collection{collection.GetCollection(collectionName)}

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip []byte
	db  collection.Collection
}

// BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte
	db          collection.Collection
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	lastBlock := blockCollection.GetLastRecord()
	fmt.Println("lastBlock", lastBlock[0])

	var lb Block
	err := bson.Unmarshal(lastBlock[0], &lb)
	if err != nil {
		log.Fatal("detail:", err)
	}
	lastHash := lb.Hash
	fmt.Println("lastHash", lastHash)
	newBlock := NewBlock(data, lastHash)
	_, err1 := blockCollection.CreateByLambda(newBlock)
	if err1 != nil {
		log.Fatal("Create failed", err1)
	}
	(*bc).tip = newBlock.Hash

}

// Iterator ...
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	var block Block

	var Conditions struct {
		CurrentHash []byte `bson:"hash"`
	}
	Conditions.CurrentHash = i.currentHash

	//GET USER FROM DB
	bl := blockCollection.FindByLambda(Conditions)
	if len(bl) == 0 {
		log.Fatal("block is not existed: ", Conditions.CurrentHash)
	}

	err := bson.Unmarshal(bl[0], &block)
	if err != nil {
		log.Fatal("detail:", err)
	}

	i.currentHash = block.PrevBlockHash

	return &block
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	var bc Blockchain
	lastBlock := blockCollection.GetLastRecord()
	if len(lastBlock) == 0 {
		genesis := NewGenesisBlock()
		_, err := blockCollection.CreateByLambda(genesis)
		if err != nil {
			log.Fatal("Create failed", err)
		}
		// fmt.Println("bl    ", bl)

		bc = Blockchain{genesis.Hash, blockCollection}
	} else {
		var block Block

		err := bson.Unmarshal(lastBlock[0], &block)
		if err != nil {
			log.Fatal("detail:", err)
		}
		// fmt.Println("bl    ", block)
		bc = Blockchain{block.Hash, blockCollection}

	}

	return &bc
}
