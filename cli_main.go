package main

import (
	blockchain "mycoin/blockchain"
)

func main() {

	// Bc.AddBlock("Send 1 BTC to Ivan")
	// bc.AddBlock("Send 2 more BTC to Ivan")

	cli := blockchain.CLI{}
	cli.Run()
}
