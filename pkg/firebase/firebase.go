package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"mocerize-api/pkg/config"
	"os"
)

var (
	firebaseApp *firebase.App
)

//type Firebase struct {
//	*firebase.App
//}

func Setup(ctx context.Context) error {

	currentDirectory, err := os.Getwd()
	if err != nil {
		return err
	}

	options := option.WithCredentialsFile(currentDirectory + config.Get("FIREBASE_CREDENTIALS_FILE"))

	app, err := firebase.NewApp(ctx, nil, options)
	if err != nil {
		return err
	}
	firebaseApp = app

	return nil
}

func GetFirebaseApp() *firebase.App {
	return firebaseApp
}
