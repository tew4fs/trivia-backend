package config

import (
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppHost  string
	Port     int
	LogLevel string
	Env      string
}

var (
	CONFIG_FILE = "./configs/config.yml"
)

func LoadConfigs() AppConfig {
	viper.SetConfigFile(CONFIG_FILE)
	viper.ReadInConfig()

	env := getEnvVarOrString("ENV", "local")

	envValues := viper.GetStringMap(env)

	cfg := AppConfig{}
	cfg.Env = env
	err := mapstructure.Decode(envValues, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func getEnvVarOrString(path string, defaultValue string) string {
	if value, ok := os.LookupEnv(path); ok {
		return value
	}
	return defaultValue
}
