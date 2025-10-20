package mock

func (m *MockClient) EncryptData(data string) (string, error) {
	return data, nil
}
