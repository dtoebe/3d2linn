package main

import (
	"fmt"
	"os"
)

func getHeader() []string {
	header := []string{"SKU", "Is Variation Group", "Variation SKU", "Variation Group Name", "Title", "Purchase Price",
		"Listing Title (default)", "Listing Description (default)", "Listing Price (default)", "Retail Price", "Brand",
		"image URL 1", "image URL 2", "Image URL 3", "Category", "Level", "Location"}
	return header
}

func parseForExport(data1, data2 [][]string) [][]string {

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

					return
				}
			}
		}
	}

	fmt.Println("Plese make sude your Products and advanced options files exist...")
	fmt.Println("... and your output csv and output errors file does not")
	os.Exit(1)
}
