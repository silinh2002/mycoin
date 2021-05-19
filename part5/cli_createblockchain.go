package blockchain

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	CreateBlockchain(address)
	// bc.db.Close()
	fmt.Println("Done!")
}

func InitBlockchain(address string) string {
	if !ValidateAddress(address) {
		return "ERROR: Address is not valid"
	}
	CreateBlockchain(address)

	return "Done!"
}
