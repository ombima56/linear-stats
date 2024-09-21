package main

import (
	"fmt"
	"log"
	"os"

	linearregressionline "github.com/ombima56/linear-stats/linear-stats/linearRegressionLine"
	correlation "github.com/ombima56/linear-stats/linear-stats/pearsonCorrelationCoefficient"
	"github.com/ombima56/linear-stats/linear-stats/readfile"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Error Usage: <file path> <argument>")
		return
	}

	filePath := os.Args[1]
	data, err := readfile.ReadData(filePath)
	if err != nil {
		log.Printf("Failed to read Data: %s", err)
		return
	}

	slope, intercept := linearregressionline.LinearRegression(data)
	correlation := correlation.Correlation(data)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", correlation)
}
