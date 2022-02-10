package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	Env                 string        `mapstructure:"ENV"`
	StripeKey           string        `mapstructure:"STRIPE_KEY"`
	StripeSecretKey     string        `mapstructure:"STRIPE_SECRET_KEY"`
	ZoomApiKey          string        `mapstructure:"ZOOM_API_KEY"`
	ZoomApiSecret       string        `mapstructure:"ZOOM_API_SECRET"`
	ZoomUserId          string        `mapstructure:"ZOOM_USER_ID"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
