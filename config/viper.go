package config

import "github.com/spf13/viper"

func GetViper() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./config")

	err := config.ReadInConfig()
	return config, err
}
