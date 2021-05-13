package blockchain

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain() {
	bc := NewBlockchain()
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Println("============ Block ", block.Hash)

		// fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		fmt.Println("Prev. block: ", block.PrevBlockHash)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// type response struct {
// 	Block struct {
// 		Hash         string
// 		PrevHash     string
// 		Transactions []Transaction
// 	}
// }

type TransactionRes struct {
	TXID      interface{}
	From      interface{}
	To        interface{}
	Timestamp int64
	Value     int
}

func PrintChain() interface{} {
	var blocksResponse []Block
	var transRes []TransactionRes
	bc := NewBlockchain()
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)

			for _, txout := range tx.Vout {
				fmt.Println(txout)

				for _, txin := range tx.Vin {
					fmt.Println(txin)
					var item TransactionRes
					item.TXID = tx.ID
					item.From = txin.Signature
					item.To = txout.PubKeyHash
					item.Timestamp = block.Timestamp
					item.Value = txout.Value
					transRes = append(transRes, item)
				}

			}
		}
		fmt.Printf("\n\n")

		blocksResponse = append(blocksResponse, *block)
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return transRes

}
