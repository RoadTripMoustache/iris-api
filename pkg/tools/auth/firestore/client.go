// Package firebase contains all methods implementation to communicate with Firebase authentication tools.
package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/logging"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	Client  *auth.Client
	Context context.Context
}

// New - Generates a FirebaseClient instance from the credentials file and the project id.
func New() *FirebaseClient {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.GetConfigs().Firebase.CredentialsFilePath)
	conf := &firebase.Config{ProjectID: config.GetConfigs().Firebase.ProjectID}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		logging.Fatal(err, nil)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		logging.Fatal(err, map[string]interface{}{"package": "firebase", "class": "FirebaseClient"})
	}

	return &FirebaseClient{
		Client:  client,
		Context: ctx,
	}
}
