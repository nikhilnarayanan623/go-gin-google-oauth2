package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	GoathClientID      string `mapstructure:"GOAUTH_CLIENT_ID"`
	GoauthClientSecret string `mapstructure:"GOAUTH_CLENT_SECRET"`
}

var envsNames = []string{

	"GOAUTH_CLIENT_ID", "GOAUTH_CLENT_SECRET",
}

var config Config

func LoadConfig() (Config, error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envsNames {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(config); err != nil {
		return config, err
	}

	return config, nil
}

func GetCofig() Config {
	return config
}
