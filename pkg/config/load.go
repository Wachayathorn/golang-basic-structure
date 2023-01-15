package config

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type ConfigInterface interface {
	Load(env string) error
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DDname   string `mapstructure:"dbname"`
	Timezone string `mapstructure:"timezone"`
}

type Config struct {
	Env        string   `mapstructure:"env"`
	Port       string   `mapstructure:"port"`
	Postgresql Postgres `mapstructure:"postgresql"`
}

type State string

const (
	StateLocal State = "local"
	StateDev   State = "dev"
)

func (c *Config) ParseState(env string) State {
	switch env {
	case "dev":
		return StateDev
	default:
		return StateLocal
	}
}

func (c *Config) Load(env string) (string, error) {
	vn := viper.New()
	vn.AddConfigPath(filepath.Join(".", "pkg", "config", "resource"))     // Set path to load config.yml
	vn.SetConfigName(fmt.Sprintf("config.%s", c.ParseState(string(env)))) // Load config file

	if err := vn.ReadInConfig(); err != nil { // Read config from file
		return "", err
	}

	if err := mapstructure.Decode(vn.AllSettings(), &c); err != nil {
		return "", err
	}

	return c.Port, nil
}
