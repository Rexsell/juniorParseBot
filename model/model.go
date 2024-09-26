package model

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	ErrEmptyBotUrl    = errors.New("empty bot url")
	ErrEmptyToken     = errors.New("empty token")
	ErrEmptyKeywords  = errors.New("empty keywords")
	ErrEmptyForwardTo = errors.New("empty forward to")
)

type Config struct {
	BotUrl    string   `yaml:"bot_url" mapstructure:"bot_url"`
	Token     string   `yaml:"token" mapstructure:"token"`
	Keywords  []string `yaml:"keywords" mapstructure:"keywords"`
	ForwardTo []string `yaml:"forward_to" mapstructure:"forward_to"`
}

func InitConfig(path string) (*Config, error) {
	config := &Config{}
	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(cfgFile, &config)
	if err != nil {
		return nil, err
	}
	return config, err
}

func (c *Config) Validate() error {
	if c.BotUrl == "" {
		return ErrEmptyBotUrl
	}
	if c.Token == "" {
		return ErrEmptyToken
	}
	if len(c.Keywords) == 0 {
		return ErrEmptyKeywords
	}
	if len(c.ForwardTo) == 0 {
		return ErrEmptyForwardTo
	}
	return nil
}
