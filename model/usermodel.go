package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_Name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name     *string            `json:"last_name"validate:"required,min=2,max=100"`
	Password      *string            `json:"password"validate:"required,min=6"`
	Email         *string            `json:"email"validate: "required,email"`
	Phone         *string            `json:"phone"validate: "required"`
	Token         *string            `json:"token"`
	User_Type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_Token *string            `json:"refresh_token"`
	Create_At     time.Time          `json:"create_at"`
	Update_At     time.Time          `json:"update_at"`
	User_Id       string             `json:"user_id"`
}
