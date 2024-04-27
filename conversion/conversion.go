package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floatForms := make([]float64, len(strings))
	for stringIndex, stringVal := range strings {
		price, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floatForms[stringIndex] = price
	}
	return floatForms, nil
}
