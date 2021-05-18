package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model for the User collection in the DB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	LastName  string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Birthdate time.Time          `bson:"birthdate" json:"birthdate"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Banner    string             `bson:"banner" json:"banner"`
	Bio       string             `bson:"bio" json:"bio"`
	Location  string             `bson:"location" json:"location"`
	Website   string             `bson:"website" json:"website"`
}
