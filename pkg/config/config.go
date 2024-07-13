package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"port"`
	DB   *dateBase
}

type dateBase struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
	SSLMode  string `mapstructure:"sslMode"`
}

func InitConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
