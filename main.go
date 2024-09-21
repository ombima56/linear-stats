package main

import (
	"fmt"
	"log"
	"os"

	correlation "learn.zone01kisumu.ke/git/hiombima/linear-stats/pearsonCorrelationCoefficient"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Error Usage: <file path> <argument>")
		return
	}

	filePath := os.Args[1]
	data, err := ReadData(filePath)
	if err != nil {
		log.Printf("Failed to read Data: %s", err)
		return
	}

	slope, intercept := LinearRegression(data)
	correlation := correlation.Correlation(data)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", correlation)
}
