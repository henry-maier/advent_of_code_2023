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

type pos struct {
	x   int
	y   int
	dir rune
}

func (p pos) String() string {
	return fmt.Sprintf("x:%v, y:%v, dir:%s", p.x, p.y, string(p.dir))
}

func (p *pos) getConnectedPos(grid [][]rune) []pos {
	m := len(grid)
	n := len(grid[0])
	res := make([]pos, 0)
	// north
	tester := ' '
	if p.y-1 > 0 {
		tester = grid[p.y-1][p.x]
		if tester == '|' || tester == '7' || tester == 'F' {
			res = append(res, pos{x: p.x, y: p.y, dir: 'N'})
		}
	}
	// south
	if p.y+1 < m {
		tester = grid[p.y+1][p.x]
		if tester == '|' || tester == 'L' || tester == 'J' {
			res = append(res, pos{x: p.x, y: p.y, dir: 'S'})
		}
	}
	// east
	if p.x+1 < n {
		tester = grid[p.y][p.x+1]
		if tester == '-' || tester == 'J' || tester == '7' {
			res = append(res, pos{x: p.x, y: p.y, dir: 'E'})
		}
	}
	// west
	if p.x-1 > 0 {
		tester = grid[p.y][p.x-1]
		if tester == '-' || tester == 'L' || tester == 'F' {
			res = append(res, pos{x: p.x, y: p.y, dir: 'W'})
		}
	}
	return res
}

func (p *pos) getNextPos(grid [][]rune) pos {
	var nextPos pos
	var nextRune rune
	if p.dir == 'N' {
		nextPos = pos{y: p.y - 1, x: p.x}
		nextRune = grid[nextPos.y][nextPos.x]
		if nextRune == '|' {
			nextPos.dir = 'N'
		} else if nextRune == '7' {
			nextPos.dir = 'W'
		} else if nextRune == 'F' {
			nextPos.dir = 'E'
		} else {
			log.Fatal("Moving ", string(p.dir), " from ", p, "is invalid. Found ", string(nextRune), " at ", nextPos)
		}
	} else if p.dir == 'S' {
		nextPos = pos{y: p.y + 1, x: p.x}
		nextRune = grid[nextPos.y][nextPos.x]
		if nextRune == '|' {
			nextPos.dir = 'S'
		} else if nextRune == 'L' {
			nextPos.dir = 'E'
		} else if nextRune == 'J' {
			nextPos.dir = 'W'
		} else {
			log.Fatal("Moving ", string(p.dir), " from ", p, "is invalid. Found ", string(nextRune), " at ", nextPos)
		}
	} else if p.dir == 'W' {
		nextPos = pos{y: p.y, x: p.x - 1}
		nextRune = grid[nextPos.y][nextPos.x]
		if nextRune == '-' {
			nextPos.dir = 'W'
		} else if nextRune == 'L' {
			nextPos.dir = 'N'
		} else if nextRune == 'F' {
			nextPos.dir = 'S'
		} else {
			log.Fatal("Moving ", string(p.dir), " from ", p, "is invalid. Found ", string(nextRune), " at ", nextPos)
		}
	} else if p.dir == 'E' {
		nextPos = pos{y: p.y, x: p.x + 1}
		nextRune = grid[nextPos.y][nextPos.x]
		if nextRune == '-' {
			nextPos.dir = 'E'
		} else if nextRune == '7' {
			nextPos.dir = 'S'
		} else if nextRune == 'J' {
			nextPos.dir = 'N'
		} else {
			log.Fatal("Moving ", string(p.dir), " from ", p, "is invalid. Found ", string(nextRune), " at ", nextPos)
		}
	} else {
		log.Fatal("Invalid dir ", string(p.dir), " for pos ", p)
	}
	fmt.Println("Moved ", string(p.dir), " from ", p, " to get to ", nextPos, " facing ", string(nextPos.dir))
	return nextPos
}

func getStartPos(grid [][]rune) pos {
	m := len(grid)
	n := len(grid[0])
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if grid[y][x] == 'S' {
				return pos{x: x, y: y}
			}
		}
	}
	log.Fatal("Unreachable! Didn't find startPos")
	return pos{x: 0, y: 0}
}

