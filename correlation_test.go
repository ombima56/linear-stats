package main

import (
	"math"
	"testing"
)

func TestCorrelation(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{
			name:     "Perfect Positive Correlation",
			data:     []float64{1, 2, 3, 4, 5},
			expected: 1.0,
		},
		{
			name:     "Perfect Negative Correlation",
			data:     []float64{5, 4, 3, 2, 1},
			expected: -1.0,
		},
		{
			name:     "No Correlation",
			data:     []float64{1, 0, 1, 0, 1},
			expected: 0.0,
		},
		{
			name:     "Some Positive Correlation",
			data:     []float64{1, 3, 2, 4, 5},
			expected: 0.9,
		},
		{
			name:     "Empty Data",
			data:     []float64{},
			expected: 0.0,
		},
		{
			name:     "Single Value Data",
			data:     []float64{42},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Correlation(tt.data)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
