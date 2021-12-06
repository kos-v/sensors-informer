package config

import (
	"fmt"
	notifySend "github.com/kos-v/sensors-informer/internal/notify-send"
	"github.com/kos-v/sensors-informer/internal/temperature"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ChannelsCommonConfig struct {
	Enable bool
}

type Config struct {
	Channels struct {
		File struct {
			ChannelsCommonConfig `yaml:",inline"`
			Path                 string
		}
		NotifySend struct {
			ChannelsCommonConfig `yaml:",inline"`
			Command              string
			Urgency              notifySend.Urgency
			ExpireTime           int `yaml:"expireTime"`
			Hint                 string
		} `yaml:"notifySend"`
		TelegramBot struct {
			ChannelsCommonConfig `yaml:",inline"`
			Token                string
			ChatId               int64 `yaml:"chatId"`
		} `yaml:"telegramBot"`
	}
	LmSensors struct {
		Command string
	} `yaml:"lmSensors"`
	Report struct {
		Format struct {
			TemperatureUnit temperature.Unit `yaml:"temperatureUnit"`
		}
		Message struct {
			Title struct {
				Text string
			}
		}
		RepeatTimeout uint `yaml:"repeatTimeout"`
	}
	Sensors struct {
		PollingInterval uint `yaml:"pollingInterval"`
		Temperature     struct {
			Critical temperature.Value
			Unit     temperature.Unit
		}
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
	if !temperature.IsSupportedUnit(config.Sensors.Temperature.Unit) {
		config.Sensors.Temperature.Unit = temperature.UnitCelsius
	}
	if !temperature.IsSupportedUnit(config.Report.Format.TemperatureUnit) {
		config.Report.Format.TemperatureUnit = temperature.UnitCelsius
	}
	if !notifySend.IsValidUrgency(config.Channels.NotifySend.Urgency) {
		config.Channels.NotifySend.Urgency = notifySend.UrgencyNormal
	}
}
