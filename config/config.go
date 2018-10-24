package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// MetricsConfigItem represents metrics config item
type MetricsConfigItem struct {
	StatName string `yaml:"stat_name"`
	Name     string `yaml:"name"`
	Help     string `yaml:"help"`
}

// Config represents config file structure
type Config struct {
	Address      string              `yaml:"address"`
	StatsAddress string              `yaml:"stats_address"`
	Interval     int                 `yaml:"interval"`
	Metrics      []MetricsConfigItem `yaml:"metrics"`
}

const (
	// DefaultInterval represents default value of interval
	DefaultInterval = 5
	// DefaultAddress represents default value of address
	DefaultAddress = ":8080"
	// DefaultStatsAddress represents default value of mtproto_proxy stats URL
	DefaultStatsAddress = "http://localhost:2398/stats"
)

func readConfigFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return []byte(data), err
}

func parseConfig(configString []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(configString, &config)
	return config, err
}

// InitFromFile initializes config data from file
func InitFromFile(path string) (Config, error) {
	var config Config

	if len(path) == 0 {
		config = Config{}
		fmt.Println("Using default config options.\n")
	} else {
		configData, err := readConfigFile(path)
		if err != nil {
			return config, fmt.Errorf("could not read config file \"%s\": %v", path, err)
		}
		config, err = parseConfig(configData)
		if err != nil {
			return config, fmt.Errorf("could not parse config: %v", err)
		}
	}

	if config.Interval < 0 {
		return config, fmt.Errorf("scan interval can't be less than or equal to 0")
	} else if config.Interval == 0 {
		config.Interval = DefaultInterval
	}

	if config.Address == "" {
		config.Address = DefaultAddress
	}

	if config.StatsAddress == "" {
		config.StatsAddress = DefaultStatsAddress
	}

	return config, nil
}
