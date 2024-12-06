package main

import (
	"bufio"
	"fmt"
	"guard"
	"io"
	"os"
)

func parseInput(path string) ([][]rune, int, int) {
    var inputData [][]rune
    file, _ := os.Open(path)
    reader := bufio.NewReader(file)
    var currentLine []rune
    var x, y, startingX, startingY int
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            inputData = append(inputData, currentLine)
            break
        }
        if char == '\n' {
            appendCopy := make([]rune, len(currentLine))
            copy(appendCopy, currentLine)
            inputData = append(inputData, appendCopy)
            currentLine = currentLine[:0]
            y += 1
            x = 0
        } else {
            if char == '^' {
                startingX, startingY = x, y
                char = '.' // we're not displaying the map, so we don't need to have the guard char in it
            }
            currentLine = append(currentLine, char)
            x += 1
        }
    }

    return inputData, startingX, startingY
}

func findAllInfiniteLoops(floorplan [][]rune, startingX int, startingY int) int {
    infiniteLoopCount := 0
    for x := range len(floorplan[0]) {
        for y := range len(floorplan) {
            if floorplan[y][x] != '#' && !(x == startingX && y == startingY) {
                floorplan[y][x] = '#'
                myGuard := guard.NewGuard(floorplan, startingX, startingY)
                if myGuard.Patrol() == -1 {
                    infiniteLoopCount += 1
                }
                floorplan[y][x] = '.'
            }
        }
    }
    return infiniteLoopCount
}

func main() {
    floorplan, startingX, startingY := parseInput("./input.txt")

    // Part 1
    myGuard := guard.NewGuard(floorplan, startingX, startingY)
    visitedLocationCount := myGuard.Patrol()
    fmt.Printf("(Part 1) - unique visited locations: %d\n", visitedLocationCount)

    // Part2
    possibleInfiniteLoops := findAllInfiniteLoops(floorplan, startingX, startingY)
    fmt.Printf("(Part 2) - possible infinite loops: %d\n", possibleInfiniteLoops)
}
