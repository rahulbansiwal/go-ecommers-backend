package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl                string        `mapstructure:"DB_URL"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	HttpServerAddr       string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	LOGPATH              string        `mapstructure:"LOG_PATH"`
	LOGFILEPREFIX        string        `mapstructure:"LOG_FILE_PREFIX"`
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
