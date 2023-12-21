package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Standard models for an user account
type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name" `
	Email    string             `json:"email"  bson:"email"`
	Password string             `json:"password" bson:"password"`
	// Will be the type of user is
	Role            int32           `json:"role" bson:"role"`
	Gender          string          `json:"gender" bson:"gender"`
	Birthday        string          `json:"bday" bson:"bday"`
	Location        Location        `json:"location" bson:"location"`
	CreatorDetails  CreatorDetails  `json:"creator_details" bson:"creator_details"`
	ConsumerDetails ConsumerDetails `json:"consumer_details" bson:"consumer_details"`
	// security section
	RecoverCode int64  `json:"recover_code" bson:"recover_code"`
	ValidCode   bool   `json:"valid_code" bson:"valid_code"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
	UpdatedAt   string `json:"updated_at" bson:"updated_at"`
	LastUpdated string `json:"last_updated" bson:"last_updated"`
}

// BusinessDetails standard model for a business profile
type BusinessDetails struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	OwnerID          string             `json:"owner_id" bson:"owner_id"`
	BusinessNickname string             `json:"business_nickname" bson:"business_nickname"`
	BusinessName     string             `json:"business_name" bson:"business_name"`
	BusinessEmail    string             `json:"business_email" bson:"business_email"`
	CoverImage       string             `json:"cover_image" bson:"cover_image"`
	Banner           string             `json:"banner" bson:"banner"`
	Verified         bool               `json:"verified" bson:"verified"`
}
