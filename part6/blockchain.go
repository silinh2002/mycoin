package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	collection "mycoin/database/collections"

	"go.mongodb.org/mongo-driver/bson"
)

const genesisCoinbaseData = "The Times 05/May/2021 Chancellor on brink of second bailout for banks"

var blockCollection = collection.Collection{collection.GetCollection("blocks")}

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
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	lastBlock := blockCollection.GetLastRecord()
	fmt.Println("lastBlock", lastBlock[0])

	var lb Block
	err := bson.Unmarshal(lastBlock[0], &lb)
	if err != nil {
		log.Fatal("detail:", err)
	}
	lastHash := lb.Hash
	fmt.Println("lastHash", lastHash)
	newBlock := NewBlock(transactions, lastHash)
	_, err1 := blockCollection.CreateByLambda(newBlock)
	if err1 != nil {
		log.Fatal("Create failed", err1)
	}
	(*bc).tip = newBlock.Hash

}

// MineBlock mines a new block with the provided transactions
func (bc *Blockchain) MineBlock(transactions []*Transaction) *Block {
	for _, tx := range transactions {
		if bc.VerifyTransaction(tx) != true {
			log.Panic("ERROR: Invalid transaction")
		}
	}

	lastBlock := blockCollection.GetLastRecord()
	fmt.Println("lastBlock", lastBlock[0])

	var lb Block
	err := bson.Unmarshal(lastBlock[0], &lb)
	if err != nil {
		log.Fatal("detail:", err)
	}
	lastHash := lb.Hash
	fmt.Println("lastHash", lastHash)

	newBlock := NewBlock(transactions, lastHash)
	_, err1 := blockCollection.CreateByLambda(newBlock)
	if err1 != nil {
		log.Fatal("Create failed", err1)
	}
	(*bc).tip = newBlock.Hash

	bc.tip = newBlock.Hash
	return newBlock
}

// SignTransaction signs inputs of a Transaction
func (bc *Blockchain) SignTransaction(tx *Transaction, privKey ecdsa.PrivateKey) {
	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	tx.Sign(privKey, prevTXs)
}

// FindUnspentTransactions returns a list of transactions containing unspent outputs
func (bc *Blockchain) FindUnspentTransactions(pubKeyHash []byte) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOutIdx := range spentTXOs[txID] {
						if spentOutIdx == outIdx {
							continue Outputs
						}
					}
				}

				if out.IsLockedWithKey(pubKeyHash) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.UsesKey(pubKeyHash) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTXs
}

// FindUTXO finds all unspent transaction outputs and returns transactions with spent outputs removed
func (bc *Blockchain) FindUTXO() map[string]TXOutputs {
	UTXO := make(map[string]TXOutputs)
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOutIdx := range spentTXOs[txID] {
						if spentOutIdx == outIdx {
							continue Outputs
						}
					}
				}

				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return UTXO
}

// FindSpendableOutputs finds and returns unspent outputs to reference in inputs
func (bc *Blockchain) FindSpendableOutputs(pubKeyHash []byte, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(pubKeyHash)
	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.IsLockedWithKey(pubKeyHash) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}

// FindTransaction finds a transaction by its ID
func (bc *Blockchain) FindTransaction(ID []byte) (Transaction, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			if bytes.Compare(tx.ID, ID) == 0 {
				return *tx, nil
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return Transaction{}, errors.New("Transaction is not found")
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

// NewBlockchain get a last block in db
func NewBlockchain() *Blockchain {
	var bc Blockchain
	lastBlock := blockCollection.GetLastRecord()
	if len(lastBlock) == 0 {
		// cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
		// genesis := NewGenesisBlock(cbtx)
		// _, err := blockCollection.CreateByLambda(genesis)
		// if err != nil {
		// 	log.Fatal("Create failed", err)
		// }
		// bc = Blockchain{genesis.Hash, blockCollection}
		fmt.Println("No existing blockchain found. Create one first.")
	} else {
		var block Block

		err := bson.Unmarshal(lastBlock[0], &block)
		if err != nil {
			log.Fatal("detail:", err)
		}
		bc = Blockchain{block.Hash, blockCollection}
	}
	return &bc
}

// CreateBlockchain creates a new blockchain DB
func CreateBlockchain(address string) *Blockchain {
	var bc Blockchain
	lastBlock := blockCollection.GetLastRecord()
	if len(lastBlock) == 0 {
		cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)
		_, err := blockCollection.CreateByLambda(genesis)
		if err != nil {
			log.Fatal("Create failed", err)
		}
		bc = Blockchain{genesis.Hash, blockCollection}

	} else {
		var block Block

		err := bson.Unmarshal(lastBlock[0], &block)
		if err != nil {
			log.Fatal("detail:", err)
		}
		bc = Blockchain{block.Hash, blockCollection}
	}
	return &bc
}

// VerifyTransaction verifies transaction input signatures
func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	return tx.Verify(prevTXs)
}
