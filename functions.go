package stats

import "math"
import "errors"

type Point struct {
	X float64
	Y float64
}

// Calculate the mean of a set of float64s
func Mean(numbers []float64) float64 {
	sum := 0.0
	num := len(numbers)

	if num == 0 {
		return sum
	}

	for _, number := range numbers {
		sum += number
	}

	return (float64(sum) / float64(num))
}

// Calculate the variance of a set of numbers
func Variance(numbers []float64) (float64, error) {
	num := len(numbers)

	if ( num < 2 ) {
		return 0.0, errors.New("Variance takes at least two numbers")
	}

	mean := Mean(numbers)
	differences := make([]float64, len(numbers))

	// Calculate the squared differences of the numbers.
	for i, number := range numbers {
		differences[i] = math.Pow(mean - number, 2)
	}

	// Return the mean of those squared numbers.
	return Mean(differences), nil
}

func StandardDeviation(numbers []float64) (float64, error) {
	variance, err := Variance(numbers)

	if err != nil {
		return 0.0, err
	}

	return math.Sqrt(variance), nil
}

func getXData(points []Point) ([]float64) {
	point_length := len(points)
	x_array := make([]float64, point_length)

	for i,point := range points {
		x_array[i] = point.X
	}

	return x_array
}

func getYData(points []Point) ([]float64) {
	point_length := len(points)
	y_array := make([]float64, point_length)

	for i,point := range points {
		y_array[i] = point.Y
	}

	return y_array
}

func Covariance(points []Point) float64 {
	num := len(points)
	sum := 0.0

	if num == 0 {
		return sum
	}

	x_mean := Mean(getXData(points))
	y_mean := Mean(getYData(points))

	for _, point := range points {
		sum += (point.X - x_mean) * (point.Y - y_mean)
	}

	return sum / float64(num - 1)
}
