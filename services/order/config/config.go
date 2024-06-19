package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"os"
)

type Config struct {
	GRPCPort                string `json:"grpc_port"`
	DatabaseServerHost      string `json:"database_server_host"`
	DatabaseServerPort      string `json:"database_server_port"`
	DatabaseName            string `json:"database_name"`
	DatabaseUser            string `json:"database_user"`
	DatabasePwd             string `json:"database_pwd"`
	OrderServiceName        string `json:"order_service_name"`
	ConsumerServiceName     string `json:"consumer_service_name"`
	LBServiceHost           string `json:"lb_service_host"`
	QueueAddress            string `json:"queue_address"`
	QueueOrchestrator       string `json:"queue_orchestrator"`
	QueueCreateOrderSubject string `json:"queue_create_order_subject"`
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
	remoteProviderEndpoint := getEnv("REMOTE_PROVIDER_ENDPOINT", "localhost:8500")
	remoteProviderPath := getEnv("REMOTE_PROVIDER_PATH", "env/orders")

	switch appEnv {
	case "development", "production":
		vp.AddRemoteProvider("consul", remoteProviderEndpoint, remoteProviderPath)
		vp.SetConfigType("json") // Need to explicitly set this to json
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
