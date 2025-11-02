// Package config contains all the configuration and methods to define the config project structure and how to load them.
package config

type Config struct {
	Firebase FirebaseConfig `yaml:"firebase"`
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
	Images   ImageConfig    `yaml:"images"`
}
