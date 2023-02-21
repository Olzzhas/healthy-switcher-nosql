package token

import (
	"server/internal/validator"
	"time"
)

const (
	ScopeActivation = "activation"
)

type Token struct {
	Plaintext string    `json:"plaintext"bson:"plaintext"`
	Hash      []byte    `json:"hash"bson:"hash"`
	UserID    string    `json:"user_id"bson:"user_id"`
	Expiry    time.Time `json:"expiry"bson:"expiry"`
	Scope     string    `json:"scope"bson:"scope"`
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}
