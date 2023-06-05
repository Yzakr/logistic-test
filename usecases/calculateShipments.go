package usecases

import (
	"github.com/Yzakr/logistic-test/models"
	"fmt"
)

// Main matching function
func ShipmentMatch(drivers []models.Driver, addresses []models.Address) ([]models.Shipment, error) {
	var matchShipments []models.Shipment
	// Giving priority to Address SS than Driver perspective SS to assign routes, but can be done otherwise
	for id, address := range addresses {
		scores := []float64{}
		for _, driver := range drivers {
			scores = append(scores, CalculateSuitabilityScore(driver.Name, address.OddOrEven, address.Name))
		}
		// If you want to print the scores for every address vs every driver
		fmt.Println(address.Name, scores)

		//Getting best driver for the address
		max_num, index := GetMaxNumber(scores)
		// Assigning driver and address to shipment
		matchShipments = append(matchShipments, models.Shipment{ID: id, SS: max_num, Driver: drivers[index], Address: address})
		// Pop driver from slice, already assignated
		drivers = append(drivers[:index], drivers[index+1:]...)
	}

	return matchShipments, nil
}