package config

import (
	_ "system-monitor-backend/models"

	"github.com/joho/godotenv"
)

type Config struct {
	settingsPairs map[string]string
}

func NewConfig() (Config, error) {
	settingsPairs, err := godotenv.Read("server.env")
	return Config{
		settingsPairs: settingsPairs,
	}, err
}

func (cfg Config) GetSetting(key string) string {
	return cfg.settingsPairs[key]
}
