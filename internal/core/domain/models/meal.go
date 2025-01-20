package models

type Meal struct {
	MealID           string `json:"mealID" dynamodbav:"mealID"`
	Description      string `json:"description" dynamodbav:"description"`
	AvailableClasses string `json:"availableClasses" dynamodbav:"availableClasses"`
}
