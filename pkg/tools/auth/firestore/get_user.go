package firebase

import (
	"firebase.google.com/go/auth"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/logging"
)

// GetUser - Get the firebase user from its userID
func (f *FirebaseClient) GetUser(userID string) *auth.UserRecord {
	user, err := f.Client.GetUser(f.Context, userID)

	if err != nil {
		logging.Error(err, nil)
		return nil
	}

	return user
}

// GetUserByEmail - Get the firebase user from its email
func (f *FirebaseClient) GetUserByEmail(email string) *auth.UserRecord {
	user, err := f.Client.GetUserByEmail(f.Context, email)

	if err != nil {
		logging.Error(err, nil)
		return nil
	}

	return user
}
