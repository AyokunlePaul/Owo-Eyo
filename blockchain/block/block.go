package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Block struct {
	PreviousHash string
	Data         string
	Hash         string
	Index        int
	Timestamp    string
	Nonce        int
}

func (block *Block) DeriveHash() {
	blockInfo := bytes.Join([][]byte{
		[]byte(block.Data),
		[]byte(block.PreviousHash),
		[]byte(block.Timestamp),
		[]byte(strconv.Itoa(block.Index))},
		[]byte{},
	)
	hash := sha256.Sum256(blockInfo)
	block.Hash = fmt.Sprintf("%x", hash)
}
