// Package mock contains all the methods to simulate the calls to an authentication service.
package mock

import (
	"firebase.google.com/go/auth"
)

// VerifyIDToken - Verify the UID Token from the user
func (m *MockClient) VerifyIDToken(idToken string) *auth.Token {
	dataList := m.getDataFromMockFiles()

	var token *auth.Token

	for _, data := range dataList {
		if data.TokenID == idToken {
			token = &auth.Token{
				Subject: data.ID,
				Firebase: auth.FirebaseInfo{
					Identities: map[string]interface{}{
						"email": []interface{}{
							data.Email,
						},
					},
				},
			}
		}
	}

	return token
}
