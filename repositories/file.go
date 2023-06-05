package repositories

import (
	"bufio"
	"os"
)

// Dynamic utility for retrieving data from files
func GetDataFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err != nil {
		return nil, err
	}
	return data, err
}
