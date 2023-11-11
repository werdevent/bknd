package models 

import "go.mongodb.org/mongo-driver/bson/primitive"

// SuperAdmin is an exclusive model that have access to all parts of the app
type SuperAdmin struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name" `
	Email string `json:"email"  bson:"email"`
	Password string `json:"password" bson:"password"`
	Role int32`json:"role" bson:"role"`
	RecoverCode int64 `json:"recover_code" bson:"recover_code"`
	ValidCode bool `json:"valid_code" bson:"valid_code"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	LastUpdated string `json:"last_updated" bson:"last_updated"`
	UKEY UKEY `json:"ukey" bson:"ukey"` 
}