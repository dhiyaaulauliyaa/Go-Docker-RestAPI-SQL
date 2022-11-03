package config

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func SetupFirebase() *firebase.App {
	/* Get Service Account Key File */
	serviceAccountKeyFilePath, err := filepath.Abs("./credentials/firebaseServiceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	/* Create option from key file */
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	/* Init Firebase Admin SDK */
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Init Firebase Error")
	}

	return app

}

func SetupFirebaseAuth(app *firebase.App) *auth.Client {
	/* Init Auth */
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Init Firebase Auth Error")
	}
	return auth
}
