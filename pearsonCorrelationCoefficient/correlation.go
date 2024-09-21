package correlation

import "math"


func Correlation(data []float64) float64 {
	num := float64(len(data))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i, y := range data {
		x := float64(i)
        sumX += x
        sumY += y
        sumXY += x * y
        sumX2 += x * x
        sumY2 += y * y
	}

	numerator := num*sumXY - sumX*sumY
    denominator := math.Sqrt((num*sumX2 - sumX*sumX) * (num*sumY2 - sumY*sumY))

	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}