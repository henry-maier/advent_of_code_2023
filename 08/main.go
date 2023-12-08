package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseMapLine(line string) (string, string, string) {
	var nodeFound, leftFound bool
	var node, left, right string
	for _, char := range line {
		switch char {
		case '=':
			nodeFound = true
		case ',':
			leftFound = true
		case ' ', '(', ')':
			continue
		default:
			if leftFound {
				right += string(char)
			} else if nodeFound {
				left += string(char)
			} else {
				node += string(char)
			}
		}
	}
	return node, left, right
}

func getLRMaps(mapLines []string) (map[string]string, map[string]string) {
	var node, left, right string
	lMap := make(map[string]string)
	rMap := make(map[string]string)
	for _, line := range mapLines {
		node, left, right = parseMapLine(line)
		lMap[node] = left
		rMap[node] = right
	}
	return lMap, rMap
}

func getStartAndEndNodes(m map[string]string) ([]string, map[string]bool) {
	startNodes := make([]string, 0)
	endNodes := make(map[string]bool)
	var endingChar rune
	for k, _ := range m {
		endingChar = getEndChar(k)
		if endingChar == 'A' {
			startNodes = append(startNodes, k)
		} else if endingChar == 'Z' {
			endNodes[k] = true
		}
	}
	return startNodes, endNodes
}

func getEndChar(s string) rune {
	runeArr := []rune(s)
	return runeArr[len(runeArr)-1]
}

func solvePart2(instructions string, mapLines []string) {
	leftMap, rightMap := getLRMaps(mapLines)
	nodePositions, endingNodes := getStartAndEndNodes(leftMap)
	fmt.Println("startnodes: ", nodePositions)
	fmt.Println("endnodes: ", endingNodes)
	instructionsArr := []rune(instructions)
	var direction rune
	var newPos string
	res := 0
	cycles := make([]int, len(nodePositions))
	for cyclesFound := 0; cyclesFound < len(nodePositions); {
		direction = instructionsArr[res%len(instructionsArr)]
		for idx, pos := range nodePositions {
			if direction == 'L' {
				newPos = leftMap[pos]
			} else if direction == 'R' {
				newPos = rightMap[pos]
			} else {
				log.Fatal("Invalid direction ", direction)
			}
			fmt.Println(res+1, ": ", "Went ", string(direction), " from ", pos, " to ", newPos)
			if cycles[idx] == 0 && getEndChar(newPos) == 'Z' {
				cycles[idx] = res + 1
				cyclesFound += 1
			}
			nodePositions[idx] = newPos
		}
		res += 1
	}
	fmt.Println("found cycles: ", cycles)
	fmt.Println("Result: ", lcm(cycles))
}

func lcm(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	newNums := make([]int, len(nums)-1)
	for idx, num := range nums[1:] {
		newNums[idx] = num / gcd(num, nums[0])
	}
	return nums[0] * lcm(newNums)
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}

}

func solvePart1(instructions string, mapLines []string) {
	leftMap, rightMap := getLRMaps(mapLines)
	fmt.Println("leftmap: ", leftMap)
	fmt.Println("rightmap: ", rightMap)
	res := 0
	instructionsArr := []rune(instructions)
	var direction rune
	var newPos string
	for pos := "AAA"; pos != "ZZZ"; {
		direction = instructionsArr[res%len(instructionsArr)]
		if direction == 'L' {
			newPos = leftMap[pos]
		} else if direction == 'R' {
			newPos = rightMap[pos]
		} else {
			log.Fatal("Invalid direction ", direction)
		}
		fmt.Println(res+1, ": ", "Went ", string(direction), " from ", pos, " to ", newPos)
		pos = newPos
		res += 1
	}
	fmt.Println("Result: ", res)
}

func main() {
	// thanks https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	instructions := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		instructions += line
	}
	var mapLines []string
	for scanner.Scan() {
		mapLines = append(mapLines, scanner.Text())
	}
	//solvePart1(instructions, mapLines)
	solvePart2(instructions, mapLines)
}
