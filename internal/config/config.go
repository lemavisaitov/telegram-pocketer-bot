package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lemavisaitov/telegram-pocketer-bot/pkg/logging"
	"sync"
)

type Config struct {
	ConsumerKey string `env:"CONSUMER_KEY"`
	Token       string `env:"TOKEN"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig(".env", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
