package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/AyokunlePaul/Owo-Eyo/blockchain/block"
	"log"
	"time"
)

type Blockchain struct {
	chain   []*block.Block
	genesis block.Block
}

func (blockchain *Blockchain) createBlock(proof int, previousHash string, data string) *block.Block {
	chainLength := len(blockchain.chain)
	newBlock := &block.Block{
		Index:        chainLength + 1,
		PreviousHash: previousHash,
		Proof:        proof,
		Timestamp:    time.Now().String(),
		Data:         []byte(data),
	}
	newBlock.DeriveHash()
	blockchain.chain = append(blockchain.chain, newBlock)
	return newBlock
}

func (blockchain *Blockchain) getLastBlock() *block.Block {
	return blockchain.chain[len(blockchain.chain)-1]
}

func (blockchain *Blockchain) proofOfWork(previousProof int) int {
	newProof := 1
	checkProof := false
	for checkProof {
		proofHash := sha256.Sum256(
			bytes.Join([][]byte{
				ToHex(int64((newProof * newProof) - (previousProof - previousProof))),
			}, []byte{}))
		checkProof = fmt.Sprintf("%x", proofHash[:4]) == "0000"
		if !checkProof {
			newProof += 1
		}
	}
	return newProof
}

func (blockchain *Blockchain) isChainValid() bool {
	previousBlock := blockchain.chain[0]
	blockIndex := 1
	for blockIndex < len(blockchain.chain) {
		currentBlock := blockchain.chain[blockIndex]
		if isValidLink := previousBlock.Hash == currentBlock.PreviousHash; !isValidLink {
			return false
		}
		proofHash := sha256.Sum256(
			bytes.Join([][]byte{
				ToHex(int64((currentBlock.Proof * currentBlock.Proof) - (previousBlock.Proof - previousBlock.Proof))),
			}, []byte{}))
		if fmt.Sprintf("%x", proofHash[:4]) != "0000" {
			return false
		}

		blockIndex += 1
		previousBlock = currentBlock
	}
	return false
}

func ToHex(value int64) []byte {
	buff := new(bytes.Buffer)
	writeError := binary.Write(buff, binary.BigEndian, value)
	if writeError != nil {
		log.Panic(writeError)
	}
	return buff.Bytes()
}
