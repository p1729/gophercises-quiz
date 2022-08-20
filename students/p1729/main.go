package main

import (
	"fmt"
	"os"
)

func main() {
	var flag string
	if len(os.Args[1:]) == 0 {
		flag = ""
	} else {
		flag = os.Args[1]
	}

	switch flag {
	case "-h":
		fmt.Println("Usage of quiz")
		fmt.Println("\t-csv string")
		fmt.Println("\t\ta csv file in the format of 'question,answer' (default \"problems.csv\")")
	case "-csv":
		if len(os.Args[2:]) == 0 {
			fmt.Println("Please provide the filename")
		} else {
			fmt.Println("You provided filename", os.Args[2])
		}
	default:
		fmt.Println("Default filename is problems.csv")
	}
}
