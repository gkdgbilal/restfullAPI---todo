package todo

import (
	"RestFullAPI-todo/api/utils"
	"RestFullAPI-todo/configs/logg"
	"RestFullAPI-todo/pkg/entities"
	"RestFullAPI-todo/pkg/enums"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

type Repository interface {
	Create(todo *entities.Todo) (*entities.Todo, error)
	Reads(*utils.Pageable) ([]entities.Todo, error)
	Read(ID primitive.ObjectID) (*entities.Todo, error)
	Update(todo *entities.Todo) (*entities.Todo, error)
	Completed(todo *entities.Todo) (*entities.Todo, error)
	Delete(ID primitive.ObjectID) error
}

type repository struct {
	Collection *mongo.Collection
}

func (r repository) Completed(todo *entities.Todo) (*entities.Todo, error) {
	todo.UpdatedAt = utils.PtrDate()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": todo.ID}, bson.M{"$set": todo})
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r repository) Create(todo *entities.Todo) (*entities.Todo, error) {
	todo.ID = primitive.NewObjectID()
	todo.CreatedAt = time.Now()
	todo.Status = enums.Active
	_, err := r.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		logg.L.Error("todo create", zap.Error(err))
		return nil, err
	}
	return todo, err
}

func (r repository) Reads(pageable *utils.Pageable) ([]entities.Todo, error) {
	var todos []entities.Todo

	var opts options.FindOptions
	skip := pageable.Page * pageable.Size
	opts.Limit = &pageable.Size
	opts.Skip = &skip
	opts.Sort = bson.D{{pageable.GetSortKey(), pageable.GetSortValue()}}
	filter := bson.D{}

	count, err := r.Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	pageable.TotalElements = count
	cursor, err := r.Collection.Find(context.Background(), filter, &opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var todo entities.Todo
		_ = cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r repository) Read(ID primitive.ObjectID) (*entities.Todo, error) {
	var todo entities.Todo
	err := r.Collection.FindOne(context.Background(), bson.D{{"_id", ID}, {"status", enums.Active}}).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r repository) Update(todo *entities.Todo) (*entities.Todo, error) {
	todo.UpdatedAt = utils.PtrDate()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": todo.ID}, bson.M{"$set": todo})
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r repository) Delete(ID primitive.ObjectID) error {
	//todoID, err := primitive.ObjectIDFromHex(ID)
	//if err != nil {
	//	return err
	//}
	//if hardDelete {
	//	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": todoID})
	//	return err
	//}
	//todo, err := r.Read(todoID)
	//if err != nil {
	//	return err
	//}
	//todo.DeletedAt = utils.PtrDate()
	//todo.Status = enums.Deleted
	//_, err = r.Collection.UpdateOne(context.Background(), bson.M{"_id": todoID}, bson.M{"$set": todo})
	//if err != nil {
	//	return err
	//}
	//return nil
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": ID})
	if err != nil {
		return err
	}
	return nil
}
