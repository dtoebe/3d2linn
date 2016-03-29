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

	csvRaw, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return csvRaw

}

func writeToCSV(data [][]string, header []string, path string) {
	csvFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	err = writer.Write(header)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < len(data); i++ {
		err := writer.Write(data[i])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	writer.Flush()
}
