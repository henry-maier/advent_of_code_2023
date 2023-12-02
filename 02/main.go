package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var COLOR_MAP = map[string]int{"red": 12, "blue": 14, "green": 13}

func get_id_if_valid(s string) (int, error) {
	fmt.Println("Handling game: ", s)
	for color, m := range COLOR_MAP {
		if !match_nums(s, color, m) {
			return 0, nil
		}
	}
	res, err := get_id(s)
	fmt.Println("Parsed ", res, " from id for valid game")
	return res, err
}

func get_id(s string) (int, error) {
	r, err := regexp.Compile("Game [0-9]+")
	match := r.FindString(s)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Atoi(match[5:])

}

func match_nums(s string, c string, t int) bool {
	r, err := regexp.Compile("[0-9]+ " + c)
	if err != nil {
		log.Fatal(err)
	}
	matches := r.FindAllString(s, -1)
	for _, match := range matches {
		num, err := strconv.Atoi(match[:len(match)-(len(c)+1)])
		if err != nil {
			log.Fatal(err)
		}
		if num > t {
			fmt.Println("Invalid game, found ", num, " but expected only ", t, " for ", c)
			return false
		}
	}
	return true
}

func get_power_of_game(s string) (int, error) {
	fmt.Println("Handling game: ", s)
	res := 1
	for color := range COLOR_MAP {
		lowest, err := get_max(s, color)
		if err != nil {
			log.Fatal(err)
		}
		res *= lowest
	}
	fmt.Println("Found ", res, " as power from game")
	return res, nil
}

func get_max(s string, c string) (int, error) {
	r, err := regexp.Compile("[0-9]+ " + c)
	if err != nil {
		log.Fatal(err)
	}
	matches := r.FindAllString(s, -1)
	m := 0
	for _, match := range matches {
		num, err := strconv.Atoi(match[:len(match)-(len(c)+1)])
		if err != nil {
			log.Fatal(err)
		}
		if num > m {
			m = num
		}
	}
	fmt.Println("Found max ", m, " for color ", c)
	return m, nil
}

func main() {
	// thanks https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		new, err := get_power_of_game(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += new
		fmt.Println("Added ", new, " for a new total of ", sum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:", sum)
}
