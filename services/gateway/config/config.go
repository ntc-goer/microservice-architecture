package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

type Database struct {
	ServerHost string `json:"server_host"`
	ServerPort string `json:"server_port"`
	DBName     string `json:"db_name"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
}

type Queue struct {
	Orchestrator string `json:"orchestrator"`
	Mail         string `json:"mail"`
}

type Broker struct {
	Address string  `json:"address"`
	Subject Subject `json:"subject"`
	Queue   Queue   `json:"queue"`
}

type Subject struct {
	CreateOrder string `json:"create_order"`
	TestSubject string `json:"test_subject"`
	SendMail    string `json:"send_mail"`
}

type Service struct {
	LBServiceHost           string `json:"lb_service_host"`
	OrderServiceName        string `json:"order_service_name"`
	ConsumerServiceName     string `json:"consumer_service_name"`
	OrchestratorServiceName string `json:"orchestrator_service_name"`
	MailServiceName         string `json:"mail_service_name"`
	GatewayServiceName      string `json:"gateway_service_name"`
}

type Config struct {
	ServicePort string   `json:"service_port"`
	Database    Database `json:"database"`
	Service     Service  `json:"service"`
	Broker      Broker   `json:"broker"`
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
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to unmarshal config into struct: %v", err)
	}
	return &cfg, nil
}
