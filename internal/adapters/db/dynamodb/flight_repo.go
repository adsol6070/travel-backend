package dynamodb

import (
	"context"
	"errors"
	"log"
	"travel-backend/internal/core/domain/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type FlightRepo struct {
	client *dynamodb.Client
}

func NewFlightRepo(client *dynamodb.Client) *FlightRepo {
	return &FlightRepo{client: client}
}

// GetAllFlights retrieves all flights from the DynamoDB table
func (r *FlightRepo) GetAllFlights() ([]models.Flight, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Flights"),
	}

	result, err := r.client.Scan(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching flights: %v", err)
		return nil, err
	}

	var flights []models.Flight
	err = attributevalue.UnmarshalListOfMaps(result.Items, &flights)
	if err != nil {
		log.Printf("Error unmarshalling flights: %v", err)
		return nil, err
	}

	return flights, nil
}

// GetFlightByID retrieves a flight by its ID
func (r *FlightRepo) GetFlightByID(id string) (*models.Flight, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Flights"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	result, err := r.client.GetItem(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching flight: %v", err)
		return nil, err
	}

	var flight models.Flight
	err = attributevalue.UnmarshalMap(result.Item, &flight)
	if err != nil {
		log.Printf("Error unmarshalling flight: %v", err)
		return nil, err
	}

	return &flight, nil
}

// CreateFlight creates a new flight record in DynamoDB
func (r *FlightRepo) CreateFlight(flight *models.Flight) error {
	if flight == nil {
		return errors.New("flight details cannot be nil")
	}

	// Marshal the flight struct into a map to store in DynamoDB
	item, err := attributevalue.MarshalMap(flight)
	if err != nil {
		log.Printf("Error marshalling flight: %v", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Flights"),
		Item:      item,
	}

	_, err = r.client.PutItem(context.Background(), input)
	if err != nil {
		log.Printf("Error creating flight: %v", err)
		return err
	}

	return nil
}

// UpdateFlight updates an existing flight by ID
func (r *FlightRepo) UpdateFlight(id string, flight *models.Flight) (*models.Flight, error) {
	if id == "" || flight == nil {
		return nil, errors.New("invalid flight ID or flight details")
	}

	// Marshal the updated flight details into a DynamoDB item
	_, err := attributevalue.MarshalMap(flight)
	if err != nil {
		log.Printf("Error marshalling updated flight: %v", err)
		return nil, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Flights"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET #flightName = :flightName, #destination = :destination, #otherAttributes = :otherAttributes"),
		ExpressionAttributeNames: map[string]string{
			"#flightName":      "FlightName",
			"#destination":     "Destination",
			"#otherAttributes": "OtherAttributes", // Add other fields to update
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			// ":flightName":      &types.AttributeValueMemberS{Value: flight.FlightName},
			":destination": &types.AttributeValueMemberS{Value: flight.Destination},
			// ":otherAttributes": &types.AttributeValueMemberM{Value: flight.OtherAttributes}, // Add other fields here
		},
	}

	_, err = r.client.UpdateItem(context.Background(), input)
	if err != nil {
		log.Printf("Error updating flight: %v", err)
		return nil, err
	}

	// Return the updated flight
	return flight, nil
}

// DeleteFlight deletes a flight by ID
func (r *FlightRepo) DeleteFlight(id string) error {
	if id == "" {
		return errors.New("invalid flight ID")
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("Flights"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	_, err := r.client.DeleteItem(context.Background(), input)
	if err != nil {
		log.Printf("Error deleting flight: %v", err)
		return err
	}

	return nil
}

// GetFlightBookings retrieves all bookings for a given flight ID.
func (r *FlightRepo) GetFlightBookings(flightID string) ([]models.Booking, error) {
	if flightID == "" {
		return nil, errors.New("flight ID cannot be empty")
	}

	// Define the DynamoDB query input to fetch bookings for the specified flight ID
	input := &dynamodb.QueryInput{
		TableName:              aws.String("Bookings"), // Assuming there's a 'Bookings' table
		KeyConditionExpression: aws.String("FlightID = :flightID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":flightID": &types.AttributeValueMemberS{Value: flightID},
		},
	}

	// Execute the query
	result, err := r.client.Query(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching bookings for flight %s: %v", flightID, err)
		return nil, err
	}

	// Unmarshal the result into a list of bookings
	var bookings []models.Booking
	err = attributevalue.UnmarshalListOfMaps(result.Items, &bookings)
	if err != nil {
		log.Printf("Error unmarshalling bookings for flight %s: %v", flightID, err)
		return nil, err
	}

	return bookings, nil
}
