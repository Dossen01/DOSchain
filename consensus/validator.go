package consensus

// Validator represents a participant in the Proof of Stake consensus mechanism.
type Validator struct {
	ID             string  // Unique identifier for the validator
	Stake          float64 // Amount of stake the validator has
	CommissionRate  float64 // Commission rate taken from the rewards
	Active         bool    // Status of the validator (active or inactive)
}