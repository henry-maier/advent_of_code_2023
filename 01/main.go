package main

import (
	"fmt"
    "bufio"
    "log"
    "os"
	"regexp"
	"strconv"
)
func possible_string_to_int_string(s string) (string, error) {
	switch s {
	case "one":
		return "1", nil
	case "two":
		return "2", nil
	case "three":
		return "3", nil
	case "four":
		return "4", nil
	case "five":
		return "5", nil
	case "six":
		return "6", nil
	case "seven":
		return "7", nil
	case "eight":
		return "8", nil
	case "nine":
		return "9", nil
	default:
		return s, nil
	
	}
}
func calc_sum_part_2 (s string) (int, error) {
	fmt.Println("Getting val for string", s)
	pattern := "([0-9]|one|two|three|four|five|six|seven|eight|nine)"
	reverse_pattern := "([0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)"
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	rr, err := regexp.Compile(reverse_pattern)
	if err != nil {
		log.Fatal(err)
	}
	first := r.FindString(s)
	last := reverse(rr.FindString(reverse(s)))
	fmt.Println("Found first: ", first, " last: ", last)
	fInt, _ := possible_string_to_int_string(first)
	lInt, _ := possible_string_to_int_string(last)
	return strconv.Atoi(fInt + lInt)
}

func calc_sum (s string) (int, error) {
	fmt.Println("Getting val for string", s)
	var first, last string
	r, _ := regexp.Compile("[0-9]")
	for i:=0; i<len(s); i++{
		f := string(s[i])
		l := string(s[len(s) - 1 - i])
		if first == "" && r.MatchString(f) {
			first = f
		}
		if last == "" && r.MatchString(l) {
			last = l
		}
		if first != "" && last != "" {
			break
		}
	}
	fmt.Println("Found first: ", first, " last: ", last)
	return strconv.Atoi(first + last)
}
// thanks https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
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
		new, err := calc_sum_part_2(scanner.Text())
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