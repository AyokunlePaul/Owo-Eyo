package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

type Blockchain struct {
	Chain   []*Block
	genesis Block
}

func (blockchain *Blockchain) CreateBlock(previousHash string, data string) *Block {
	chainLength := len(blockchain.Chain)
	newBlock := &Block{
		Index:        chainLength + 1,
		PreviousHash: previousHash,
		Timestamp:    time.Now().String(),
		Data:         data,
	}
	proofOfWork := NewProof(newBlock)
	nonce := proofOfWork.Run()
	newBlock.Nonce = nonce
	newBlock.DeriveHash()
	blockchain.Chain = append(blockchain.Chain, newBlock)
	return newBlock
}

func (blockchain *Blockchain) GetPreviousBlock() *Block {
	return blockchain.Chain[len(blockchain.Chain)-1]
}

func (blockchain *Blockchain) ProofOfWork(previousProof int, data string) int {
	newProof := 1
	checkProof := true
	proofStartTime := time.Now().String()
	for checkProof {
		proofHash := sha256.Sum256(
			bytes.Join([][]byte{
				ToHex(int64(previousProof)),
				ToHex(int64(newProof)),
				[]byte(data),
				[]byte(proofStartTime),
			}, []byte{}))
		checkProof = fmt.Sprintf("%x", proofHash[:3]) != "000000"
		if checkProof {
			newProof += 1
		}
	}
	return newProof
}

func (blockchain *Blockchain) IsChainValid() bool {
	previousBlock := blockchain.Chain[0]
	blockIndex := 1
	for blockIndex < len(blockchain.Chain) {
		currentBlock := blockchain.Chain[blockIndex]
		if isValidLink := previousBlock.Hash == currentBlock.PreviousHash; !isValidLink {
			return false
		}
		proofHash := sha256.Sum256(
			bytes.Join([][]byte{
				ToHex(int64((currentBlock.Nonce * currentBlock.Nonce) - (previousBlock.Nonce * previousBlock.Nonce))),
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
