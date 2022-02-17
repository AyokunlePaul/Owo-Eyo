package block

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

const Difficulty = 16

type OfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(block *Block) *OfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return &OfWork{
		Block:  block,
		Target: target,
	}
}

func (proofOfWork *OfWork) InitData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			[]byte(proofOfWork.Block.PreviousHash),
			[]byte(proofOfWork.Block.Data),
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		}, []byte{})
}

func (proofOfWork *OfWork) Run() int {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := proofOfWork.InitData(nonce)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(proofOfWork.Target) == -1 {
			break
		} else {
			nonce += 1
		}
	}
	return nonce
}

func (proofOfWork *OfWork) Validate() bool {
	var intHash big.Int
	
	data := proofOfWork.InitData(proofOfWork.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(proofOfWork.Target) == -1
}
