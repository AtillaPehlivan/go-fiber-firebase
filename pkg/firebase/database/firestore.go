package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"mocerize-api/pkg/firebase"
)

var (
	firestoreClient *firestore.Client
)

//type Firestore struct {
//	*firestore.Client
//}

func Setup(ctx context.Context) error {

	client, err := firebase.GetFirebaseApp().Firestore(ctx)

	if err != nil {
		return err
	}

	firestoreClient = client

	return nil
}

func Client() *firestore.Client {
	return firestoreClient
}
