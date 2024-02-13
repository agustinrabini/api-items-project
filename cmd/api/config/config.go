package config

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	InternalBaseUrl = "http://127.0.0.1"
	LogRatio        = 100
	LogBodyRatio    = 100
)

var (
	Environment              = os.Getenv("GO_ENVIRONMENT")
	InternalBasePricesClient = os.Getenv("API_PRICES_INTERNAL_BASE_URL")
	InternalBaseShopsClient  = os.Getenv("API_SHOPS_INTERNAL_BASE_URL")
)

// Configuration estructura
type Configuration struct {
	APIRestServerHost string `mapstructure:"jopit_api_host"`
	APIRestServerPort string `mapstructure:"jopit_api_port"`
	APIRestUsername   string `mapstructure:"jopit_api_username"`
	APIRestPassword   string `mapstructure:"jopit_api_password"`
	APIBaseEndpoint   string `mapstructure:"jopit_api_base_endpoint"`
	LoggingPath       string `mapstructure:"jopit_api_logpath"`
	LoggingFile       string `mapstructure:"jopit_api_logfile"`
	LoggingLevel      string `mapstructure:"jopit_api_loglevel"`
	MongoUser         string
	MongoPassword     string
	MongoHost         string
	MongoDataBase     string
	MonngoConnString  string
}

// Config is package struct containing conf params
var ConfMap Configuration

func Load() {

	// Setting defaults if the config not read
	// API
	viper.SetDefault("jopit_api_host", "127.0.0.1")
	viper.SetDefault("jopit_api_port", ":8080")
	viper.SetDefault("jopit_api_username", "jopit")
	viper.SetDefault("jopit_api_password", "changeme")
	viper.SetDefault("jopit_api_base_endpoint", "http://nginx")
	// LOG
	viper.SetDefault("jopit_api_logpath", "/var/log/jopit")
	viper.SetDefault("jopit_api_logfile", "jopit_api.log")
	viper.SetDefault("jopit_api_loglevel", "trace")

	fmt.Println("\nSetting default values for: jopit_api_host, jopit_api_port, jopit_api_username, jopit_api_password, jopit_api_base_endpoint, jopit_api_logpath, jopit_api_logfile and jopit_api_loglevel")

	err := viper.Unmarshal(&ConfMap)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %+v", err)
	}
	spew.Dump(ConfMap)

	fmt.Printf("Loading database configuration...")

	if os.Getenv("MONGO_USERNAME") == "" {
		log.Fatal("MONGO_USERNAME is empty)")
	}
	ConfMap.MongoUser = os.Getenv("MONGO_USERNAME")

	if os.Getenv("MONGO_PASSWORD") == "" {
		log.Fatal("MONGO_PASSWORD is empty)")
	}
	ConfMap.MongoPassword = os.Getenv("MONGO_PASSWORD")

	if os.Getenv("MONGO_HOST") == "" {
		log.Fatal("MONGO_HOST is empty)")
	}
	ConfMap.MongoHost = os.Getenv("MONGO_HOST")

	if os.Getenv("MONGO_DATABASE") == "" {
		log.Fatal("MONGO_DATABASE is empty)")
	}
	ConfMap.MongoDataBase = os.Getenv("MONGO_DATABASE")

	if os.Getenv("API_SHOPS_INTERNAL_BASE_URL") == "" {
		log.Fatal("API_SHOPS_INTERNAL_BASE_URL is empty)")
	}

	if os.Getenv("API_PRICES_INTERNAL_BASE_URL") == "" {
		log.Fatal("API_PRICES_INTERNAL_BASE_URL is empty)")
	}

	fmt.Println("\n All good!!")
}

func IsProductionEnvironment() bool {
	return Environment == "production"
}
