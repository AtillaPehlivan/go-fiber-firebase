package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
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

	var options option.ClientOption

	if config.Get("FIREBASE_CREDENTIALS_FILE") != "" {
		log.Println("read FIREBASE_CREDENTIALS_FILE")
		options = option.WithCredentialsFile(currentDirectory + config.Get("FIREBASE_CREDENTIALS_FILE"))
	} else {
		log.Println("read FIREBASE_CREDENTIALS")
		log.Println(config.Get("FIREBASE_CREDENTIALS"))
		options = option.WithCredentialsJSON([]byte(config.Get("FIREBASE_CREDENTIALS")))
	}

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
