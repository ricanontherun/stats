package stats

import "math"

func Test() string {
	return "Hello from the stats library!"
}

// Calculate the mean of a set of float64s
func Mean(numbers []float64) float64 {
	var sum = 0.0

	for _, number := range numbers {
		sum += number
	}

	return (float64(sum) / float64(len(numbers)))
}

// Calculate the variance of a set of numbers
func Variance(numbers []float64) float64 {
	var mean = Mean(numbers)
	var differences = make([]float64, len(numbers))

	// Calculate the squared differences of the numbers.
	for i, number := range numbers {
		differences[i] = math.Pow(mean - number, 2)
	}

	// Return the mean of those squared numbers.
	return Mean(differences)
}

func StandardDeviation(numbers []float64) float64 {
	return math.Sqrt(Variance(numbers))
}
