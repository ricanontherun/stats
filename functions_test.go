package stats

import "testing"

func TestMean(t * testing.T) {
	var cases = []struct {
		in []float64
		expected float64
	}{
		{[]float64{1.0,2.0,3.0}, 2.0},
		{[]float64{0.0}, 0.0},
		{[]float64{1.0}, 1.0},
		{[]float64{}, 0.0},
		{[]float64{10.0, -1.0, 3.0}, 4.0},
	}

	for _, test_case := range cases {
		mean := Mean(test_case.in)

		if ( mean != test_case.expected ) {
			t.Errorf("Mean(%v) == %v, expected %v", test_case.in, mean, test_case.expected)
		}
	}
}
