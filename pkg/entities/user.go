package entities

import (
	"RestFullAPI-todo/pkg/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at"`
	CreatedBy string             `bson:"created_by"`
	UpdatedBy string             `bson:"updated_by"`
	Status    enums.Status       `bson:"status"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
}
