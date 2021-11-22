package service

import (
	"mocerize-api/model"
	"mocerize-api/pkg/repository"
)

type MockService interface {
	Index(userUID string) (*[]model.Mock, error)
	FindByUID(string) (*model.Mock, error)
	Update(string, *model.Mock) (*model.Mock, error)
	Create(userUID string, mock *model.Mock) (*model.Mock, error)
	Destroy(mockUID string, userUID string) (bool, error)
}

type mockService struct {
	mockRepository repository.MockRepository
}

// create new user service via UserRepository

func NewMockService(r repository.MockRepository) MockService {
	return &mockService{mockRepository: r}
}

func (s mockService) Index(userUID string) (*[]model.Mock, error) {
	return s.mockRepository.Index(userUID)
}

func (s mockService) FindByUID(UID string) (*model.Mock, error) {
	return s.mockRepository.FindByUID(UID)
}

func (s mockService) Create(userUID string, mock *model.Mock) (*model.Mock, error) {
	return s.mockRepository.Create(userUID, mock)
}

func (s mockService) Update(UID string, Mock *model.Mock) (*model.Mock, error) {
	return s.mockRepository.Update(UID, Mock)
}

func (s mockService) Destroy(mockUID string, userUID string) (bool, error) {
	return s.mockRepository.Destroy(mockUID, userUID)
}
