package post

import (
	"context"
	"go-simple-api/utils/schemas"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type Repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Database, collection string) *Repository {
	return &Repository{
		db: db.Collection(collection),
	}
}

func newPost(userId, title, description string) (schemas.Post, error) {
	objectId, err := bson.ObjectIDFromHex(userId)

	if err != nil {
		return schemas.Post{}, err
	}

	newPost := schemas.Post{
		ID:          bson.NewObjectID(),
		Title:       title,
		Description: description,
		UserId:      objectId,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now()}

	return newPost, nil
}

func (r Repository) CreatePost(ctx context.Context, userId, title, description string) (*schemas.Post, error) {
	entity, err := newPost(userId, title, description)

	if err != nil {
		return nil, err
	}

	_, err = r.db.InsertOne(ctx, entity)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r Repository) GetPosts(ctx context.Context) ([]*schemas.Post, error) {
	var entities = make([]*schemas.Post, 0)

	cursor, err := r.db.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}

func (r Repository) GetPostById(ctx context.Context, postId string) (*schemas.Post, error) {
	objectId, err := bson.ObjectIDFromHex(postId)

	if err != nil {
		return nil, err
	}

	entity := new(schemas.Post)

	err = r.db.FindOne(ctx, bson.M{"_id": objectId}).Decode(entity)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r Repository) UpdatePost(ctx context.Context, userId, postId, title, description string) (*schemas.Post, error) {
	userObjectId, err := bson.ObjectIDFromHex(userId)
	objectId, err := bson.ObjectIDFromHex(postId)

	if err != nil {
		return nil, err
	}

	entity := new(schemas.Post)
	filter := bson.M{"_id": objectId, "user_id": userObjectId}
	update := bson.M{"$set": bson.M{"title": title, "description": description, "updated_at": time.Now()}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = r.db.FindOneAndUpdate(ctx, filter, update, opts).Decode(entity)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r Repository) DeletePost(ctx context.Context, userId, postId string) (*schemas.Post, error) {
	userObjectId, err := bson.ObjectIDFromHex(userId)
	objectId, err := bson.ObjectIDFromHex(postId)

	if err != nil {
		return nil, err
	}

	entity := new(schemas.Post)

	err = r.db.FindOneAndDelete(ctx, bson.M{"_id": objectId, "user_id": userObjectId}).Decode(entity)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
