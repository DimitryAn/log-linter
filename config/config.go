package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AvailableLevels  []string `yaml:"available_levels"`
	AvailableLoggers []string `yaml:"available_loggers"`
	BannedKeywords   []string `yaml:"banned_sensitive_keywords"`
	AvailableSymbols []string `yaml:"available_special_symbols"`
}

func Load(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	var conf Config
	err = yaml.Unmarshal(data, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func (c *Config) Map() (levels, loggers, names map[string]struct{}, symboles map[rune]struct{}) {
	levels = make(map[string]struct{})

	for _, v := range c.AvailableLevels {
		levels[v] = struct{}{}
	}

	loggers = make(map[string]struct{})
	for _, v := range c.AvailableLoggers {
		loggers[v] = struct{}{}
	}

	names = make(map[string]struct{})
	for _, v := range c.BannedKeywords {
		names[v] = struct{}{}
	}

	symboles = make(map[rune]struct{})
	for _, v := range c.AvailableSymbols {
		if len(v) == 1 {
			symboles[rune(v[0])] = struct{}{}
		} else {
			fmt.Fprintf(os.Stderr, "loglinter: invalid format of specials symboles")
		}
	}

	return levels, loggers, names, symboles
}

func FromMap(m map[string]interface{}) (*Config, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
