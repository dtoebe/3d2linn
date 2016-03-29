package main

import (
	"fmt"
	"os"
	"strconv"
)

func formatHiddenAll(data [][]string) [][]string {
	rows := [][]string{}
	for i := 0; i < len(data); i++ {
		rows = append(rows, []string{data[i][1], data[i][2], data[i][13], data[i][7], data[i][8], data[i][11], data[i][12], data[i][19]})
	}
	return rows
}

func formatHWI(data [][]string) [][]string {
	rows := [][]string{}
	for i := 0; i < len(data); i++ {
		inv, err := strconv.Atoi(data[i][13])
		if err != nil {
			fmt.Println(err)
			continue
		}
		if inv > 0 {
			rows = append(rows, []string{data[i][1], data[i][2], data[i][13]})
		}
	}
	return rows
}

func parseHidden(data [][]string) [][]string {
	hidden := [][]string{}
	for i := 0; i < len(data); i++ {
		if data[i][48] == "1" {
			hidden = append(hidden, data[i])
		}
	}

	return hidden
}

func getHidden(inp, out string, t string) {
	if checkFileExist(inp) {
		var header []string
		var hiddenRows [][]string
		if !checkFileExist(out) {
			rawData := getRawData(inp)
			hiddenData := parseHidden(rawData)
			switch t {
			case "hidden":
				header = []string{"ID", "Name", "Inventory", "Cost", "Price", "Sale Price", "Is On Sale", "Date Created"}
				hiddenRows = formatHiddenAll(hiddenData)
			case "hwi":
				header = []string{"ID", "Name", "Inventory"}
				hiddenRows = formatHWI(hiddenData)
			default:
				return
			}
			writeToCSV(hiddenRows, header, out)
		} else {
			fmt.Printf("%s exists. Please choose a filename that does not exist\n", out)
			os.Exit(1)
		}
	} else {
		fmt.Printf("%s does not exist\n", inp)
		os.Exit(1)
	}
}
