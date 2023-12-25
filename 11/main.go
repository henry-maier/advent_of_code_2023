package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type planet struct {
	x         int
	y         int
	originalX int
	originalY int
}

func newPlanet(x int, y int) planet {
	return planet{x: x, y: y, originalX: x, originalY: y}
}

func (p *planet) expandX(by int) {
	p.x += by
}

func (p *planet) expandY(by int) {
	p.y += by
}

func (p *planet) dist(p1 *planet) int {
	return dist(p, p1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dist(p1 *planet, p2 *planet) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func solvePart(lines []string, expandBy int) {
	fmt.Println("Expanding universe by ", expandBy)
	planets := make([]planet, 0)
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				fmt.Println("Found planet at (", x, ",", y, ").")
				planets = append(planets, newPlanet(x, y))
			}
		}
	}
	for y, line := range lines {
		if !strings.Contains(line, "#") {
			fmt.Println("Expanding row ", y)
			for idx := range planets {
				if planets[idx].originalY > y {
					planets[idx].expandY(expandBy)
				}
			}
			fmt.Println("expanded: ", planets)
		}
	}
	for x := 0; x < utf8.RuneCountInString(lines[0]); x++ {
		foundPlanet := false
		for _, line := range lines {
			if line[x] == '#' {
				foundPlanet = true
				break
			}
		}
		if !foundPlanet {
			fmt.Println("Expanding column ", x)
			for idx := range planets {
				if planets[idx].originalX > x {
					planets[idx].expandX(expandBy)
				}
			}
		}
	}
	fmt.Println("Planets after expansion: ", planets)
	totalDist := 0
	for idx, p := range planets {
		for i := idx + 1; i < len(planets); i++ {
			dist := p.dist(&planets[i])
			fmt.Println("Calculating dist between planet ", idx, " and planet ", i, " as ", dist, ". For new total: ", totalDist)
			totalDist += dist
		}
	}
	fmt.Println("Total distance: ", totalDist)
}

func solvePart2(lines []string) {

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
	//solvePart(lines, 1)
	solvePart(lines, 1000000-1)
}
