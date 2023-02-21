package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct{
	DBHost string `mapstructure:"POSTGRES_HOST"`
	DBUsername string `mapstructure:"POSTGRES_USERNAME"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"POSTGRES_NAME"`
	DBPort string `mapstructure:"POSTGRES_PORT"`

	ServerPort string `mapstructure:"PORT"`

	TokenSecret string `mapstructure:"TOKEN_SECRET"`
	TokenExpirationTime time.Duration `mapstructure:"TOEKN_EXPIRED_IN"`
	TokenMaxAge int `mapstructure:"TOKEN_MAXAGE"`
}


func LoadConfig(path string) (config Config, err error){

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil{
		return
	}

	err = viper.Unmarshal(&config)
	return

}