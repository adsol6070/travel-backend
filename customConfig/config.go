package customConfig

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AWS struct {
		Region          string
		AccessKeyID     string
		SecretAccessKey string
		DynamoDB        struct {
			Endpoint    string
			TablePrefix string
		}
	}
}

var AppConfig *Config

func LoadConfig() {
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(("./config"))

	viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file: %v", err)
	// }

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config file: %v", err)
	}

	// Set AWS Credentials from Environment Variables
	AppConfig.AWS.Region = os.Getenv("AWS_REGION")
	AppConfig.AWS.AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	AppConfig.AWS.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	log.Println("Configuration loaded successfully.")
}
