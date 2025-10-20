package mock

import (
	"encoding/json"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/utils"
	"os"
)

func (m *MockClient) getDataFromMockFiles() []MockAuthData {
	filePath := *config.GetConfigs().Firebase.Mock.DataFilePath

	if _, err := os.Stat(filePath); err == nil {
		f, err := os.Open(filePath)
		if err != nil {
			utils.ProcessError(err)
		}
		defer f.Close()

		var list = make([]MockAuthData, 0)

		decoder := json.NewDecoder(f)
		err = decoder.Decode(&list)
		if err != nil {
			utils.ProcessError(err)
		}

		return list
	}

	return nil
}
