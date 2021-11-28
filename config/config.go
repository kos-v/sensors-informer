package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Channels struct {
		TelegramBot struct {
			Enable bool
			Token  string
			ChatId int64 `yaml:"chatId"`
		} `yaml:"telegramBot"`
	}
	LmSensors struct {
		Command string
	} `yaml:"lmSensors"`
	Report struct {
		RepeatTimeout uint `yaml:"repeatTimeout"`
	}
	Sensors struct {
		CriticalTemperature float32 `yaml:"criticalTemperature"`
		PollingInterval     uint    `yaml:"pollingInterval"`
	}
}

func LoadConfig() (*Config, error) {
	data, err := ioutil.ReadFile(getConfig())
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	setDefaultValues(&config)

	return &config, err
}

func getConfig() string {
	return "./config.yml"
}

func setDefaultValues(config *Config) {
	if config.Sensors.PollingInterval < 1 {
		config.Sensors.PollingInterval = 1
	}
	if config.Report.RepeatTimeout < 60 {
		config.Report.RepeatTimeout = 60
	}
}
