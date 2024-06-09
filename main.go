package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"Documents/functions"
)

func main() {
	// Check if there is the required number of command-line arguments parsed
	if len(os.Args) != 2 {
		fmt.Println("Required: go run . data.txt")
		return
	}

	// Check if the argument ends with ".txt"
	if !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("Only txt file required")
		os.Exit(1)
	}

	// Read data from the file
	dataBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(dataBytes) == 0 {
		fmt.Println("Empty file")
		return
	}

	// Convert data to float64 slice
	dataStr := string(dataBytes)
	dataLines := strings.Split(dataStr, "\n")
	var data []float64
	count := 0
	for _, line := range dataLines {
		count++
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error parsing data in line:", count)
			return
		}
		data = append(data, value)
	}

	// Calculate statistics
	Average := functions.Average(data)
	Median := functions.Median(data)
	Variance := functions.Variance(data)
	StandardDev := functions.StandardDev(data)

	// Print the results
	fmt.Println("Average:", int(math.Round(Average)))
	fmt.Println("Median:", int(math.Round(Median)))
	fmt.Println("Variance:", int(math.Round(Variance)))
	fmt.Println("Standard Deviation:", int(math.Round(StandardDev)))
}
