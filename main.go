package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func checkFileExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	start := time.Now()
	var hidden bool
	var hwi bool
	var inv bool
	var linn bool
	flag.BoolVar(&hidden, "hidden", false, "Sort all hiddeden products and output to new csv")
	flag.BoolVar(&hwi, "hwi", false, "Sort Hidden with inventory and output to new csv")
	flag.BoolVar(&inv, "inv", false, "Sort low inventory based on third param (third param must be a whole number)")
	flag.BoolVar(&linn, "linn", false, "Pull from products and advanced options csv and output csv ready for linnworks and another for any errors")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n %s [options] <product>.csv <output>.csv\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	switch true {
	case linn:
		if flag.NArg() > 3 {
			createLinnCSV(flag.Arg(0), flag.Arg(1), flag.Arg(2), flag.Arg(3))
			elapsed := time.Since(start)
			fmt.Printf("Ran in %s\n", elapsed)

			os.Exit(0)
		}
	case inv:
		if flag.NArg() > 2 {
			lowInv, err := strconv.Atoi(flag.Arg(2))
			if err != nil {
				fmt.Println(err)
				elapsed := time.Since(start)
				fmt.Printf("Ran in %s\n", elapsed)

				os.Exit(1)
			}
			getLowInv(flag.Arg(0), flag.Arg(1), lowInv)
		} else {
			flag.Usage()
			os.Exit(1)
		}
	case hwi:
		if flag.NArg() > 1 {
			getHidden(flag.Arg(0), flag.Arg(1), "hwi")
		} else {
			flag.Usage()
			os.Exit(1)
		}
		os.Exit(0)
	case hidden:
		if flag.NArg() > 1 {
			getHidden(flag.Arg(0), flag.Arg(1), "hidden")
		} else {
			flag.Usage()
			os.Exit(1)
		}
		os.Exit(0)
	}
	elapsed := time.Since(start)
	fmt.Printf("Ran in %s\n", elapsed)
}
