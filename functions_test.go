package stats

import "testing"
import "math"

func TestMean(t * testing.T) {
	var cases = []struct {
		in []float64
		expected float64
	}{
		{[]float64{1.0,2.0,3.0}, 2.0},
		{[]float64{1.0}, 1.0},
		{[]float64{10.0, -1.0, 3.0}, 4.0},

		// Special cases
		{[]float64{0.0}, 0.0},
		{[]float64{}, 0.0},
	}

	for _, test_case := range cases {
		mean := Mean(test_case.in)

		if ( mean != test_case.expected ) {
			t.Errorf("Mean(%v) == %v, expected %v", test_case.in, mean, test_case.expected)
		}
	}
}

func TestVariance(t * testing.T) {
	// First we test the interface. It should return an error when an array
	// containing less than 2 values is provided.
	_, err := Variance([]float64{})

	if err == nil {
		t.Errorf("Interface should have returned an error")
	}

	_, err = Variance([]float64{1,2})

	if err != nil {
		t.Errorf("Interface should have NOT returned an error")
	}

	// Run through some test cases
	cases := []struct {
		in []float64
		expected float64
	}{
		{[]float64{2.0,10.0,34.0,54.0}, 419.0},
		{[]float64{-20.0,1.23,100.32}, 2749.6048222222},
	}

	for _, test_case := range cases {
		variance, _ := Variance(test_case.in)

		if ( math.Ceil(variance) != math.Ceil(test_case.expected) ) {
			t.Errorf("Variance(%v) == %v, expected %v", test_case.in, variance, test_case.expected)
		}
	}
}
