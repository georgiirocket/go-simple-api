package schemas

import (
	"go-simple-api/utils/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Username  string        `bson:"username"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"created_at"`
}

func (u *User) ToModel() models.UserModel {
	return models.UserModel{
		ID:        u.ID.Hex(),
		Username:  u.Username,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}
}
