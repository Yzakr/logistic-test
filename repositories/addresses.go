package repositories

import (
	"github.com/Yzakr/logistic-test/models"
	"strings"
)

// Utility to obtain if Odd or Even
func CalculateOddOrEven(name string) string {
	// Trimming address street length
	word_length := len(strings.ReplaceAll(name, " ", ""))
	if word_length % 2 == 0 {
		return "odd"
	}
	return "even"
}

// Obtaining data for Addresses and attaching Odd or Even to the model
func GetAddressesData(filename string) ([]models.Address, error) {
	if filename == "" { 
		filename = "./repositories/data/addresses.txt"
	}
	data, err := GetDataFromFile(filename)
	if err != nil {
		return nil, err
	}
	var addressesData []models.Address
	for _, item := range data {
		temp_address := strings.Split(item, ", ")
		temp_street := strings.Split(temp_address[0], " ")
		temp_state := strings.Split(temp_address[2], " ")
		address_name := temp_street[1] + " " + temp_street[2]
		notes := ""
		if len(temp_street) > 3 {
			notes = temp_street[3] + " " + temp_street[4]
		}
		addressesData = append(addressesData, 
			models.Address {
				Name: address_name, 
				Notes: notes, 
				HouseNumber: 
				temp_street[0], 
				City: temp_address[1], 
				State: temp_state[0], 
				PostalCode: temp_state[1],
				OddOrEven: CalculateOddOrEven(address_name),
			})
	}
	return addressesData, nil
}