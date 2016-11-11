package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)
//This is a user model for store a user information.
type User struct {
	Id bson.ObjectId    `bson:"_id,omitempty" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	HashPassword []byte `json:"hashpwd"`
	CreatedAt time.Time `json:"createAt"`
}

type RegisterUser struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Password2  string    `json:"password2"`
}