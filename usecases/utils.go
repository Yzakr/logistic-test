package usecases

import (
	"math"
)

func GetMaxNumber(items []float64) (float64, int) {
	var max_num float64
	var index_max_num int
	max_num = items[0]
 
	for index, element := range items {
		if max_num < element {
			max_num = element
			index_max_num = index
		}
	}
	return max_num, index_max_num
}

func GetCommonFactors(a, b int) []int {
	CF := []int{}
	for num := 2; num <= int(math.Min(float64(a), float64(b))); num++ {
		if a % num == 0 && b % num == 0 {
			CF = append(CF, num)
		}
	}
	return CF
}