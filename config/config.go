package config

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	Server struct {
		BaseUrl     string `yaml:"baseUrl"`
		Environment string `yaml:"environment"`
		Port        string `yaml:"port"`
		SecurityKey string `yaml:"securityKey"`
		Timeout     int    `yaml:"timeout"`
	} `yaml:"server"`
	DB struct {
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		DbName   string `yaml:"dbName"`
	} `yaml:"db"`
}

func NewConfig() *Config {
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	viper.SetConfigName("config.local")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error load config: %s", err)
		}
	}

	config := Config{}
	_ = viper.Unmarshal(&config)
	Cfg = &config
	return &config
}
