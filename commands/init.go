package commands

import (
	"auto/models"
	"github.com/spf13/viper"
	"log"
)

var Config models.Config

func init() {
	viperConfig := viper.New()
	viperConfig.SetConfigName("config")
	viperConfig.SetConfigType("yaml")
	viperConfig.AddConfigPath(".")
	viperConfig.AddConfigPath("../")
	if err := viperConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("viper.ConfigFileNotFoundError: %v", err)

			return
		} else {
			log.Fatalf("viperConfig.ReadInConfig() err: %v", err)
			return
		}
	}
	// 映射到结构体
	if err := viperConfig.Unmarshal(&Config); err != nil {
		log.Fatalf("viperConfig.Unmarshal() err: %v", err)
	}
}
