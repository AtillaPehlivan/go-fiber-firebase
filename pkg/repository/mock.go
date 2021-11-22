package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"mocerize-api/model"
	"mocerize-api/pkg/helper"
)

type MockRepository interface {
	Index(userUID string) (*[]model.Mock, error)
	FindByUID(string) (*model.Mock, error)
	Update(string, *model.Mock) (*model.Mock, error)
	Create(userUID string, mock *model.Mock) (*model.Mock, error)
	Destroy(mockUID string, userUID string) (bool, error)
}

type firestoreDBMockRepository struct {
	DB *firestore.Client
}

// create new mock repository via firestoreDB

func NewMockRepository(fireStoreClient *firestore.Client) MockRepository {
	return &firestoreDBMockRepository{DB: fireStoreClient}
}

func (r firestoreDBMockRepository) Index(userUID string) (*[]model.Mock, error) {

	var mocks []model.Mock

	mockIter := r.DB.Collection("mocks").Where("user_uid", "==", userUID).Documents(context.Background())

	for {
		doc, err := mockIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var mock model.Mock

		if err := doc.DataTo(&mock); err != nil {
			return nil, err
		}

		mocks = append(mocks, mock)

	}

	return &mocks, nil
}

func (r firestoreDBMockRepository) FindByUID(UID string) (*model.Mock, error) {

	doc, err := r.DB.Collection("mocks").Doc(UID).Get(context.Background())

	if err != nil {
		return nil, err
	}

	var mockData model.Mock
	err = doc.DataTo(&mockData)

	if err != nil {
		return nil, err
	}

	return &mockData, nil
}

func (r firestoreDBMockRepository) Update(UID string, mock *model.Mock) (*model.Mock, error) {

	updateData, _ := helper.StructToMapString(mock)

	_, err := r.DB.Collection("mocks").Doc(UID).Set(context.Background(), updateData, firestore.MergeAll)

	if err != nil {
		return nil, err
	}

	return mock, nil

}

func (r firestoreDBMockRepository) Create(userUID string, mock *model.Mock) (*model.Mock, error) {

	newUid := uuid.NewString()

	mock.Uid = newUid
	mock.UserUid = userUID
	mock.Status = "default_status"

	createdData, _ := helper.StructToMapString(mock)

	_, err := r.DB.Collection("mocks").Doc(newUid).Set(context.Background(), createdData)

	if err != nil {
		return nil, err
	}

	return mock, nil
}

func (r firestoreDBMockRepository) Destroy(mockUID string, userUID string) (bool, error) {

	mockIter := r.DB.Collection("mocks").Where("uid", "==", mockUID).Where("user_uid", "==", userUID).Documents(context.Background())

	for {
		doc, err := mockIter.Next()

		if err == iterator.Done {
			break
		}

		_, err = doc.Ref.Delete(context.Background())
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
