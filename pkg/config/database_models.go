package config

type DatabaseConfig struct {
	Mock  DatabaseMockConfig  `yaml:"mock"`
	Mongo DatabaseMongoConfig `yaml:"mongo"`
}

type DatabaseMockConfig struct {
	Enabled        bool    `yaml:"enabled" envconfig:"DATABASE_MOCK_ENABLED"`
	DataFolderPath *string `yaml:"data_folder_path" envconfig:"DATABASE_MOCK_DATA_PATH"`
}

type DatabaseMongoConfig struct {
	URI        *string `yaml:"uri" envconfig:"DATABASE_MONGO_URI"`
	Name       *string `yaml:"name" envconfig:"DATABASE_MONGO_NAME"`
	PrivateKey *string `yaml:"private_key" envconfig:"DATABASE_MONGO_PRIVATE_KEY"`
	PublicKey  *string `yaml:"public_key" envconfig:"DATABASE_MONGO_PUBLIC_KEY"`
}
