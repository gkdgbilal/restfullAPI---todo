package user

import (
	"RestFullAPI-todo/configs/logg"
	"RestFullAPI-todo/pkg/entities"
	"RestFullAPI-todo/pkg/enums"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

type Repository interface {
	Create(user *entities.User) (*entities.User, error)
	Reads() ([]entities.User, error)
	Read(ID primitive.ObjectID) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Delete(ID string, hardDelete bool) error
	FindByEmail(email string) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}
func (r *repository) Create(user *entities.User) (*entities.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.CreatedBy = user.Username
	user.Status = enums.Active
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		logg.L.Error("user create", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (r *repository) Reads() ([]entities.User, error) {
	var users []entities.User
	cursor, err := r.Collection.Find(context.Background(), bson.D{{"status", enums.Active}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user entities.User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

func (r *repository) Read(ID primitive.ObjectID) (*entities.User, error) {
	var user entities.User
	err := r.Collection.FindOne(context.Background(), bson.D{{"_id", ID}, {"status", enums.Active}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.Collection.FindOne(context.Background(), bson.D{{"email", email}, {"status", enums.Active}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.Collection.FindOne(context.Background(), bson.D{{"username", username}, {"status", enums.Active}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Update(user *entities.User) (*entities.User, error) {
	user.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) Delete(ID string, hardDelete bool) error {
	userID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	if hardDelete {
		_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": userID})
		return err
	}

	user, err := r.Read(userID)
	if err != nil {
		return err
	}
	user.DeletedAt = time.Now()
	user.Status = enums.Deleted
	_, err = r.Collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}
