package repositories

import (
	"github.com/Yzakr/logistic-test/models"
)

// Obtaining drivers data and model
func GetDriversData(filename string) ([]models.Driver, error) {
	if filename == "" {
		filename = "./repositories/data/drivers.txt"
	}
	data, err := GetDataFromFile(filename)
	if err != nil {
		return nil, err
	}
	var driversData []models.Driver
	for _, item := range data {
		driversData = append(driversData, models.Driver{Name: item})
	}
	return driversData, nil
}