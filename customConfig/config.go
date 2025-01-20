package customConfig

import (
	"log"

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
	// Automatically read environment variables
	viper.AutomaticEnv()

	// Read .env file if it exists
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("No .env file found, falling back to system environment variables: %v", err)
	}

	// Initialize AppConfig
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	// Set AWS Credentials from Environment Variables
	AppConfig.AWS.Region = viper.GetString("AWS_REGION")
	AppConfig.AWS.AccessKeyID = viper.GetString("AWS_ACCESS_KEY_ID")
	AppConfig.AWS.SecretAccessKey = viper.GetString("AWS_SECRET_ACCESS_KEY")

	log.Println("Configuration loaded successfully.")
}
