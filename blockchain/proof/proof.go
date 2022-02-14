package proof

import (
	"github.com/AyokunlePaul/Owo-Eyo/blockchain/block"
	"math/big"
)

const Difficulty = 10

type ProofOfWork struct {
	Block  *block.Block
	Target *big.Int
}

func NewProof(block *block.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return nil
}
