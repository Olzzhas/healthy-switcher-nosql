package dish

import "time"

type Dish struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	Title       string    `json:"title"bson:"title"`
	Description string    `json:"description"bson:"description"`
	Rating      int64     `json:"rating"bson:"rating"`
	Comments    []Comment `json:"comments"bson:"comments"`
}

type Comment struct {
	ID string `json:"id" bson:"_id,omitempty"`
	//User User
	CommentBody string `json:"comment_body" bson:"comment_body"`
	Rating      int64  `json:"rating" bson:"rating"`
}
