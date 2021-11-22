package service

import (
	"mocerize-api/model"
	"mocerize-api/pkg/repository"
)

type UserService interface {
	FindByUID(UID string) (*model.User, error)
	FindByToken(Token string) (*model.User, error)
	Update(UID string, User *model.User) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

// create new user service via UserRepository

func NewUserService(r repository.UserRepository) UserService {
	return &userService{userRepository: r}
}

func (s userService) FindByUID(UID string) (*model.User, error) {
	return s.userRepository.FindByUID(UID)
}

func (s userService) FindByToken(Token string) (*model.User, error) {
	return s.userRepository.FindByToken(Token)
}

func (s userService) Update(UID string, User *model.User) (*model.User, error) {
	return s.userRepository.Update(UID, User)
}
