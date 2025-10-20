package mock

func (m *MockClient) DecryptData(data string) (string, error) {
	return data, nil
}
