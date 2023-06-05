package models

// Shipment Model Definition
type Shipment struct {
	ID int `json:"ID"`
	SS float64 `json:"ss"`
	Driver
	Address
}