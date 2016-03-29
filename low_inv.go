package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseLowInv(data [][]string, invMin int) [][]string {
	rows := [][]string{}
	for i := 0; i < len(data); i++ {
		if i > 0 {
			iInv, err := strconv.Atoi(data[i][13])
			if err != nil {
				fmt.Println(err)
				continue
			}
			if iInv <= invMin {
				rows = append(rows, []string{data[i][1], data[i][2], data[i][13], data[i][48]})
			}

		}
	}
	return rows
}

func getLowInv(inp, out string, lowInv int) {
	if checkFileExist(inp) {
		if !checkFileExist(out) {
			rawData := getRawData(inp)
			lowInvData := parseLowInv(rawData, lowInv)
			header := []string{"ID", "Name", "Inventory", "Is Hidden"}
			writeToCSV(lowInvData, header, out)
		} else {
			fmt.Printf("%s exists. Please choose a file name that does not exist\n", out)
			os.Exit(1)
		}
	} else {
		fmt.Printf("%s does not exist\n", inp)
		os.Exit(1)
	}
}
