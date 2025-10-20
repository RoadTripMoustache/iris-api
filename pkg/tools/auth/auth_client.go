package auth

import "firebase.google.com/go/auth"

type AuthClient interface {
	// VerifyIDToken - Verify the UID Token from the user
	VerifyIDToken(idToken string) *auth.Token

	// GetUser - Get the firebase user from its userID
	GetUser(userID string) *auth.UserRecord

	// GetUserByEmail - Get the firebase user from its email
	GetUserByEmail(email string) *auth.UserRecord

	// DeleteUser - Delete an user in firebase
	DeleteUser(userID string) error
}
