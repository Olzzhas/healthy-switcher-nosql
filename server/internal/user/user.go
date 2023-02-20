package user

import "time"

type User struct {
	ID        string    `json:"id"bson:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"-" bson:"password"`
	Activated bool      `json:"activated" bson:"activated"`
	Version   int       `json:"-" bson:"version"`
}

type CreatedUserDTO struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password password `json:"password "`
}

type password struct {
	plaintext *string
	hash      []byte
}
