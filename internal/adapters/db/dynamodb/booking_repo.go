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

type BookingRepo struct {
	client *dynamodb.Client
}

func NewBookingRepo(client *dynamodb.Client) *BookingRepo {
	return &BookingRepo{client: client}
}

func (r *BookingRepo) GetAllBookings() ([]models.Booking, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Bookings"),
	}

	result, err := r.client.Scan(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching bookings: %v", err)
		return nil, err
	}

	var bookings []models.Booking
	err = attributevalue.UnmarshalListOfMaps(result.Items, &bookings)
	if err != nil {
		log.Printf("Error unmarshalling bookings: %v", err)
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepo) GetBookingByID(id string) (*models.Booking, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Bookings"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	result, err := r.client.GetItem(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching booking: %v", err)
		return nil, err
	}

	var booking models.Booking
	err = attributevalue.UnmarshalMap(result.Item, &booking)
	if err != nil {
		log.Printf("Error unmarshalling booking: %v", err)
		return nil, err
	}

	return &booking, nil
}

// CreateBooking adds a new booking to the database
func (r *BookingRepo) CreateBooking(booking *models.Booking) error {
	if booking == nil {
		return errors.New("booking details are nil")
	}

	item, err := attributevalue.MarshalMap(booking)
	if err != nil {
		log.Printf("Error marshalling booking: %v", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Bookings"),
		Item:      item,
	}

	_, err = r.client.PutItem(context.Background(), input)
	if err != nil {
		log.Printf("Error inserting booking: %v", err)
		return err
	}

	return nil
}

// UpdateBookingStatus updates the status of a booking by ID
func (r *BookingRepo) UpdateBookingStatus(id string, status string) error {
	if id == "" || status == "" {
		return errors.New("id or status cannot be empty")
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Bookings"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET #status = :status"),
		ExpressionAttributeNames: map[string]string{
			"#status": "Status",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":status": &types.AttributeValueMemberS{Value: status},
		},
	}

	_, err := r.client.UpdateItem(context.Background(), input)
	if err != nil {
		log.Printf("Error updating booking status: %v", err)
		return err
	}

	return nil
}

// GetBookingsByUserID retrieves bookings for a specific user
func (r *BookingRepo) GetBookingsByUserID(userID string) ([]models.Booking, error) {
	if userID == "" {
		return nil, errors.New("userID cannot be empty")
	}

	input := &dynamodb.QueryInput{
		TableName:              aws.String("Bookings"),
		IndexName:              aws.String("UserID-index"), // Assuming you have a secondary index on UserID
		KeyConditionExpression: aws.String("UserID = :userID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userID": &types.AttributeValueMemberS{Value: userID},
		},
	}

	result, err := r.client.Query(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching bookings for userID %s: %v", userID, err)
		return nil, err
	}

	var bookings []models.Booking
	err = attributevalue.UnmarshalListOfMaps(result.Items, &bookings)
	if err != nil {
		log.Printf("Error unmarshalling bookings: %v", err)
		return nil, err
	}

	return bookings, nil
}

// DeleteBooking deletes a booking by ID
func (r *BookingRepo) DeleteBooking(id string) error {
	if id == "" {
		return errors.New("invalid booking ID")
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("Bookings"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	_, err := r.client.DeleteItem(context.Background(), input)
	if err != nil {
		log.Printf("Error deleting booking: %v", err)
		return err
	}

	return nil
}

// UpdateBooking updates an existing booking by ID
func (r *BookingRepo) UpdateBooking(id string, booking *models.Booking) (*models.Booking, error) {
	if id == "" || booking == nil {
		return nil, errors.New("invalid booking ID or booking details")
	}

	// Marshall the updated booking details into a DynamoDB item
	_, err := attributevalue.MarshalMap(booking)
	if err != nil {
		log.Printf("Error marshalling updated booking: %v", err)
		return nil, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Bookings"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET #status = :status, #userID = :userID, #otherAttributes = :otherAttributes"),
		ExpressionAttributeNames: map[string]string{
			"#status":          "Status",
			"#userID":          "UserID",
			"#otherAttributes": "OtherAttributes", // Add other fields you want to update
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			// ":status":          &types.AttributeValueMemberS{Value: booking.Status},
			":userID": &types.AttributeValueMemberS{Value: booking.UserID},
			// ":otherAttributes": &types.AttributeValueMemberM{Value: booking.OtherAttributes}, // Add other fields here
		},
	}

	_, err = r.client.UpdateItem(context.Background(), input)
	if err != nil {
		log.Printf("Error updating booking: %v", err)
		return nil, err
	}

	// Return the updated booking
	return booking, nil
}
