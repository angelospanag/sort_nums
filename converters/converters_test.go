package converters_test

import (
	"testing"

	"github.com/angelospanag/sort_nums/converters"
)

var intSliceToCSVStringTests = []struct {
	n        []int
	expected string
}{
	{[]int{3, 5, 0, 1, 2}, "3,5,0,1,2"},
	{[]int{}, ""},
	{[]int{3}, "3"},
}

func TestIntSliceToCSVString(t *testing.T) {
	for _, tt := range intSliceToCSVStringTests {
		actual := converters.IntSliceToCSVString(tt.n)

		if actual != tt.expected {
			t.Errorf("IntSliceToCSVString(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

var csvStringToIntSliceTests = []struct {
	n        string
	expected []int
}{
	{"3,5,0,1,2", []int{3, 5, 0, 1, 2}},
	{"", []int{}},
	{"3", []int{3}},
}

func TestCSVStringToIntSlice(t *testing.T) {
	for _, tt := range csvStringToIntSliceTests {
		actual, _ := converters.CSVStringToIntSlice(tt.n)

		for i, _ := range actual {
			if actual[i] != tt.expected[i] {
				t.Errorf("CSVStringToIntSlice(%d): expected %d, actual %d", tt.n, tt.expected[i], actual[i])
			}
		}
	}
}
