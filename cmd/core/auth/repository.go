package auth

import (
	"context"
	"go-simple-api/libs/schemas"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type Repository struct {
	db *mongo.Collection
}

func NewAuthRepository(db *mongo.Database, collection string) *Repository {
	return &Repository{
		db: db.Collection(collection),
	}
}

func newUser(username, password string) schemas.User {
	return schemas.User{
		ID:        bson.NewObjectID(),
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}
}

func (r Repository) CreateUser(ctx context.Context, username, password string) (*schemas.User, error) {
	model := newUser(username, password)
	_, err := r.db.InsertOne(ctx, model)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r Repository) GetUser(ctx context.Context, username string) (*schemas.User, error) {
	user := new(schemas.User)
	err := r.db.FindOne(ctx, bson.M{
		"username": username,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
