package config

import (
	"strings"

	"github.com/spf13/viper"
)

var Config = &config{}

type config struct {
	General generalConfig
	DB      dbConfig
}

type generalConfig struct {
	Port string
}

type dbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSL      string
}

func init() {
	appName := "example"
	viper.SetConfigName(appName)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetEnvPrefix(appName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.Unmarshal(Config)
}
