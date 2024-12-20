package main

import (
	"bufio"
	"fmt"
	"os"
)

type xy struct {
    x int
    y int
}

func parseInput(path string) ([][]rune, xy, xy) {
    var racetrack [][]rune
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    y := 0
    var start, end xy    
    for scanner.Scan() {
        currentLine := scanner.Text()
        for x, char := range currentLine {
            if char == 'S' {
                start.x, start.y = x, y
            } else if char == 'E' {
                end.x, end.y = x, y
            }
        }
        racetrack = append(racetrack, []rune(currentLine))
        y += 1
    }
    racetrack[start.y][start.x] = '.'
    racetrack[end.y][end.x] = '.'
    return racetrack, start, end
}

func getStandardPath(racetrack [][]rune, start xy, end xy) []xy {
    var path []xy
    path = append(path, start)
    loc := xy{start.x, start.y}
    racetrack[loc.y][loc.x] = 'O'
    for loc.x != end.x || loc.y != end.y {
        if racetrack[loc.y - 1][loc.x] == '.' { // up
            loc.y -= 1
        } else if racetrack[loc.y + 1][loc.x] == '.' {// down
            loc.y += 1
        } else if racetrack[loc.y][loc.x - 1] == '.' { // left
            loc.x -= 1
        } else if racetrack[loc.y][loc.x + 1] == '.' { // right
            loc.x += 1
        }
        path = append(path, xy{loc.x, loc.y})
        racetrack[loc.y][loc.x] = 'O'
    }
    return path
}

func absDiff(a int, b int) int {
    c := a - b
    if c < 0 {
        return -c
    }
    return c
}

func isValidCheat(locA xy, locB xy, cheatTime, minTimeSaved int, savings int) bool {
    totalDiff := absDiff(locA.x, locB.x) + absDiff(locA.y, locB.y)
    if totalDiff <= cheatTime && savings - totalDiff >= minTimeSaved {
        return true
    }
    return false
}

func findNumberOfCheats(path []xy, minTimeSaved int, cheatTime int) int {
    numCheats := 0
    for i := range len(path) - minTimeSaved {
        for j := i+minTimeSaved; j < len(path); j++ {
            if isValidCheat(path[i], path[j], cheatTime, minTimeSaved, j-i) {
                numCheats += 1
            }
        }
    }
    return numCheats
}

func main() {
    racetrack, start, end := parseInput("input.txt")

    // Part1
    standardPath := getStandardPath(racetrack, start, end)
    fmt.Printf("(Part 1) - Possible 2 picosecond cheats: %d\n", findNumberOfCheats(standardPath, 100, 2))

    // Part2
    fmt.Printf("(Part 2) - Possible 20 picosecond cheats: %d\n", findNumberOfCheats(standardPath, 100, 20))
}