package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseLine(line string) []int {
	res := make([]int, 0)
	curr := ""
	for _, c := range line {
		if c == ' ' {
			num, err := strconv.Atoi(curr)
			if err != nil {
				log.Fatal(err)
			}
			res = append(res, num)
			curr = ""
		} else {
			curr += string(c)
		}
	}
	if curr != "" {
		num, err := strconv.Atoi(curr)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, num)
	}
	return res
}

func solveLine(line string) int {
	curr := parseLine(line)
	lastVals := make([]int, 1)
	lastVals[0] = curr[len(curr)-1]
	diffs := make([]int, len(curr)-1)
	solved := false
	for !solved {
		solved = true
		for idx, val := range curr {
			if idx < len(curr)-1 {
				diffs[idx] = -val
			}
			if idx > 0 {
				diffs[idx-1] += val
				solved = solved && (diffs[idx-1] == 0)
			}
		}
		fmt.Println("Found differences ", diffs)
		lastVals = append(lastVals, diffs[len(diffs)-1])
		curr = diffs
		diffs = make([]int, len(curr)-1)
	}
	res := 0
	fmt.Println("Using lastvals ", lastVals)
	for idx, _ := range lastVals {
		if idx != 0 {
			num := lastVals[len(lastVals)-1-idx]
			fmt.Println("Need ", res+num, "to get diff of ", res, " from ", num)
			res += num
		}
	}
	return res
}

func solvePart1(lines []string) {
	res := 0
	var num int
	for idx, line := range lines {
		fmt.Println("Solving line ", line)
		num = solveLine(line)
		fmt.Println(idx, ": Found next value for ", line, " as ", num)
		res += num
	}
	fmt.Println("Result: ", res)
}

func solveLine2(line string) int {
	curr := parseLine(line)
	firstVals := make([]int, 1)
	firstVals[0] = curr[0]
	diffs := make([]int, len(curr)-1)
	solved := false
	for !solved {
		solved = true
		for idx, val := range curr {
			if idx < len(curr)-1 {
				diffs[idx] = -val
			}
			if idx > 0 {
				diffs[idx-1] += val
				solved = solved && (diffs[idx-1] == 0)
			}
		}
		fmt.Println("Found differences ", diffs)
		firstVals = append(firstVals, diffs[0])
		curr = diffs
		diffs = make([]int, len(curr)-1)
	}
	res := 0
	fmt.Println("Using firstvals ", firstVals)
	for idx, _ := range firstVals {
		if idx != 0 {
			num := firstVals[len(firstVals)-1-idx]
			fmt.Println("Need ", num-res, "to get diff of ", res, " from ", num)
			res = num - res
		}
	}
	return res
}

func solvePart2(lines []string) {
	res := 0
	var num int
	for idx, line := range lines {
		fmt.Println("Solving line ", line)
		num = solveLine2(line)
		fmt.Println(idx, ": Found next value for ", line, " as ", num)
		res += num
	}
	fmt.Println("Result: ", res)
}

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//solvePart1(lines)
	solvePart2(lines)
}
