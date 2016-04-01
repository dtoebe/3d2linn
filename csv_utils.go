package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func getRawData(file string) [][]string {
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvRaw, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("[ERR] Error reading %s: %v\n", file, err)
		os.Exit(1)
	}

	return csvRaw

}

func writeToCSV(data [][]string, header []string, path string) {
	csvFile, err := os.Create(path)
	if err != nil {
		fmt.Printf("[ERR] Cannot create output csv file: %v\n", err)
		os.Exit(1)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	err = writer.Write(header)
	if err != nil {
		fmt.Printf("[ERR] Error writing header: %v\n", err)
		os.Exit(1)
	}
	for i := 0; i < len(data); i++ {
		err := writer.Write(data[i])
		if err != nil {
			fmt.Println("[ERR] Error writing line %d: %v\n", i+1, err)
			return
		}
	}
	writer.Flush()
}
