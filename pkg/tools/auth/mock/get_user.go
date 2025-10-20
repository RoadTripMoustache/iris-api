package mock

import (
	"firebase.google.com/go/auth"
)

func (m *MockClient) GetUser(userID string) *auth.UserRecord {
	dataList := m.getDataFromMockFiles()

	var record *auth.UserRecord

	for _, data := range dataList {
		if data.ID == userID {
			record = &auth.UserRecord{
				UserInfo: &auth.UserInfo{
					DisplayName: data.DisplayName,
				},
				TenantID: data.TokenID,
			}
		}
	}

	return record
}

func (m *MockClient) GetUserByEmail(email string) *auth.UserRecord {
	dataList := m.getDataFromMockFiles()

	var record *auth.UserRecord

	for _, data := range dataList {
		if data.Email == email {
			record = &auth.UserRecord{
				UserInfo: &auth.UserInfo{
					DisplayName: data.DisplayName,
				},
				TenantID: data.TokenID,
			}
		}
	}

	return record
}
