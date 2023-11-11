package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentIntents struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
}

// CreatorDetails model that will allow the user to create events, if the model is empty the user will not be able to be a creator
type CreatorDetails struct {
	SellerID        string `json:"seller_id" bson:"seller_id"`
	SellerStatus    int32  `json:"seller_status" bson:"seller_status"`
	HasRequirements int32  `json:"has_requirements" bson:"has_requirements"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
	UpdatedAt       string `json:"updated_at" bson:"updated_at"`
}

// ConsumerDetails model that will allow the user to buy from other creators
type ConsumerDetails struct {
	BuyerID   string `json:"buyer_id" bson:"buyer_id"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}
