package main

import (
	"fmt"
	"log"

	"github.com/techmexdev/blockchain"
	"github.com/techmexdev/blockchain/minerpool"
)

var mp MinerPool

func main() {
	bc := blockchain.NewBlockchain()
	var miners []Miner
	for i = 0; i < 5; i++ {
		miners = append(minerpool.NewMiner("Miner "+i, *bc))
	}
	mp = minerpool.NewMinerPool(miners)
	mp.ReceiveTransaction("Satoshi", "Ashley", 20)
	mp.ReceiveTransaction("Ashley", "Bob", 10)
	mp.ReceiveTransaction("Ashley", "Charly", 10)
	mp.ReceiveTransaction("Daniel", "Ashley", 500)

	satBal := mp.GetBalance("Satoshi")
	log.Printf("Satoshi's balance: %v", satBal)

	ashBal := mp.GetBalance("Ashley")
	log.Printf("Ashley's balance: %v", ashBal)

	bobBal := mp.GetBalance("Bob")
	log.Printf("Bob's balance: %v", bobBal)

	danBal := mp.GetBalance("Daniel")
	log.Printf("Daniel's balance: %v", danBal)

	bc = mp.Blockchain()
	fmt.Printf("Miner Pool Blockchain: %v", bc)
}
