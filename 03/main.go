package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type schematic struct {
	m, n     int
	contents [][]rune
}

func newSchematic(fileName string) *schematic {
	// thanks https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	fmt.Println("Opening file to create schmatic ", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var contents [][]rune
	for scanner.Scan() {
		contents = append(contents, []rune(scanner.Text()))
	}
	fmt.Println("Read in file.")
	ret := schematic{m: len(contents), n: len(contents[0]), contents: contents}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Constructed schematic with ", ret.m, " strings of len ", ret.n)
	return &ret
}

func (s schematic) isNumber(i int, j int) bool {
	return unicode.IsDigit(s.get(i, j))
}

func (s schematic) isGear(i int, j int) bool {
	return s.get(i, j) == '*'
}

func (s schematic) get(i int, j int) rune {
	if i < 0 || j < 0 || i >= s.m || j >= s.n {
		return '.'
	}
	return s.contents[i][j]
}

func (s schematic) isSymbolAround(i int, start int, end int) bool {
	// try one left and one right
	if s.isSymbol(i, start-1) || s.isSymbol(i, end) {
		return true
	}
	// try row above and below
	for j := start - 1; j <= end; j++ {
		if s.isSymbol(i-1, j) || s.isSymbol(i+1, j) {
			return true
		}
	}
	return false
}

func (s schematic) convertGearLocation(i int, j int) int {
	return i*s.m + j
}

func (s schematic) getGearRowFromLocation(l int) int {
	return l / s.m
}

func (s schematic) getGearColFromLocation(l int) int {
	return l % s.m
}

func (s schematic) getGearsAround(i int, start int, end int) []int {
	var res []int
	for r := i - 1; r <= i+1; r++ {
		for c := start - 1; c <= end; c++ {
			if s.isGear(r, c) {
				res = append(res, s.convertGearLocation(r, c))
			}
		}
	}
	return res
}

func (s schematic) isSymbol(i int, j int) bool {
	return s.get(i, j) != '.'
}

func (s schematic) getPartNumsSum() (int, error) {
	sum := 0
	for i := 0; i < s.m; i++ {
		for j := 0; j < s.n; j++ {
			if s.isNumber(i, j) {
				start := j
				end := start + 1
				for ; s.isNumber(i, end); end++ {
				}
				// number is s.contents[i][start:end] end is exclusive
				if s.isSymbolAround(i, start, end) {
					num, err := strconv.Atoi(string(s.contents[i][start:end]))
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Found part num ", num)
					sum += num
				}
				j = end - 1
			}
		}
	}
	return sum, nil
}

func (s schematic) getGearRatioSum() (int, error) {
	m := make(map[int][]int)
	for i := 0; i < s.m; i++ {
		for j := 0; j < s.n; j++ {
			if s.isNumber(i, j) {
				start := j
				end := start + 1
				for ; s.isNumber(i, end); end++ {
				}
				for _, gear := range s.getGearsAround(i, start, end) {
					if m[gear] == nil {
						var tmp []int
						m[gear] = tmp
					}
					num, err := strconv.Atoi(string(s.contents[i][start:end]))
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Found part num next to gear ", num)
					m[gear] = append(m[gear], num)
				}
				// number is s.contents[i][start:end] end is exclusive
				j = end - 1
			}
		}
	}
	sum := 0
	for k, v := range m {
		if len(v) == 2 {
			fmt.Println("Gear at ", s.getGearRowFromLocation(k), ",", s.getGearColFromLocation(k), " has exactly 2 part nums, ", v)
			sum += v[0] * v[1]
		}
	}
	return sum, nil
}

func main() {
	fileToUse := "input1.txt"
	s := newSchematic(fileToUse)
	// fmt.Println(s.getPartNumsSum())
	fmt.Println(s.getGearRatioSum())
}
