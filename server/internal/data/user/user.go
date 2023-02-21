package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"server/internal/data/dish"
	"server/internal/validator"
	"time"
)

type User struct {
	ID        string    `json:"id"bson:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  password  `json:"-" bson:"-"`
	Activated bool      `json:"activated" bson:"activated"`
	Orders    []Order   `json:"orders"bson:"orders"`
	Version   int       `json:"-" bson:"version"`
}

//type CreatedUserDTO struct {
//	Name     string   `json:"name"`
//	Email    string   `json:"email"`
//	Password password `json:"password "`
//}

type password struct {
	plaintext *string
	hash      []byte
}

type Order struct {
	ID        string    `json:"id"bson:"id"`
	Dish      dish.Dish `json:"dish"bson:"dish"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))

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

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}
