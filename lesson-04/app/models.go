package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	UID   int                `bson:"uid" json:"uid"`
	Lname string             `bson:"lname" json:"lname" form:"lname"`
	Fname string             `bson:"fname" json:"fname" form:"fname"`
	Email string             `bson:"email" json:"email" form:"email"`
}

// Users array
type Users = []User
