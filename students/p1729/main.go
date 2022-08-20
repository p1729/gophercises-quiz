package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type question struct {
	problem string
	answer  string
}

func (q question) String() string {
	return q.problem
}

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
		questions := parseQuizFile("problems.csv")
		for index, question := range questions {
			fmt.Printf("Problem #%d: %s =\n", index+1, question)
		}
	}
}

func parseQuizFile(filename string) []question {
	questions := make([]question, 0, 100)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("not able to read file ", filename)
		return []question{}
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		values := strings.Split(text, ",")
		questions = append(questions, question{values[0], values[1]})
	}
	return questions
}
