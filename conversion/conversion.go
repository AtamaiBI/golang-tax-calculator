package conversion

import (
	"strconv"
"errors"
"math")

// Receives a slice of strings and returns a slice of floats and an error
func StringsToFloats(strings []string) ([]float64, error) {

	var floats []float64

		for _, stringVal := range strings {

		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}
floats = append(floats, floatVal)
	}

	return floats, nil
	}

	func RoundToTwo(num float64) float64 {
		return math.Round(num*100) / 100
	}


