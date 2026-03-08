package core

// Account structure for state management

type Account struct {
    ID      string  `json:"id"`
    Balance float64 `json:"balance"`
    Name    string  `json:"name"`
}