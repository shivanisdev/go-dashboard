package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User comment
type User struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID string             `json:"userID" bson:"userID"`
	Name   string             `json:"name" bson:"name"`
	Email  string             `json:"email" bson:"email"`
}
