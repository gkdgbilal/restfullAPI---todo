package todo

import (
	"RestFullAPI-todo/api/utils"
	"RestFullAPI-todo/configs/logg"
	"RestFullAPI-todo/pkg/dto"
	"RestFullAPI-todo/pkg/entities"
	"RestFullAPI-todo/pkg/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	Create(todo *dto.TodoDTO) (*dto.TodoResponse, error)
	Read(id string) (*dto.TodoResponse, error)
	Reads(*utils.Pageable) (*dto.ListResponse, error)
	Update(todo *dto.TodoDTO) (*entities.Todo, error)
	Completed(todo *dto.TodoDTO) (*entities.Todo, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func (s service) Completed(todoDto *dto.TodoDTO) (*entities.Todo, error) {
	todoID, err := primitive.ObjectIDFromHex(todoDto.ID)
	if err != nil {
		return nil, err
	}
	//if todoDto.Completed == false {
	//	todoDto.Completed = true
	//} else {
	//	todoDto.Completed = false
	//}
	todo := &entities.Todo{
		ID:          todoID,
		Title:       todoDto.Title,
		Description: todoDto.Description,
		Completed:   !todoDto.Completed,
		Status:      enums.Active,
	}
	return s.repository.Completed(todo)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) Create(todo *dto.TodoDTO) (*dto.TodoResponse, error) {
	t := new(entities.Todo)
	t.Title = todo.Title
	t.Description = todo.Description
	t.Completed = todo.Completed
	create, err := s.repository.Create(t)
	if err != nil {
		logg.L.Error("repository.Create: ", zap.Error(err))
		return nil, err
	}

	return &dto.TodoResponse{
		ID:          create.ID.Hex(),
		Title:       create.Title,
		Description: create.Description,
		Completed:   create.Completed,
	}, nil
}

func (s service) Read(id string) (*dto.TodoResponse, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	todo, err := s.repository.Read(ID)
	if err != nil {
		return nil, err
	}
	return &dto.TodoResponse{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}, nil
}

func (s service) Reads(pageable *utils.Pageable) (*dto.ListResponse, error) {
	todos, err := s.repository.Reads(pageable)
	if err != nil {
		return nil, err
	}
	var todoResponse []dto.TodoResponse
	for _, todo := range todos {
		todoResponse = append(todoResponse,
			dto.TodoResponse{
				ID:          todo.ID.Hex(),
				Title:       todo.Title,
				Description: todo.Description,
				Completed:   todo.Completed,
				CreatedAt:   todo.CreatedAt,
				UpdatedAt:   time.Now(),
			})
	}
	res := dto.ListResponse{
		Page:          pageable.Page,
		Size:          pageable.Size,
		TotalPages:    pageable.GetTotalPage(),
		TotalElements: pageable.TotalElements,
		IsLastPage:    pageable.IsLast(),
		HasNext:       pageable.HasNext(),
		Content:       todoResponse,
	}
	return &res, nil
}

func (s service) Update(todoDto *dto.TodoDTO) (*entities.Todo, error) {
	todoID, err := primitive.ObjectIDFromHex(todoDto.ID)
	if err != nil {
		return nil, err
	}
	todo := &entities.Todo{
		ID:          todoID,
		Title:       todoDto.Title,
		Description: todoDto.Description,
		Completed:   todoDto.Completed,
		Status:      enums.Active,
	}
	return s.repository.Update(todo)
}

func (s service) Delete(id string) error {
	//return s.repository.Delete(id, hardDelete)
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(ID)
}
