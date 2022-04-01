package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required, min=2, max=100"`
	Last_name     *string            `json:"last_name" validate:"reduired, min=2, max=100"`
	Password      *string            `json:"password" validate:"required, min=8"`
	Email         *string            `json:"email" validate:"email, reduired"`
	Phone         *string            `json:"phone" validate:"reduired"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"reduired, eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updtaed_at    time.Time          `json:"updated_at"`
	User_id       *string            `json:"user_id"`
}
