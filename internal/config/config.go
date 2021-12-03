package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Channels struct {
		File struct {
			Enable bool
			Path   string
		}
		NotifySend struct {
			Enable     bool
			Command    string
			ExpireTime int `yaml:"expireTime"`
			Hint       string
		} `yaml:"notifySend"`
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

func LoadConfig(specificConfig string) (*Config, error) {
	configPath, err := getConfigPath(specificConfig)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(configPath)
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

func getConfigPath(specificConfig string) (string, error) {
	localConfig := "./config.yml"
	globalConfig := "/etc/sensors-informer/config.yml"

	if specificConfig != "" {
		if isExists(specificConfig) {
			return specificConfig, nil
		}
	} else if isExists(localConfig) {
		return localConfig, nil
	} else if isExists(globalConfig) {
		return globalConfig, nil
	}

	return "", fmt.Errorf("config not found")
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil || !os.IsNotExist(err) {
		return true
	}

	return false
}

func setDefaultValues(config *Config) {
	if config.Sensors.PollingInterval < 1 {
		config.Sensors.PollingInterval = 1
	}
	if config.Report.RepeatTimeout < 60 {
		config.Report.RepeatTimeout = 60
	}
}
