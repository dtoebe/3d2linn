package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func checkFileExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	var hidden bool
	var hwi bool
	var inv bool
	flag.BoolVar(&hidden, "hidden", false, "Sort all hiddeden products and output to new csv")
	flag.BoolVar(&hwi, "hwi", false, "Sort Hidden with inventory and output to new csv")
	flag.BoolVar(&inv, "inv", false, "Sort low inventory based on third param (third param must be a whole number)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n %s [options] <product>.csv <output>.csv\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	switch true {
	case inv:
		if flag.NArg() > 2 {
			lowInv, err := strconv.Atoi(flag.Arg(2))
			if err != nil {
				fmt.Println(err)
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
}
