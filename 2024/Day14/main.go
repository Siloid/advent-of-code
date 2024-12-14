package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	roomX = 100 // 101 spaces 0 -> 100
	roomY = 102 // 103 spaces 0 -> 102

	outputDir = "./tree_ouput"

	// quadrant boundaries
	uL, uR, lL, lR                 = 0, 1, 2, 3
	uLminX, uLminY, uLmaxX, uLmaxY = 0, 0, int((roomX - 1) / 2), int((roomY - 1) / 2)
	uRminX, uRminY, uRmaxX, uRmaxY = int(roomX/2) + 1, 0, roomX, int((roomY - 1) / 2)
	lLminX, lLminY, lLmaxX, lLmaxY = 0, int(roomY/2) + 1, int((roomX - 1) / 2), roomY
	lRminX, lRminY, lRmaxX, lRmaxY = int(roomX/2) + 1, int(roomY/2) + 1, roomX, roomY
)

type xy struct {
	x int
	y int
}

type robot struct {
	position xy
	velocity xy
}

func (r *robot) move(times int) {
	newX := (r.position.x + (r.velocity.x * times)) % (roomX + 1)
	newY := (r.position.y + (r.velocity.y * times)) % (roomY + 1)
	if newX < 0 {
		newX = roomX + newX + 1
	}
	if newY < 0 {
		newY = roomY + newY + 1
	}
	r.position.x = newX
	r.position.y = newY
}

func (r robot) getQuadrant() int {
	// returns quadrant in which this robot resides or -1 if between quadrants
	if r.position.x >= uLminX && r.position.y >= uLminY && r.position.x <= uLmaxX && r.position.y <= uLmaxY {
		return uL
	}
	if r.position.x >= uRminX && r.position.y >= uRminY && r.position.x <= uRmaxX && r.position.y <= uRmaxY {
		return uR
	}
	if r.position.x >= lLminX && r.position.y >= lLminY && r.position.x <= lLmaxX && r.position.y <= lLmaxY {
		return lL
	}
	if r.position.x >= lRminX && r.position.y >= lRminY && r.position.x <= lRmaxX && r.position.y <= lRmaxY {
		return lR
	}
	return -1
}

func parseInput(path string) []robot {
	var robots []robot
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		xyRegexp := regexp.MustCompile(`[-]*\d+,[-]*\d+`)
		details := xyRegexp.FindAllString(currentLine, -1)
		position := strings.Split(details[0], ",")
		velocity := strings.Split(details[1], ",")
		pX, _ := strconv.Atoi(position[0])
		pY, _ := strconv.Atoi(position[1])
		vX, _ := strconv.Atoi(velocity[0])
		vY, _ := strconv.Atoi(velocity[1])
		robots = append(robots, robot{xy{pX, pY}, xy{vX, vY}})
	}
	file.Close()
	return robots
}

func calculateSafety(robots []robot) int {
	uLCount, uRCount, lLCount, lRCount := 0, 0, 0, 0
	for _, r := range robots {
		quadrant := r.getQuadrant()
		if quadrant == uL {
			uLCount += 1
			continue
		}
		if quadrant == uR {
			uRCount += 1
			continue
		}
		if quadrant == lL {
			lLCount += 1
			continue
		}
		if quadrant == lR {
			lRCount += 1
			continue
		}
	}
	return uLCount * uRCount * lLCount * lRCount
}

func printRobots(robots []robot, iteration int) {
	var screen [][]rune
	for range roomY + 1 {
		line := make([]rune, roomX+1)
		for x := range roomX + 1 {
			line[x] = ' '
		}
		screen = append(screen, line)
	}

	for _, r := range robots {
		screen[r.position.y][r.position.x] = '#'
	}

	f, _ := os.OpenFile(outputDir+"/tree"+strconv.Itoa(iteration)+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	f.WriteString("\n\nIteration: " + strconv.Itoa(iteration) + "\n")
	for i := range screen {
		f.WriteString(string(screen[i]) + "\n")
	}
}

func runPart2(robots []robot, iterations int) {
	os.Mkdir(outputDir, 0755)
	robots = parseInput("./input.txt")
	printRobots(robots, 0)
	for i := range iterations {
		for j, _ := range robots {
			robots[j].move(1)
		}
		printRobots(robots, i+1)
	}
}

func main() {
	robots := parseInput("./input.txt")

	// Part1
	timesToMove := 100
	for i, _ := range robots {
		robots[i].move(timesToMove)
	}
	safetyFactor := calculateSafety(robots)
	fmt.Printf("(Part 1) - Safety Factor: %d\n", safetyFactor)

	// Part2
	robots = parseInput("./input.txt")
	fmt.Println("(Part 2) - To run Part 2, uncomment the line below.")
	fmt.Println("  Warning this will create thousands of files in: " + outputDir)
	// This creates the first 10K iterations.  For my input iteration 7672 created a tree.
	// A different number of iterations may be needed for a different input

	// Uncomment next line to run Part 2
	//runPart2(robots, 10000)
}
