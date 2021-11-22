package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"mocerize-api/model"
	"mocerize-api/pkg/helper"
)

type UserRepository interface {
	FindByUID(UID string) (*model.User, error)
	FindByToken(Token string) (*model.User, error)
	Update(UID string, User *model.User) (*model.User, error)
}

type firestoreDBUserRepository struct {
	DB *firestore.Client
}

// create new user repository via firestoreDB

func NewUserRepository(fireStoreClient *firestore.Client) UserRepository {
	return &firestoreDBUserRepository{DB: fireStoreClient}
}

func (r firestoreDBUserRepository) FindByUID(UID string) (*model.User, error) {

	doc, err := r.DB.Collection("users").Doc(UID).Get(context.Background())

	if err != nil {
		return nil, err
	}

	var userData model.User
	err = doc.DataTo(&userData)

	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func (r firestoreDBUserRepository) FindByToken(Token string) (*model.User, error) {
	panic("implement me")
}

func (r firestoreDBUserRepository) Update(UID string, user *model.User) (*model.User, error) {

	updateData, _ := helper.StructToMapString(user)

	_, err := r.DB.Collection("users").Doc(UID).Set(context.Background(), updateData, firestore.MergeAll)

	if err != nil {
		return nil, err
	}

	return user, nil

}
