package section1

import (
	"testing"

	"github.com/Doniblaze/exercise-repo/internal/section1"
)

func TestSumAllEvenNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"EmptySlice", []int{}, 0},
		{"NoEvenNumbers", []int{1, 3, 5, 7}, 0},
		{"OnlyEvenNumbers", []int{2, 4, 6, 8}, 20},
		{"MixedNumbers", []int{1, 2, 3, 4, 5, 6}, 12},
		{"LargeNumbers", []int{1000000, 999999, 888888, 777777}, 1888888},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := section1.SumAllEvenNumber(tt.input)
			if result != tt.expected {
				t.Errorf("Expected sum of even numbers for %s: %d, but got: %d", tt.name, tt.expected, result)
			}
		})
	}
}
