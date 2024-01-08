package config

import (
	"github.com/spf13/viper"
)

// AppConfig holds the application configuration
type AppConfig struct {
	Server   ServerConfig
	Logging  LoggingConfig
	Database DatabaseConfig
	// ... other configurations
}

type ServerConfig struct {
	Port int
	Host string
}

type LoggingConfig struct {
	Level string
	File  string
}

type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	// ... other database configurations
}

// InitConfig initializes the application configuration using Viper
func InitConfig(configFile string) (AppConfig, error) {
	// Initialize Viper
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, err
	}

	// Create an AppConfig struct to hold the configuration
	var config AppConfig

	// Unmarshal the configuration into the AppConfig struct
	err = viper.Unmarshal(&config)
	if err != nil {
		return AppConfig{}, err
	}

	return config, nil
}
