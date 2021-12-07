package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Lname string             `bson:"lname" json:"lname"`
	Fname string             `bson:"fname" json:"fname"`
	Email string             `bson:"email" json:"email"`
}

// Users array
type Users = []User
