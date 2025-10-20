package firebase

import (
	"firebase.google.com/go/auth"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/logging"
)

// VerifyIDToken - Verify the UID Token from the user
func (f *FirebaseClient) VerifyIDToken(idToken string) *auth.Token {
	token, err := f.Client.VerifyIDToken(f.Context, idToken)

	if err != nil {
		logging.Error(err, nil)
		return nil
	}

	return token
}
