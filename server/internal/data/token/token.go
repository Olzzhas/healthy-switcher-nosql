package token

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"server/internal/validator"
	"time"
)

const (
	scopeActivation = "activation"
)

type Token struct {
	Plaintext string    `json:"plaintext"bson:"plaintext"`
	Hash      []byte    `json:"hash"bson:"hash"`
	UserID    string    `json:"user_id"bson:"user_id"`
	Expiry    time.Time `json:"expiry"bson:"expiry"`
	Scope     string    `json:"scope"bson:"scope"`
}

func generateToken(userID string, ttl time.Duration, scope string) (*Token, error) {

	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	return token, nil
}

func New(userID string, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = m.Insert(token)
	return token, err
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}
