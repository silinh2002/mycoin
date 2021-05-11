package main

import (
	part2 "mycoin/part3"
)

func main() {
	// bc := part2.NewBlockchain()

	// for _, block := range bc.Blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	pow := part2.NewProofOfWork(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }
	bc := part2.NewBlockchain()

	// Bc.AddBlock("Send 1 BTC to Ivan")
	// bc.AddBlock("Send 2 more BTC to Ivan")

	cli := part2.CLI{bc}
	cli.Run()
}
