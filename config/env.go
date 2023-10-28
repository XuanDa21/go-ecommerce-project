package env

import (
	"fmt"

	"github.com/spf13/viper"
)

type appConf struct {
	ENV             string `mapstructure:"ENV"`
	OV_SYNCHRO_MODE string `mapstructure:"OV_SYNCHRO_MODE"`
	OV_SYNCHRO_CRON string `mapstructure:"OV_SYNCHRO_CRON"`
	OV_BASE_URL     string `mapstructure:"OV_BASE_URL"`
	OV_ACCESS_TOKEN string `mapstructure:"OV_ACCESS_TOKEN"` //TODO: dynamic token
	BACKEND_API_KEY string `mapstructure:"BACKEND_API_KEY"`
	BASE_URL        string `mapstructure:"BASE_URL"` //OVNA backend
	LOG_PATH        string `mapstructure:"LOG_PATH"`
}

var Conf appConf

func LoadEnv() (err error) {
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Conf)
	fmt.Println(Conf.BACKEND_API_KEY)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return nil
}
