package core

type Block struct {
    Index     int
    Timestamp string
    Data      string
    PrevHash  string
}

type Blockchain struct {
    chain []Block
}

func NewBlockchain() *Blockchain {
    genesisBlock := Block{0, "2026-03-08 12:28:02", "", ""}
    return &Blockchain{chain: []Block{genesisBlock}}
}

func (bc *Blockchain) GetLatestBlock() Block {
    return bc.chain[len(bc.chain)-1]
}