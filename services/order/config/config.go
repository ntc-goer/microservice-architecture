package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

type Host struct {
	GRPCHost    string `json:"grpc_host"`
	GRPCPort    string `json:"grpc_port"`
	HTTPHost    string `json:"http_host"`
	HTTPPort    string `json:"http_port"`
	ServiceId   string `json:"service_id"`
	ServiceName string `json:"service_name"`
}

type Config struct {
	Host Host
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./services/order/config")
	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "json"
	}); err != nil {
		log.Fatalf("Unable to unmarshal config into struct: %v", err)
	}
	return &cfg, nil
}
