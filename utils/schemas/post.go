package schemas

import (
	"go-simple-api/utils/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Post struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	UserId      bson.ObjectID `bson:"user_id"`
	UpdatedAt   time.Time     `bson:"updated_at"`
	CreatedAt   time.Time     `bson:"created_at"`
}

func (u *Post) ToModel() models.PostModel {
	return models.PostModel{
		ID:          u.ID.Hex(),
		Title:       u.Title,
		Description: u.Description,
		UserId:      u.UserId.Hex(),
		UpdatedAt:   u.UpdatedAt.Format(time.RFC3339),
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
	}
}

func PostsToModels(arr []*Post) []models.PostModel {
	var m []models.PostModel
	for _, entity := range arr {
		m = append(m, entity.ToModel())
	}

	return m
}
