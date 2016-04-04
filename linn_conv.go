package main

import (
	"fmt"
	"os"
	"strings"
)

func getHeader() []string {
	header := []string{"SKU", "Is Variation Group", "Variation SKU", "Variation Group Name", "Title", "Purchase Price",
		"Listing Title (default)", "Listing Description (default)", "Listing Price (default)", "Retail Price", "Brand",
		"Range", "Variation Title", "image URL 1", "image URL 2", "Image URL 3", "Category", "Level", "Location"}
	return header
}

func looseSelVarName(s string) string {
	if strings.Contains(strings.ToLower(s), "by vol") {
		return "Strength"
	} else if strings.Contains(strings.ToLower(s), "byvol") {
		return "Strength"
	} else if strings.Contains(strings.ToLower(s), "ohm") {
		return "Resistance"
	} else if strings.Contains(strings.ToLower(s), "mah") {
		return "Size"
	}
	return "Color"
}

func checkVarCost(s1, s2 string) string {
	if s2 == "0" || s2 == "0.00" {
		return s1
	}
	return s2
}

func isOnSale(price, salePrice, saleBool string) string {
	if saleBool == "1" {
		return salePrice
	}
	return price
}

func cleanCategory(s string) string {
	at := []string{}
	slash := strings.Split(s, "/")
	if len(slash) >= 2 {
		at = strings.Split(slash[1], "@")
	} else {
		at = slash
	}
	return at[0]
}

func cleanDesc(sku, name, s, path string) string {
	var newDesc string
	if strings.Contains(s, "\ufffd") {
		writeData := []string{}
		writeData = append(writeData, fmt.Sprintf("%s - %s\n", sku, name))
		for i, c := range s {
			if string(c) == "\ufffd" {
				var begin int
				var end int
				if i < 40 {
					begin = 0
				} else {
					begin = i - 40
				}
				if i > len(s)-40 {
					end = len(s)
				} else {
					end = i + 40
				}
				writeData = append(writeData, fmt.Sprintf("\tBefore: %s\n", s[begin:i]))
				writeData = append(writeData, fmt.Sprintf("\tAfter:  %s\n\n", s[i:end]))
			}
		}
		writeData = append(writeData, "\n\n")
		writeToTxt(writeData, path)

		newDesc = strings.Replace(s, "\ufffd", "", -1)
		newDesc = strings.TrimSpace(newDesc)
	} else {
		newDesc = s
	}

	return newDesc

}

func parseForExport(data1, data2 [][]string, outPath string) [][]string {
	exportdata := [][]string{}
	for i := 0; i < len(data1); i++ {
		if data1[i][0] != "catalogid" {
			rows := [][]string{}
			isVariance := "No"   //MATCH Is Variation Group
			variationGroup := "" //MATCH Variation Group Name
			variationSKU := ""   //MATCH Variation SKU
			variationRange := "" //MATCH Variation Name
			variationTitle := "" //MATCH Variation Title
			price := isOnSale(data1[i][8], data1[i][11], data1[i][12])
			for j := 0; j < len(data2); j++ {
				if data1[i][0] == data2[j][0] {
					isVariance = "Yes"
					variationGroup = data1[i][2]
					variationSKU = data1[i][1]
					variationRange = looseSelVarName(data2[j][4])
					variationTitle = data2[j][4]
					varTitle := data2[j][1] + " - " + data2[j][4]
					varCost := checkVarCost(data1[i][7], data2[j][5])
					rows = append(rows, []string{data2[j][3], "No", data1[i][1], data1[i][2], varTitle, varCost,
						varTitle, "", price, price, data1[i][5], variationRange, variationTitle, "", "", "", cleanCategory(data1[i][3]),
						"4", "Default"})
				}
			}
			//TODO Fix Image URL
			desc := cleanDesc(data1[i][1], data1[i][2], data1[i][21], outPath)
			exportdata = append(exportdata, []string{data1[i][1], isVariance, variationSKU, variationGroup, data1[i][2], data1[i][7],
				data1[i][2], desc, price, price, data1[i][5], "", "", data1[i][25], data1[i][26], data1[i][27],
				cleanCategory(data1[i][3]), "4", "Default"})

			for _, row := range rows {
				exportdata = append(exportdata, row)
			}
		}
	}
	return exportdata
}

func createLinnCSV(inp1, inp2, out, errOut string) {
	if checkFileExist(inp1) {
		if checkFileExist(inp2) {
			if !checkFileExist(out) {
				if !checkFileExist(errOut) {
					rawData1 := getRawData(inp1)
					rawData2 := getRawData(inp2)
					header := getHeader()
					wantedData := rawData1
					if true { //TODO make flag for sorting hidden or not
						wantedData = parseHidden(rawData1, false)
					}
					allData := parseForExport(wantedData, rawData2, errOut)
					writeToCSV(allData, header, out)

					return
				}
			}
		}
	}

	fmt.Println("Plese make sude your Products and advanced options files exist...")
	fmt.Println("... and your output csv and output errors file does not")
	os.Exit(1)
}
