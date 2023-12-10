package main

import (
	"fmt"
	"os"
	"strings"
)

func preprocess_text(text string) [][]int {
	var numbers []string
	numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var nrs [][]int
	for _, line := range strings.Split(text, "\n") {
		fmt.Printf("%s\n", line)
		var line_nrs []int = []int{}
		for index, char := range line {
			for n_in, number := range numbers {
				if strings.HasPrefix(line[index:], number) {
					line_nrs = append(line_nrs, n_in)
				}
			}
			if char >= '0' && char <= '9' {
				line_nrs = append(line_nrs, int(char-'0'))
			}
		}
		nrs = append(nrs, line_nrs)
		fmt.Printf("line_nrs: %v\n", line_nrs)
	}
	return nrs
}

func main() {
	var fc []byte
	fc, err := os.ReadFile("inputs/d1_in.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var input = string(fc)
	// var input = `two1nine
	// eightwothree
	// abcone2threexyz
	// xtwone3four
	// 4nineeightseven2
	// zoneight234
	// 7pqrstsixteen`
	//iterate through each line
	var res = preprocess_text(input)

	var total int = 0
	for _, line := range res {
		total += line[0]*10 + line[len(line)-1]
		fmt.Printf("first %d last %d\n", line[0], line[len(line)-1])
	}
	fmt.Printf("total: %d\n", total)
}