func solvePart1(lines []string) {
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	startPos := getStartPos(grid)
	fmt.Println("Using starting pos ", startPos, "with val ", string(grid[startPos.y][startPos.x]))
	curr := startPos.getConnectedPos(grid)
	if len(curr) != 2 {
		log.Fatal("Got bad initial positions ", curr)
	}
	curr1 := curr[0]
	curr2 := curr[1]
	fmt.Println("Tracking two positions. 1: ", curr1, ", 2: ", curr2)
	steps := 0
	for {
		steps += 1
		fmt.Println("1:")
		curr1 = curr1.getNextPos(grid)
		if curr1.x == curr2.x && curr1.y == curr2.y {
			fmt.Println("Found meeting point ", curr1, " meets ", curr2, " after ", steps, " steps.")
			break
		}
		fmt.Println("2:")
		curr2 = curr2.getNextPos(grid)
		if curr1.x == curr2.x && curr1.y == curr2.y {
			fmt.Println("Found meeting point ", curr1, " meets ", curr2, " after ", steps, " steps.")
			break
		}
	}
	fmt.Println("Result: ", steps)
}

func solveLine2(line string) int {
	return 0
}

func solvePart2(lines []string) {
	var grid [][]rune
	resMap := make(map[pos]bool)
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	startPos := getStartPos(grid)
	resMap[pos{x: startPos.x, y: startPos.y}] = true
	fmt.Println("Using starting pos ", startPos, "with val ", string(grid[startPos.y][startPos.x]))
	curr := startPos.getConnectedPos(grid)
	if len(curr) != 2 {
		log.Fatal("Got bad initial positions ", curr)
	}
	curr1 := curr[0]
	curr2 := curr[1]
	fmt.Println("Tracking two positions. 1: ", curr1, ", 2: ", curr2)
	resMap[pos{x: curr1.x, y: curr1.y}] = true
	resMap[pos{x: curr2.x, y: curr2.y}] = true
	// lazy hack for real input, replace S with |
	grid[startPos.y][startPos.x] = '|'
	steps := 0
	for {
		steps += 1
		fmt.Println("1:")
		curr1 = curr1.getNextPos(grid)
		resMap[pos{x: curr1.x, y: curr1.y}] = true
		if curr1.x == curr2.x && curr1.y == curr2.y {
			fmt.Println("Found meeting point ", curr1, " meets ", curr2, " after ", steps, " steps.")
			break
		}
		fmt.Println("2:")
		curr2 = curr2.getNextPos(grid)
		resMap[pos{x: curr2.x, y: curr2.y}] = true
		if curr1.x == curr2.x && curr1.y == curr2.y {
			fmt.Println("Found meeting point ", curr1, " meets ", curr2, " after ", steps, " steps.")
			break
		}
	}
	fmt.Println("Found loop of length: ", steps)
	inLoopCount := 0
	m := len(grid)
	n := len(grid[0])
	insideLoop := false
	var r rune
	for y := 0; y < m; y++ {
		if insideLoop {
			log.Fatal("Failure. Said we were inside the loop at start of row ", y)
		}
		for x := 0; x < n; x++ {
			r = grid[y][x]
			if (resMap[pos{x: x, y: y}]) {
				// part of the loop
				fmt.Println("Intersecting loop at (", x, ", ", y, "). Found char ", string(r))
				if r == '|' {
					insideLoop = !insideLoop
					fmt.Println("Switching! Now insideLoop is", insideLoop)
				} else if r == 'F' {
					c := x + 1
					for ; c < n && grid[y][c] == '-'; c += 1 {
					}
					if grid[y][c] == 'J' {
						insideLoop = !insideLoop
						fmt.Println("Switching! Now insideLoop is", insideLoop)
					}
					x = c
				} else if r == 'L' {
					c := x + 1
					for ; c < n && grid[y][c] == '-'; c += 1 {
					}
					if grid[y][c] == '7' {
						insideLoop = !insideLoop
						fmt.Println("Switching! Now insideLoop is", insideLoop)
					}
					x = c
				}
			} else if insideLoop {
				fmt.Println("Found (", x, ", ", y, " to be inside the loop.")
				inLoopCount += 1
			}
		}
	}
	fmt.Println("Result: ", inLoopCount, " inside loop.")

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
