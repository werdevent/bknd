package models

// SIGNATURES
const (
	// HIGHER_SIGNATURE is a signature that will represent a higher signature and hierarchy
	HIGHER_SIGNATURE = "EDM111023JEHNHS666"
	// STANDARD_SIGNATURE is a signature that will represent a standard signature and hierarchy
	STANDARD_SIGNATURE = "EDM111O23JEHNSS265"
	// LOWER_SIGNATURE is a signature that will represent a low signature and hierarchy
	LOWER_SIGNATURE = "EDM111O23JEHNLS55"
)

// UKEY is a unique type that will encapsulate the unnique id of the user and basic information for the application to start running
type UKEY struct {
	Signature string `json:"signature" bson:"signature"`
	OwnerID string `json:"owner_id" bson:"owner_id" `
	LoggedIn bool `json:"logged_in" bson:"logged_in"`
	SellerID string `json:"seller_id"`
	BuyerID string `json:"buyer_id" bson:"buyer_id"`
}

