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

type result struct {
	correct int
	total   int
}

func (q question) String() string {
	return q.problem
}

func (r result) String() string {
	return fmt.Sprintf("%d/%d", r.correct, r.total)
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
		result := playQuiz(questions)
		fmt.Printf("You scored %d out of %d ", result.correct, result.total)
	}
}

func playQuiz(questions []question) result {
	var res result
	reader := bufio.NewReader(os.Stdin)
	for index, question := range questions {
		fmt.Printf("Problem #%d: %s = ", index+1, question)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error while reading input")
		}
		if question.answer == strings.TrimRight(input, "\n") {
			res.correct++
		}
		res.total++
	}
	return res
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
