package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	Enable        bool   `mapstructure:"ENABLE"`
	Host          string `mapstructure:"REDIS_HOST"`
	Port          int    `mapstructure:"REDIS_PORT"`
	Password      string `mapstructure:"PASSWORD"`
	Database      int    `mapstructure:"DATABASE"`
	ExpiryTime    int    `mapstructure:"EXPIRY_TIME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)

	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	return

}
