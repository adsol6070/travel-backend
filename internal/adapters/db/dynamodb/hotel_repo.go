package dynamodb

import (
	"context"
	"log"
	"travel-backend/internal/core/domain/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type HotelRepo struct {
	client *dynamodb.Client
}

func NewHotelRepo(client *dynamodb.Client) *HotelRepo {
	return &HotelRepo{client: client}
}

func (r *HotelRepo) GetAllHotels() ([]models.Hotel, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String("Hotels"),
	}

	result, err := r.client.Scan(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching hotels: %v", err)
		return nil, err
	}

	var hotels []models.Hotel
	err = attributevalue.UnmarshalListOfMaps(result.Items, &hotels)
	if err != nil {
		log.Printf("Error unmarshalling hotels: %v", err)
		return nil, err
	}

	return hotels, nil
}

func (r *HotelRepo) GetHotelByID(id string) (*models.Hotel, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Hotels"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	result, err := r.client.GetItem(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching hotel: %v", err)
		return nil, err
	}

	var hotel models.Hotel
	err = attributevalue.UnmarshalMap(result.Item, &hotel)
	if err != nil {
		log.Printf("Error unmarshalling hotel: %v", err)
		return nil, err
	}

	return &hotel, nil
}

func (r *HotelRepo) CreateHotel(hotel *models.Hotel) error {
	av, err := attributevalue.MarshalMap(hotel)
	if err != nil {
		log.Printf("Error marshalling hotel: %v", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Hotels"),
		Item:      av,
	}

	_, err = r.client.PutItem(context.Background(), input)
	return err
}

func (r *HotelRepo) GetHotelBookings(hotelID string) ([]models.Booking, error) {
	// Create a query input to fetch bookings by hotelID
	input := &dynamodb.QueryInput{
		TableName:              aws.String("Bookings"),
		KeyConditionExpression: aws.String("HotelID = :hotelID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hotelID": &types.AttributeValueMemberS{Value: hotelID},
		},
	}

	result, err := r.client.Query(context.Background(), input)
	if err != nil {
		log.Printf("Error fetching bookings for hotel %s: %v", hotelID, err)
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

func (r *HotelRepo) UpdateHotel(id string, hotel *models.Hotel) (*models.Hotel, error) {
	// Marshal the updated hotel details
	updatedValues, err := attributevalue.MarshalMap(hotel)
	if err != nil {
		log.Printf("Error marshalling updated hotel: %v", err)
		return nil, err
	}

	// Prepare the update input
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Hotels"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		AttributeUpdates: map[string]types.AttributeValueUpdate{},
	}

	// Add updated attributes to the input
	for key, value := range updatedValues {
		input.AttributeUpdates[key] = types.AttributeValueUpdate{
			Value:  value,
			Action: types.AttributeActionPut,
		}
	}

	// Perform the update
	_, err = r.client.UpdateItem(context.Background(), input)
	if err != nil {
		log.Printf("Error updating hotel %s: %v", id, err)
		return nil, err
	}

	// Return the updated hotel (you can also fetch it again if needed)
	return hotel, nil
}

func (r *HotelRepo) DeleteHotel(id string) error {
	// Prepare the delete input
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("Hotels"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	// Perform the deletion
	_, err := r.client.DeleteItem(context.Background(), input)
	if err != nil {
		log.Printf("Error deleting hotel %s: %v", id, err)
		return err
	}

	return nil
}
