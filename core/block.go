package core

// Block defines a block in the blockchain

type Block struct {
    Index        int       `json:"index"`
    Timestamp    string    `json:"timestamp"`
    Data         string    `json:"data"`
    PrevBlockHash string    `json:"prev_block_hash"`
    Hash         string    `json:"hash"`
}