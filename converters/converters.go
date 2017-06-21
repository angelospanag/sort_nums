package converters

import (
	"strconv"
	"strings"
)

// IntSliceToCSVString converts a slice of integers to a 'comma separated values' string
func IntSliceToCSVString(ints []int) (string, error) {

	b := make([]string, len(ints))
	for i, v := range ints {
		b[i] = strconv.Itoa(v)
	}

	result := strings.Join(b, ",")

	return result, nil
}

// CSVStringToIntSlice converts a 'comma separated values' string to a slice of integers
func CSVStringToIntSlice(s string) ([]int, error) {

	sanitizedString := strings.Replace(s, ",", " ", -1)
	trimmedString := strings.Trim(sanitizedString, " ")
	numbersSlice := strings.Split(trimmedString, " ")

	ints := []int{}

	for _, number := range numbersSlice {
		i, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}

	return ints, nil
}
