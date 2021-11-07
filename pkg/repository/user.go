package repository

import (
	"cloud.google.com/go/firestore"
	"mocerize-api/model"
)

type UserRepository interface {
	FindByUID(UID string) (*model.User, error)
	FindByToken(Token string) (*model.User, error)
}

type firestoreDBRepository struct {
	DB *firestore.Client
}

// create new user repository via firestoreDB

func NewUserRepository(fireStoreClient *firestore.Client) UserRepository {
	return &firestoreDBRepository{DB: fireStoreClient}
}

func (r firestoreDBRepository) FindByUID(UID string) (*model.User, error) {
	panic("implement me")
}

func (r firestoreDBRepository) FindByToken(Token string) (*model.User, error) {
	panic("implement me")
}
