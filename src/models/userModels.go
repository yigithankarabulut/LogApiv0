package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UID      primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Username string             `json:"username"bson:"username"`
	Password string             `json:"password"bson:"password"`
}
