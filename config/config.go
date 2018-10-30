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
var Config struct {
	Address        string              `yaml:"address"`
	StatsAddress   string              `yaml:"stats_address"`
	Interval       int                 `yaml:"interval"`
	RequestTimeout int                 `yaml:"request_timeout"`
	Metrics        []MetricsConfigItem `yaml:"metrics"`
}

const (
	// DefaultInterval represents default value of interval
	DefaultInterval = 5
	// DefaultAddress represents default value of address
	DefaultAddress = ":8080"
	// DefaultStatsAddress represents default value of mtproto_proxy stats URL
	DefaultStatsAddress = "http://localhost:2398/stats"
	// DefaultRequestTimeout represents default value of http request timeout in seconds
	DefaultRequestTimeout = 10
)

func readConfigFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return []byte(data), err
}

func parseConfig(configString []byte) error {
	return yaml.Unmarshal(configString, &Config)
}

// InitFromFile initializes config data from file
func InitFromFile(path string) error {
	if len(path) == 0 {
		fmt.Println("Using default config options.\n")
	} else {
		configData, err := readConfigFile(path)
		if err != nil {
			return fmt.Errorf("could not read config file \"%s\": %v", path, err)
		}
		err = parseConfig(configData)
		if err != nil {
			return fmt.Errorf("could not parse config: %v", err)
		}
	}

	if Config.Interval < 0 {
		return fmt.Errorf("scan interval can't be less than or equal to 0")
	} else if Config.Interval == 0 {
		Config.Interval = DefaultInterval
	}

	if Config.Address == "" {
		Config.Address = DefaultAddress
	}

	if Config.StatsAddress == "" {
		Config.StatsAddress = DefaultStatsAddress
	}

	if Config.RequestTimeout < 0 {
		return fmt.Errorf("http request timeout can't be less than or equal to 0")
	} else if Config.RequestTimeout == 0 {
		Config.RequestTimeout = DefaultRequestTimeout
	}

	return nil
}
