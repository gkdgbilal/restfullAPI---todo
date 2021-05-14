package auth

import (
	"RestFullAPI-todo/api/utils"
	"RestFullAPI-todo/api/utils/jwt"
	"RestFullAPI-todo/api/utils/password"
	"RestFullAPI-todo/configs/logg"
	"RestFullAPI-todo/pkg/dto"
	"RestFullAPI-todo/pkg/entities"
	"RestFullAPI-todo/pkg/user"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service interface {
	SignUp(register *dto.RegisterDTO) (*dto.AuthResponse, error)
	Login(login *dto.LoginDTO) (*dto.AuthResponse, error)
	IsUserActiveByUsername(username string) bool
}

type service struct {
	repository user.Repository
}

func NewService(repository user.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) SignUp(register *dto.RegisterDTO) (*dto.AuthResponse, error) {
	u, err := s.repository.FindByEmail(register.Email)

	//If email already exists, return
	if u != nil {
		return nil, utils.ErrEmailAlreadyTaken
	}

	u, err = s.repository.FindByUsername(register.Username)

	//If email username exists, return
	if u != nil {
		return nil, utils.ErrUsernameAlreadyTaken
	}

	u = &entities.User{
		Username: register.Username,
		Password: password.Generate(register.Password),
		Email:    register.Email,
	}

	// Create a user, if error return
	u, err = s.repository.Create(u)

	if err != nil {
		return nil, utils.NewError(fiber.StatusBadRequest, "Something went wrong!")
	}

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID:       u.ID.Hex(),
		Username: u.Username,
	})

	return &dto.AuthResponse{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Email:    u.Email,
		Token:    t,
	}, nil
}

func (s service) Login(login *dto.LoginDTO) (*dto.AuthResponse, error) {
	u, err := s.repository.FindByUsername(login.Username)

	// If username not exists, return
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, utils.ErrInvalidPasswordUsername
	}

	// If password not match, return
	if err := password.Verify(u.Password, login.Password); err != nil {
		return nil, utils.ErrInvalidPasswordUsername
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID:       u.ID.Hex(),
		Username: u.Username,
	})

	return &dto.AuthResponse{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Email:    u.Email,
		Token:    t,
	}, nil
}

func (s *service) IsUserActiveByUsername(username string) bool {
	u, err := s.repository.FindByUsername(username)
	if err != nil {
		logg.L.Error("IsUserActiveByUsername", zap.Error(err))
		return false
	}
	return u != nil
}
