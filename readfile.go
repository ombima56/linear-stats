package main

import (
	"bufio"
	"os"
	"strconv"
)

func ReadData(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, val)
	}
	return data, scanner.Err()
}