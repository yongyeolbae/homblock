package main

import (
	"github.com/baeyongyeol/blockchain"
	"github.com/baeyongyeol/cli"
	"github.com/baeyongyeol/db"
)

func main() {
	/*blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
	blockchain.Blockchain().AddBlock("Fourth")*/
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()

}
