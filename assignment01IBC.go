package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PreviousBlock *Block
	HashValue     [32]byte
	Transaction   string
	Coins         int
}

func InsertBlock(Transaction string, chainHead *Block, Coins int) *Block {

	var newBlock *Block = new(Block)

	if chainHead == nil {

		newBlock.Transaction = Transaction
		newBlock.PreviousBlock = chainHead
		newBlock.Coins = Coins

	} else {

		newBlock.Transaction = Transaction
		newBlock.PreviousBlock = chainHead
		newBlock.HashValue = sha256.Sum256([]byte(chainHead.Transaction))
		newBlock.Coins = Coins
		println("New Block Added")
	}

	return newBlock
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	for chainHead != nil {
		if chainHead.Transaction == oldTrans {
			chainHead.Transaction = newTrans
			break
		}
		chainHead = chainHead.PreviousBlock
	}
}
func ListBlocks(chainHead *Block) {
	for chainHead.PreviousBlock != nil {
		println("Transaction " + chainHead.Transaction)
		fmt.Println(chainHead.HashValue)
		chainHead = chainHead.PreviousBlock

	}
}
func VerifyChain(chainHead *Block) {
	for chainHead.PreviousBlock != nil {
		if chainHead.HashValue != sha256.Sum256([]byte(chainHead.PreviousBlock.Transaction)) {
			println("HASH HAS CHANGED .... NOT A BLOCK CHAIN BUT OH WELL")
		}
		if chainHead.PreviousBlock == nil {
			println("Hash Not changed")
		}
		chainHead = chainHead.PreviousBlock
	}
}
