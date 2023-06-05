package models

// Address Model Definition
type Address struct {
	Name string `json:"address_name"`
	// We can convert this to int if needed
	HouseNumber string `json:"house_number"`
	City string `json:"city"`
	State string `json:"state"`
	PostalCode string `json:"postal_code"`
	Notes string `json:"notes"`
	OddOrEven string `json:"oddoreven"`
}