package models 

// Location is a 2 dimensional representation of the current location of the user 
type Location struct {
	LocationType string `json:"location_type" bson:"location_type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}