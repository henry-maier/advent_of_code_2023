package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func findRange(time int, dist int) int {
	timeFloat := float64(time)
	distFloat := float64(dist)
	// x(time - x) = dist
	// (time)x  - x^2 = dist
	// x^2 - (time)x + dist = 0
	// ((time) +/- sqrt(time^2 - 4(dist))) / 2
	inner := math.Sqrt(math.Pow(timeFloat, 2) - 4*distFloat)
	lower := (timeFloat - inner) / 2
	upper := (timeFloat + inner) / 2
	fmt.Println("upper ", upper, "lower ", lower)
	// get lower and upper bounds (neither are inclusive
	// subtract then subtract one)
	return int(math.Ceil(upper)-math.Floor(lower)) - 1
}

func processLine(s string) []int {
	sl := strings.Fields(s)
	sl = sl[1:]
	ret := make([]int, len(sl))
	for idx, s := range sl {
		ret[idx], _ = strconv.Atoi(s)
	}
	return ret
}

func processLinePart2(s string) int {
	sl := strings.Fields(s)
	sl = sl[1:]
	stringRet := ""
	for _, s := range sl {
		stringRet += s
	}
	ret, _ := strconv.Atoi(stringRet)
	return ret

}

func solvePart1(line1 string, line2 string) {
	times := processLine(line1)
	distances := processLine(line2)
	var time, dist, r int
	res := 1
	for idx := range times {
		time = times[idx]
		dist = distances[idx]
		fmt.Println("Finding range for time ", time, " distance ", dist)
		r = findRange(time, dist)
		fmt.Println("Found range ", r)
		res *= r
	}
	fmt.Println("Final answer for part 1: ", res)
}

func solvePart2(line1 string, line2 string) {
	time := processLinePart2(line1)
	distance := processLinePart2(line2)
	fmt.Println("Finding range for time ", time, " distance ", distance)
	r := findRange(time, distance)
	fmt.Println("Final answer for part 2: ", r)
}

func main() {
	// thanks https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()
	solvePart1(line1, line2)
	solvePart2(line1, line2)
}
