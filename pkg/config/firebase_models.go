package config

type FirebaseConfig struct {
	CredentialsFilePath string             `yaml:"credentials_file_path" envconfig:"FIREBASE_CREDENTIALS_FILE_PATH"`
	ProjectID           string             `yaml:"project_id" envconfig:"FIREBASE_PROJECT_ID"`
	Mock                FirebaseMockConfig `yaml:"mock"`
}

type FirebaseMockConfig struct {
	Enabled      bool    `yaml:"enabled" envconfig:"FIREBASE_MOCK_ENABLED"`
	DataFilePath *string `yaml:"data_file_path" envconfig:"FIREBASE_MOCK_DATA_PATH"`
}
