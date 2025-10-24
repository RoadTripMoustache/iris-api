package config

import (
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/utils"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	// defaultConfigPath - Default value for the environment variable CONFIG_FILE_PATH
	defaultConfigPath = "config.yaml"
)

var config Config

func LoadConfig() {
	readFile(&config)
	readEnv(&config)
	setDefaultValues(&config)
	checkConfig(&config)
}

func GetConfigs() Config {
	return config
}

func readFile(cfg *Config) {
	configFilePath, isDefined := os.LookupEnv(enum.ConfigFilePath)
	if !isDefined {
		configFilePath = defaultConfigPath
	}

	if _, err := os.Stat(configFilePath); err == nil {
		f, err := os.Open(configFilePath)
		if err != nil {
			utils.ProcessError(err)
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(cfg)
		if err != nil {
			utils.ProcessError(err)
		}
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("envconfig", cfg)
	if err != nil {
		utils.ProcessError(err)
	}
}

func setDefaultValues(cfg *Config) {
	if cfg.Server.Expose == "" {
		cfg.Server.Expose = ":8080"
	}
	if cfg.Server.MetricsExpose == "" {
		cfg.Server.MetricsExpose = ":2121"
	}
	if cfg.Server.ImagesDir == "" {
		cfg.Server.ImagesDir = "tmp/images"
	}
	if cfg.Server.ImagesBaseURL == "" {
		cfg.Server.ImagesBaseURL = "/images"
	}
}

func checkConfig(cfg *Config) {
	if cfg == nil {
		utils.ProcessError(fmt.Errorf("empty config"))
	} else {
		//  AUTH
		if cfg.Firebase.Mock.Enabled && cfg.Firebase.Mock.DataFilePath == nil {
			utils.ProcessError(fmt.Errorf("mock Firebase data path should not be empty if the mock Firebase is enabled"))
		}
		if !cfg.Firebase.Mock.Enabled && cfg.Firebase.ProjectID == "" {
			utils.ProcessError(fmt.Errorf("firebase ProjectID must be defined"))
		}
		if !cfg.Firebase.Mock.Enabled && cfg.Firebase.CredentialsFilePath == "" {
			utils.ProcessError(fmt.Errorf("firebase CredentialsFilePath must be defined"))
		}

		// DATABASE
		if cfg.Database.Mock.Enabled && cfg.Database.Mock.DataFolderPath == nil {
			utils.ProcessError(fmt.Errorf("mock DB data path should not be empty if the mock DB is enabled"))
		}
		// --- mongo
		if !cfg.Database.Mock.Enabled && cfg.Database.Mongo.URI == nil {
			utils.ProcessError(fmt.Errorf("MongoDB URI must be defined"))
		}
		if !cfg.Database.Mock.Enabled && cfg.Database.Mongo.Name == nil {
			utils.ProcessError(fmt.Errorf("MongoDB Name must be defined"))
		}
	}
}
