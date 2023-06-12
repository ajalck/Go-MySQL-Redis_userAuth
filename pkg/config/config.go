package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	SqlDSN   string `mapstructure:"MYSQL_DSN"`
	RedisURL string `mapstructure:"REDIS_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/env")
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	return
}
