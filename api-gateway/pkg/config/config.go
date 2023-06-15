package config

import "github.com/spf13/viper";

type Config struct {
	Port 			string `mapstructure:"PORT"`
	AuthSvcUrl 		string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl 	string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl 	string `mapstructure:"ORDER_SVC_URL"`
}

// 환경 설정 검수 및 설정
func LoadConfig() (c Config, err error) {

	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}