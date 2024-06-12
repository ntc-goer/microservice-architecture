package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	GRPCPort            string `json:"grpc_port"`
	OrderServiceName    string `json:"order_service_name"`
	ConsumerServiceName string `json:"consumer_service_name"`
	LBServiceHost       string `json:"lb_service_host"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
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
	// Accept to override os env if you need
	if port := os.Getenv("GRPC_PORT"); port != "" {
		cfg.GRPCPort = port
	}
	return &cfg, nil
}
