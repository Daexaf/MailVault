package config

import (
	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		AppName: viper.GetString("APP_NAME"),
		AppPort: viper.GetString("APP_PORT"),

		DBServer:   viper.GetString("DB_SERVER"),
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),

		StoragePath: viper.GetString("STORAGE_PATH"),
	}
	return cfg, nil
}
