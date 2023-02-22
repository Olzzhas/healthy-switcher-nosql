package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"server/internal/data/dish"
	"server/internal/validator"
	"time"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

var AnonymousUser = &User{}

type User struct {
	ID        string    `json:"id"bson:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  password  `json:"-" bson:"password"`
	Hash      []byte    `json:"hash"bson:"hash"`
	Activated bool      `json:"activated" bson:"activated"`
	Orders    []Order   `json:"orders"bson:"orders"`
	Version   int       `json:"-" bson:"version"`
}

type password struct {
	plaintext *string
	hash      []byte
}

type Order struct {
	ID        string    `json:"id"bson:"id"`
	Dish      dish.Dish `json:"dish"bson:"dish"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func (p *password) Set(plaintextPassword string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		log.Fatal(err)
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return hash, err
}

func (p *password) Matches(plaintextPassword string, hash []byte) (bool, error) {
	fmt.Println(hash)
	err := bcrypt.CompareHashAndPassword(hash, []byte(plaintextPassword))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}
func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}
func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	if user.Hash == nil {
		panic("missing password hash for user")
	}
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}
