package entities

import (
	"RestFullAPI-todo/pkg/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   *time.Time         `bson:"updated_at"`
	DeletedAt   *time.Time         `bson:"deleted_at"`
	Status      enums.Status       `bson:"status"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
}
