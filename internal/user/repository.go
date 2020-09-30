package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository  belongs to user
type Repository interface {
	Get(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user User) error
}

type repository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

// NewRepository creates a new album repository
func NewRepository(db *mongo.Database) Repository {
	return repository{db, db.Collection(collectionName)}
}

func (r repository) Get(ctx context.Context, id string) (User, error) {
	var user User
	err := r.collection.FindOne(ctx, bson.M{"userID": id}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r repository) Create(ctx context.Context, user User) error {
	user.ID = primitive.NewObjectID()
	fmt.Println(user)
	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

const collectionName = "users"
