package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
The model of user represent all the information the user or consumer needs to start using the app, once the
user agree or apply to be a creator, it will still be able to be a consumer and have a sort of like a second profile
but this time is for event creation
*/

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id" `
	Name          string             `json:"name" bson:"name"`
	Email         string             `json:"email" bson:"email" `
	Password      string             `json:"password" bson:"password" `
	Profile_Image string             `json:"profile_image" bson:"profile_image"`

	// INFORMATIONAL INFO ABOUT THE USER
	Role             int32           `json:"role" bson:"role"`
	Gender           string          `json:"gender" bson:"gender"`
	Birthday         string          `json:"birthday" bson:"birthday"`
	Default_Location Location        `json:"default_location" bson:"default_location"`
	Consumer         ConsumerDetails `json:"consumer_details" bson:"consumer_details"`

	// SECURITY DATA
	Recovery_Code string `json:"recovery_code" bson:"recovery_code"`
	Valid_Code    bool   `json:"valid_code" bson:"valid_code"`
	Created_At    string `json:"created_at" bson:"created_at"`
	Updated_At    string `json:"updated_at" bson:"updated_at"`
	Last_Updated  string `json:"last_updated" bson:"last_updated"`
	Verified      bool   `json:"verified" bson:"verified"`
}
