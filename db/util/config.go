package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl          string `mapstructure:"DB_URL"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	HttpServerAddr string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
