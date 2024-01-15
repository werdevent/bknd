package models

type Location struct {
	LocationType string    `json:"location_type" bson:"location_type"`
	Coordinates  []float64 `json:"coordinates" bson:"coordinates"`
}
