package usecases

import "testing"

func TestCalculateSuitabilityScore(t *testing.T) {
    SS := CalculateSuitabilityScore("Rolando Valenzuela", "even", "Dessie Lights")
    if SS != 9.0 {
        t.Errorf("SS = %f; want 9.0", SS)
    }
}