package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type antenna struct {
    frequency rune
    xLoc int
    yLoc int
}

func parseInput(path string) ([]antenna, int, int) {
    var antennas []antenna
    file, _ := os.Open(path)
    reader := bufio.NewReader(file)
    x, y := 0, 0
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            x -= 1
            break
        }
        if char == '\n' {
            x = 0
            y += 1
            continue
        } else if char == '.' {
            x += 1
            continue
        }
        antennas = append(antennas, antenna{char, x, y})
        x += 1
    }

    return antennas, x, y
}

func countUnique(stringSlice []string) int {
    stringMap := make(map[string]bool)
    for _, str := range stringSlice {
        if !stringMap[str] {
            stringMap[str] = true
        }
    }
    return len(stringMap)
}

func isValidLocation(x int, y int, maxX int, maxY int) bool {
    if x < 0 || x > maxX || y < 0 || y > maxY {
        return false
    }
    return true
}

func getAllResonantLocations(startingAntenna antenna, xDiff int, yDiff int, maxX int, maxY int) []string {
    var resonantLocations []string
    x, y := startingAntenna.xLoc, startingAntenna.yLoc
    for isValidLocation(x, y, maxX, maxY) {
        resonantLocations = append(resonantLocations, strconv.Itoa(x) + "," + strconv.Itoa(y))
        x, y = x + xDiff, y + yDiff
    }
    return resonantLocations
}

func countAntinodes(antennas []antenna, maxX int, maxY int, enableResonant bool) int {
    var antinodeLocations []string
    for i, antennaA := range antennas {
        for _, antennaB := range antennas[i + 1:] {
            if antennaA.frequency != antennaB.frequency {
                continue
            }
            
            xDiff, yDiff := antennaA.xLoc - antennaB.xLoc, antennaA.yLoc - antennaB.yLoc
            if enableResonant {
                antinodeLocationsFromA := getAllResonantLocations(antennaA, xDiff, yDiff, maxX, maxY)
                antinodeLocationsFromB := getAllResonantLocations(antennaB, -xDiff, -yDiff, maxX, maxY)
                antinodeLocations = append(antinodeLocations, antinodeLocationsFromA...)
                antinodeLocations = append(antinodeLocations, antinodeLocationsFromB...)
            } else {
                aN1x, aN1y := antennaA.xLoc + xDiff, antennaA.yLoc + yDiff
                if isValidLocation(aN1x, aN1y, maxX, maxY) {
                    antinodeLocationString := strconv.Itoa(aN1x) + "," + strconv.Itoa(aN1y)
                    antinodeLocations = append(antinodeLocations, antinodeLocationString)
                }
                aN2x, aN2y := antennaB.xLoc - xDiff, antennaB.yLoc - yDiff
                if isValidLocation(aN2x, aN2y, maxX, maxY) {
                    antinodeLocationString := strconv.Itoa(aN2x) + "," + strconv.Itoa(aN2y) 
                    antinodeLocations = append(antinodeLocations, antinodeLocationString)
                }
            }
        }
    }
    return countUnique(antinodeLocations)
}


func main() {
    antennas, maxX, maxY := parseInput("./input.txt")

    // Part1
    antinodesCount := countAntinodes(antennas, maxX, maxY, false)
    fmt.Printf("(Part 1) - total antinodes: %d\n", antinodesCount)

    // Part2
    antinodesCount = countAntinodes(antennas, maxX, maxY, true)
    fmt.Printf("(Part 2) - total resonant antinodes: %d\n", antinodesCount)
}