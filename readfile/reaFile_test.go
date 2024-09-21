package readfile

import (
	"os"
	"testing"
)

func TestReadData_Success(t *testing.T) {
	file, err := os.CreateTemp("", "testing*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(file.Name())

	content := "189.2\n134.0\n121.5\n114.4\n145.3\n110.2\n"
	if _, err := file.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}
	file.Close()

	data, err := ReadData(file.Name())
	if err != nil {
		t.Fatalf("ReadData failed: %s", err)
	}

	expected := []float64{189.2, 134.0, 121.5, 114.4, 145.3, 110.2}

	if len(data) != len(expected) {
		t.Errorf("Expected %d elements, but got %d", len(expected), len(data))
	}

	for i, v := range expected {
		if data[i] != v {
			t.Errorf("Expected value at index %d to be %.2f, but got %.2f", i, v, data[i])
		}
	}
}


func TestReadData_FileNotFound(t *testing.T) {
	_, err := ReadData("non_existing_file.txt")
	if err == nil {
		t.Error("Expected an error for non-existing file, but got none")
	}
}

func TestReadData_InvalidData(t *testing.T) {
	file, err := os.CreateTemp("", "invaliddata*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer os.Remove(file.Name()) 

	content := "invalid\n123.0\n"
	if _, err := file.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}

	file.Close()

	_, err = ReadData(file.Name())
	if err == nil {
		t.Error("Expected an error for invalid data, but got none")
	}
}


