package dish

import (
	"server/internal/validator"
	"time"
)

type Dish struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	Title       string    `json:"title"bson:"title"`
	Img         string    `json:"img"bson:"img"`
	Price       int64     `json:"price"bson:"price"`
	Description string    `json:"description"bson:"description"`
	Rating      int64     `json:"rating"bson:"rating"`
	Comments    []Comment `json:"comments"bson:"comments"`
}

type Comment struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	UserID      string `json:"user_id",bson:"user_id"`
	CommentBody string `json:"comment_body" bson:"comment_body"`
	Rating      int64  `json:"rating" bson:"rating"`
}

func ValidateDish(v *validator.Validator, dish *Dish) {
	v.Check(dish.Title != "", "title", "title cannot be empty")
	v.Check(len(dish.Title) <= 128, "title", "must not be more than 128 bytes long")

	v.Check(dish.Description != "", "description", "description cannot be empty")
	v.Check(len(dish.Description) <= 4096, "description", "must not be more than 128 bytes long")
}
