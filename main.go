package main

import (
	"github.com/Yzakr/logistic-test/usecases"
	"github.com/Yzakr/logistic-test/repositories"
	"fmt"
	"encoding/json"
)

func main() {
	var addressFile string
    print("Address File: ")
    _, err := fmt.Scanln(&addressFile)
    if err != nil {
            fmt.Println("Error: ", err)
    }

    var driversFile string
    print("Drivers File: ")
    _, err = fmt.Scanln(&driversFile)
    if err != nil {
            fmt.Println("Error: ", err)
    }

	addresses, err := repositories.GetAddressesData(addressFile)
	drivers, err := repositories.GetDriversData(driversFile)

	if err != nil {
		panic(err)	
	}

	shipments, err := usecases.ShipmentMatch(drivers, addresses)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(shipments)
	fmt.Println(string(b))
}