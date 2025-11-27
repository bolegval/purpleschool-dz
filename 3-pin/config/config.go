package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в параметрах окружения")
	}

	return &Config{Key: key}
}
