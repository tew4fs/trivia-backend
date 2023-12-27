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
	configFile = "./configs/config.yml"
	Config     = &AppConfig{}
)

func LoadConfigs() {
	viper.SetConfigFile(configFile)
	viper.ReadInConfig()

	env := getEnvVarOrString("ENV", "local")

	envValues := viper.GetStringMap(env)

	Config.Env = env
	err := mapstructure.Decode(envValues, Config)
	if err != nil {
		panic(err)
	}

}

func getEnvVarOrString(path string, defaultValue string) string {
	if value, ok := os.LookupEnv(path); ok {
		return value
	}
	return defaultValue
}
