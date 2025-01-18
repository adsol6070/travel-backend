package models

type Meal struct {
	MealID           string `json:"mealID"`
	Description      string `json:"description"`
	AvailableClasses string `json:"availableClasses"`
}
