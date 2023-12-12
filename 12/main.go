package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
)

var cache = make(map[string]int)

func parseLine(line string) ([]rune, []int) {
	springs := make([]rune, 0)
	counts := make([]int, 0)
	springsFound := false
	numString := ""
	for _, c := range line {
		if c == ' ' {
			springsFound = true
		} else if springsFound && c != ',' {
			numString += string(c)
		} else if springsFound && c == ',' {
			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}
			counts = append(counts, num)
			numString = ""
		} else {
			// springs
			springs = append(springs, c)
		}
	}
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatal(err)
	}
	counts = append(counts, num)
	return springs, counts
}

func getHash(s []rune, c []int, b int) string {
	h := sha256.New()
	fmt.Fprint(h, s, c, b)
	// h.Write([]byte(fmt.Sprint(s)))
	// cString := ""
	// for _, r := range c {
	// 	cString += fmt.Sprint(r)
	// }
	// h.Write([]byte(cString))
	// h.Write([]byte(fmt.Sprint(b)))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func solveLineRecursive(springs []rune, counts []int, brokenCount int) int {
	if len(springs) == 0 {
		if len(counts) == 0 && brokenCount == 0 || (len(counts) == 1 && brokenCount == counts[0]) {
			return 1
		} else {
			return 0
		}
	}
	h := getHash(springs, counts, brokenCount)
	if v, ok := cache[h]; ok {
		fmt.Println("Cache hit! for ", springs, counts, brokenCount, ": ", v)
		return v
	}
	lastSpring := springs[len(springs)-1]
	lastCount := 0
	if len(counts) > 0 {
		lastCount = counts[len(counts)-1]
	}
	if brokenCount > lastCount {
		fmt.Println("Doesn't work!", brokenCount, lastCount, counts)
		cache[h] = 0
		return 0
	}
	if lastSpring == '.' {
		if brokenCount == 0 {
			fmt.Println("Working spring")
			res := solveLineRecursive(springs[:len(springs)-1], counts, 0)
			cache[h] = res
			return res
		} else if brokenCount < lastCount {
			fmt.Println("Doesn't work!", brokenCount, lastCount, counts)
			cache[h] = 0
			return 0
		} else {
			fmt.Println("working spring, breaking a chain")
			res := solveLineRecursive(springs[:len(springs)-1], counts[:len(counts)-1], 0)
			cache[h] = res
			return res
		}
	} else if lastSpring == '#' {
		fmt.Println("broken spring, increasing broken count")
		res := solveLineRecursive(springs[:len(springs)-1], counts, brokenCount+1)
		cache[h] = res
		return res
	} else {
		brokenSpring := make([]rune, len(springs))
		copy(brokenSpring, springs)
		brokenSpring[len(brokenSpring)-1] = '#'
		workingSpring := make([]rune, len(springs))
		copy(workingSpring, springs)
		workingSpring[len(brokenSpring)-1] = '.'
		broken := solveLineRecursive(brokenSpring, counts, brokenCount)
		fmt.Println("Found ", broken, " options using broken spring from ", springs, " ", counts, " ", brokenCount)
		working := solveLineRecursive(workingSpring, counts, brokenCount)
		fmt.Println("Found ", working, " options using working spring from ", springs, " ", counts, " ", brokenCount)
		res := broken + working
		cache[h] = res
		return res
	}

}
func solveLine(line string) int {
	springs, counts := parseLine(line)
	return solveLineRecursive(springs, counts, 0)
}

func solveLine2(line string) int {
	springs, counts := parseLine(line)
	newSprings := make([]rune, len(springs)*5+4)
	newCounts := make([]int, len(counts)*5)
	for i := 0; i < 5; i++ {
		for idx, val := range springs {
			newSprings[i*len(springs)+idx+i] = val
		}
		if i < 4 {
			newSprings[i*len(springs)+len(springs)+i] = '?'
		}
		for idx, val := range counts {
			newCounts[i*len(counts)+idx] = val
		}
	}
	fmt.Println("Actually solving line ", newSprings, newCounts)
	return solveLineRecursive(newSprings, newCounts, 0)
}

func solvePart1(lines []string) {
	var res uint64
	var num int
	for idx, line := range lines {
		fmt.Println("Solving line ", line)
		num = solveLine(line)
		fmt.Println(idx, ": Found next value for ", line, " as ", num)
		res += uint64(num)
	}
	fmt.Println("Result: ", res)
}

func solvePart2(lines []string) {
	var res uint64
	var num int
	for idx, line := range lines {
		fmt.Println("Solving line ", line)
		num = solveLine2(line)
		fmt.Println(idx, ": Found next value for ", line, " as ", num)
		res += uint64(num)
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
