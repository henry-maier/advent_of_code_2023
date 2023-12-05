package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type mapping struct {
	mapLines []mapline
}

type mapline struct {
	sourceStart int
	destStart   int
	r           int
}

func (m *mapping) addLine(ml mapline) {
	m.mapLines = append(m.mapLines, ml)
}

func (m *mapping) toDest(i int) int {
	for _, ml := range m.mapLines {
		if ml.contains(i) {
			return ml.toDest(i)
		}
	}
	return i
}

func (ml *mapline) contains(i int) bool {
	return i >= ml.sourceStart && i-ml.sourceStart < ml.r
}

func (ml *mapline) toDest(i int) int {
	return ml.destStart + (i - ml.sourceStart)
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {
	fileToUse := "input1.txt"
	file, err := os.Open(fileToUse)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var seeds []int
	// initialize seeds
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		for _, match := range r.FindAllString(line, -1) {
			seeds = append(seeds, toInt(match))
		}
	}
	fmt.Println("Found initial seed ranges: ", seeds)
	mapList := make([]mapping, 0)
	currMapping := mapping{mapLines: make([]mapline, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// process mapping
			// for idx, val := range seeds {
			//	seeds[idx] = currMapping.toDest(val)
			// }
			mapList = append(mapList, currMapping)
			fmt.Println("finished processing mapping...")
			fmt.Println(currMapping)
			currMapping = mapping{mapLines: make([]mapline, 0)}
		} else {
			matches := r.FindAllString(line, -1)
			if len(matches) != 3 {
				fmt.Println("Processing mapping: ", line)
			} else {
				ml := mapline{
					sourceStart: toInt(matches[1]),
					destStart:   toInt(matches[0]),
					r:           toInt(matches[2])}
				currMapping.addLine(ml)
			}
		}
	}
	// process mapping one more time for off by one
	mapList = append(mapList, currMapping)
	fmt.Println("finished processing final mapping...")
	fmt.Println(currMapping)
	res := int(math.Inf(1))
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := start + seeds[i+1]
		fmt.Println("Trying seeds in range (", start, ",", end, ")")
		for curr := start; curr < end; curr++ {
			mappedVal := curr
			for _, m := range mapList {
				//fmt.Print(mappedVal, "->")
				mappedVal = m.toDest(mappedVal)
				//fmt.Print(mappedVal)
			}
			//fmt.Println()
			if mappedVal < res {
				res = mappedVal
				fmt.Println("new min: ", res, " from start ", curr)
			}
		}
	}
	fmt.Println("Lowest location: ", res)
}
