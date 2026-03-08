// Transaction represents a single transaction in the blockchain.
type Transaction struct {
    ID     string    `json:"id"`     // Unique identifier for the transaction
    Amount float64   `json:"amount"` // Amount of cryptocurrency transferred
    From   string    `json:"from"`   // Sender's address
    To     string    `json:"to"`     // Receiver's address
    Timestamp string  `json:"timestamp"` // Time when the transaction occurred
}