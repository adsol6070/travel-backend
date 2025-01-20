package dynamodb

import (
	"context"
	"log"
	"travel-backend/customConfig"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDBClient() *dynamodb.Client {
	// Log: Starting to create DynamoDB client
	log.Println("Initializing DynamoDB client...")

	// Set up static credentials provider
	credentialsProvider := aws.NewCredentialsCache(
		credentials.NewStaticCredentialsProvider(
			customConfig.AppConfig.AWS.AccessKeyID,
			customConfig.AppConfig.AWS.SecretAccessKey,
			"",
		),
	)

	// Log: Attempting to load AWS configuration
	log.Println("Loading AWS configuration...")

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(customConfig.AppConfig.AWS.Region),
		config.WithCredentialsProvider(credentialsProvider),
	)

	if err != nil {
		// Log: Failed to load AWS configuration
		log.Printf("Failed to load AWS configuration: %v\n", err)
		return nil
	}

	// Log: AWS configuration loaded successfully
	log.Println("AWS configuration loaded successfully.")

	// Create DynamoDB client from the loaded config
	client := dynamodb.NewFromConfig(cfg)

	// Log: Attempting to test DynamoDB connection
	log.Println("Testing DynamoDB connection...")

	// Attempt a simple operation to verify the connection
	_, err = client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		// Log: Failed to connect to DynamoDB
		log.Printf("Failed to connect to DynamoDB: %v\n", err)
		return nil
	}

	// Log: Successfully connected to DynamoDB
	log.Println("Successfully connected to DynamoDB.")

	return client
}
