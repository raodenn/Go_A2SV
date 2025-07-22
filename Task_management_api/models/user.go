package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username     *string            `bson:"username" json:"username" validate:"required,min=2,max=100"`
	Password     *string            `bson:"password" json:"password,omitempty" validate:"required,min=6"`
	UserType     *string            `bson:"usertype" json:"usertype" validate:"required,eq=ADMIN|eq=USER"`
	Token        *string            `bson:"token,omitempty" json:"token,omitempty"`
	RefreshToken *string            `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	UserId       string             `bson:"user_id" json:"user_id"`
}
