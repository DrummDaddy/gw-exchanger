package config

import "github.com/spf13/viper"

type Config struct {
	GRPCPort string
	DB       DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

//LoadConfig читает конфигурации из config.env
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
