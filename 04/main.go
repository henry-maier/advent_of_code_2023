package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

func getScoreForLinePartOne(s string) (int, error) {
	foundStart := false
	foundWinningNumbers := false
	currNumber := ""
	m := make(map[string]bool)
	sum := 0
	for _, char := range s {
		if !foundStart && char == ':' {
			foundStart = true
		} else if foundStart && char == '|' {
			foundWinningNumbers = true
		} else if foundStart && char == ' ' && currNumber != "" {
			// end of a number
			// if past the winning numbers, check and add
			if foundWinningNumbers && m[currNumber] {
				new, _ := strconv.Atoi(currNumber)
				fmt.Println("Found winning number, ", new)
				sum += 1
				// otherwise we are still processing winning numbers, add to the set
			} else {
				m[currNumber] = true
			}
			// reset the curr number
			currNumber = ""
		} else if foundStart && unicode.IsDigit(char) {
			currNumber += string(char)
		}
	}
	// off by one
	if foundWinningNumbers && m[currNumber] {
		new, _ := strconv.Atoi(currNumber)
		fmt.Println("Found winning number, ", new)
		sum += 1
		currNumber = ""
	}
	if sum == 0 {
		return 0, nil
	}
	return int(math.Pow(2, float64(sum-1))), nil
}

func processLine(s string, idx int, copies []int) (int, error) {
	// when processing a line, there is at least one copy of it
	copies[idx] += 1
	// find the winning numbers (copying from above I am lazy today)
	foundStart := false
	foundWinningNumbers := false
	currNumber := ""
	m := make(map[string]bool)
	sum := 0
	for _, char := range s {
		if !foundStart && char == ':' {
			foundStart = true
		} else if foundStart && char == '|' {
			foundWinningNumbers = true
		} else if foundStart && char == ' ' && currNumber != "" {
			// end of a number
			// if past the winning numbers, check and add
			if foundWinningNumbers && m[currNumber] {
				new, _ := strconv.Atoi(currNumber)
				fmt.Println("Found winning number, ", new)
				sum += 1
				// otherwise we are still processing winning numbers, add to the set
			} else {
				m[currNumber] = true
			}
			// reset the curr number
			currNumber = ""
		} else if foundStart && unicode.IsDigit(char) {
			currNumber += string(char)
		}
	}
	// off by one
	if foundWinningNumbers && m[currNumber] {
		new, _ := strconv.Atoi(currNumber)
		fmt.Println("Found winning number, ", new)
		sum += 1
		currNumber = ""
	}
	for i := 0; i < sum; i++ {
		copies[idx+i+1] += copies[idx]
	}
	return sum, nil
}

func main() {
	// thanks https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	var copies = make([]int, 1000)
	for scanner.Scan() {
		new, err := processLine(scanner.Text(), i, copies)
		if err != nil {
			log.Fatal(err)
		}
		i += 1
		fmt.Println("Line ", i, " had ", new, " winning nums and ", copies[i], " copies.")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sum := 0
	for idx := 0; idx < i+1; idx++ {
		fmt.Println(copies[idx], " copies total for card ", idx)
		sum += copies[idx]
	}
	fmt.Println("Result:", sum)
}
