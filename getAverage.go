package getAverage

// GetAverage why is branch switch between master and others so vague..
func GetAverage(numbers []float64) float64 {
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}
