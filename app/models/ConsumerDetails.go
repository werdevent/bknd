package models

/*
	ConsumerDetails is a model that hold the information for the user to be able to buy tickets and ask for

refund to their accounts
*/
type ConsumerDetails struct {
	Buyer_ID   string `json:"buyer_id" bson:"buyer_id"`
	Creator_ID string `json:"creator_id" bson:"creator_id"`
	Created_At string `json:"created_at" bson:"created_at"`
	Updated_At string `json:"updated_at" bson:"updated_at"`
}
