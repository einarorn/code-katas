package fibonacci

import (
	"errors"
	"math"
)

func CalculateByN(n int) (uint64, error) {
	value := math.Pow(1+math.Sqrt(5), float64(n)) / (math.Pow(2, float64(n)) * math.Sqrt(float64(5)))

	if math.IsNaN(value) {
		return 0, errors.New("fibonacci number to big to compute")
	}

	return uint64(math.Round(value)), nil
}
