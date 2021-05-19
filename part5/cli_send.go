package blockchain

import (
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := NewBlockchain(from)
	// defer bc.db.Close()

	tx, _ := NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("Success!")
}

func Send(from, to string, amount int) string {
	if !ValidateAddress(from) {
		return "ERROR: Sender address is not valid"
	}
	if !ValidateAddress(to) {
		return "ERROR: Recipient address is not valid"
	}

	bc := NewBlockchain(from)
	// defer bc.db.Close()

	tx, rs := NewUTXOTransaction(from, to, amount, bc)
	if rs != "" {
		return rs
	}
	cbTx := NewCoinbaseTX(from, "")
	txs := []*Transaction{cbTx, tx}

	bc.MineBlock(txs)
	return "Success!"
}
