package token

import (
	"server/internal/validator"
	"time"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
)

type Token struct {
	Plaintext string    `json:"token"bson:"token"`
	Hash      []byte    `json:"-"bson:"-"`
	UserID    string    `json:"-"bson:"user_id"`
	Expiry    time.Time `json:"expiry"bson:"expiry"`
	Scope     string    `json:"-"bson:"-"`
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}
