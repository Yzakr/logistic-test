package usecases

import "testing"

func TestGetMaxNumber(t *testing.T) {
    max, _ := GetMaxNumber([]float64{3.0,5.3,7.2})
    if max != 7.2 {
        t.Errorf("MaxNumber = %f; want 7.2", max)
    }
}

func TestGetMaxNumberIndex(t *testing.T) {
    _, index := GetMaxNumber([]float64{3.0,5.3,7.2})
    if index != 2 {
        t.Errorf("MaxNumberIndex = %d; want 2", index)
    }
}

func TestGetCommonFactors(t *testing.T) {
	CF := GetCommonFactors(15,30)
    if CF[0] != 3 {
        t.Errorf("Get CF = %d; want 3", CF[0])
    }
}
