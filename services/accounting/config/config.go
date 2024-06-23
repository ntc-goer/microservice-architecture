package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"os"
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
	AccountingServiceName   string `json:"accounting_service_name"`
}

type Config struct {
	ServicePort string   `json:"service_port"`
	Database    Database `json:"database"`
	Service     Service  `json:"service"`
	Broker      Broker   `json:"broker"`
}

func getEnv(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return val
}

func Load() (*Config, error) {
	vp := viper.New()
	appEnv := getEnv("APP_ENV", "local")

	switch appEnv {
	case "development", "production":
		remoteProviderEndpoint := getEnv("REMOTE_PROVIDER_ENDPOINT", "localhost:8500")
		remoteProviderPath := getEnv("REMOTE_PROVIDER_PATH", "env/orders")

		vp.AddRemoteProvider("consul", remoteProviderEndpoint, remoteProviderPath)
		vp.SetConfigType("yaml") // Need to explicitly set this to json
		if err := vp.ReadRemoteConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
			return nil, err
		}
	default:
		vp.SetConfigName("config")
		vp.SetConfigType("yaml")
		vp.AddConfigPath("./config")
		// Read the configuration file
		if err := vp.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
			return nil, err
		}
	}
	var cfg Config
	if err := vp.Unmarshal(&cfg, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "json"
	}); err != nil {
		log.Fatalf("Unable to unmarshal config into struct: %v", err)
	}
	return &cfg, nil
}
