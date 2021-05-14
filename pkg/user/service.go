package user

import (
	"RestFullAPI-todo/pkg/dto"
	"RestFullAPI-todo/pkg/entities"
)

type Service interface {
	Create(user *entities.User) (*entities.User, error)
	Reads() ([]dto.UserResponse, error)
	Read(id string) (*dto.UserResponse, error)
	ReadByUsername(username string) (*dto.UserResponse, error)
	Update(user *dto.UserDTO) (*entities.User, error)
	Delete(id string, hardDelete bool) error
}

type service struct {
	repository Repository
}

func (s service) Create(user *entities.User) (*entities.User, error) {
	panic("implement me")
}

func (s service) Reads() ([]dto.UserResponse, error) {
	panic("implement me")
}

func (s service) Read(id string) (*dto.UserResponse, error) {
	panic("implement me")
}

func (s service) ReadByUsername(username string) (*dto.UserResponse, error) {
	panic("implement me")
}

func (s service) Update(user *dto.UserDTO) (*entities.User, error) {
	panic("implement me")
}

func (s service) Delete(id string, hardDelete bool) error {
	panic("implement me")
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
