package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host                   string `mapstructure:"HOST"`
	Port                   string `mapstructure:"PORT"`
	Tls                    string `mapstructure:"TLS"`
	CertPath               string `mapstructure:"CERT_PATH"`
	KeyPath                string `mapstructure:"Key_PATH"`
	DbHost                 string `mapstructure:"DB_HOST"`
	DbPort                 string `mapstructure:"DB_PORT"`
	DbUser                 string `mapstructure:"DB_USER"`
	DbName                 string `mapstructure:"DB_NAME"`
	DbPassword             string `mapstructure:"DB_PASSWORD"`
	DbPreferSimpleProtocol bool   `mapstructure:"DB_PREFER_SIMPLE_PROTOCOL"`
	DbMigrate              bool   `mapstructure:"DB_MIGRATE"`
	DbSchema               string `mapstructure:"DB_SCHEMA"`
}

func Init() (config *Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
