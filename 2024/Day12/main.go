package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type coordinate struct {
    x int
    y int
}

type fence struct {
    location coordinate
    direction int
}

type plot struct {
    crop rune
    area int
    perimeter int
    sides int
    fences []fence
}

func createPlot(crop rune) plot {
    return plot{crop, 0, 0, 0, make([]fence, 0)}
}

func (p *plot) fenceCost()int {
    return p.area * p.perimeter
}

func (p *plot) fenceBulkCost()int {
    return p.area * p.sides
}

func (p *plot) computeSides() {
    // This is not efficient and is called for every section add
    // gotta be a better way
    p.sides = 0
    maxX, maxY := 0, 0
    for _, fence := range p.fences {
        if fence.location.x > maxX {maxX = fence.location.x}
        if fence.location.y > maxY {maxY = fence.location.y}
    }
    var checkedFences []fence
    for y := 0; y <= maxY; y++ {
        for x := 0; x <= maxX; x++ {
            for _, nextFence := range p.fences {
                if nextFence.location.x == x && nextFence.location.y == y {
                    areConnected := false
                    for _, checkedFence := range checkedFences {
                        if areConnectedFences(nextFence, checkedFence) {
                            areConnected = true
                            break
                        }
                    }
                    if !areConnected {
                        p.sides += 1
                    }
                    checkedFences = append(checkedFences, nextFence)
                }
            }
        }
    }
}

func (p *plot) addSection(gardenData [][]rune, location coordinate) {
    newFences := getFences(gardenData, location)
    p.area += 1
    p.perimeter += len(newFences)
    p.fences = append(p.fences, newFences...)
    p.computeSides()
}

func readInput(path string) [][]rune {
    var gardenData [][]rune
    file, _ := os.Open(path)
    reader := bufio.NewReader(file)
    var currentLine []rune
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            gardenData = append(gardenData, currentLine)
            break
        }
        if char == '\n' {
            appendCopy := make([]rune, len(currentLine))
            copy(appendCopy, currentLine)
            gardenData = append(gardenData, appendCopy)
            currentLine = currentLine[:0]
            continue
        }
        currentLine = append(currentLine, char)
    }
    return gardenData
}

func areConnectedFences(fenceA fence, fenceB fence) bool {
    if fenceA.direction == fenceB.direction {
        adjacentLocations := []coordinate{{fenceA.location.x - 1, fenceA.location.y},
                                          {fenceA.location.x + 1, fenceA.location.y},
                                          {fenceA.location.x, fenceA.location.y - 1},
                                          {fenceA.location.x, fenceA.location.y + 1}}
        for _, location := range adjacentLocations {
            if location.x == fenceB.location.x && location.y == fenceB.location.y {
                return true
            }
        }
    }
    return false
}

func getFenceDirection(locA coordinate, locB coordinate) int {
    // 0,1,2,3 -> up, right, down, left
    if locA.x == locB.x {
        if locA.y < locB.y {return 2}
        return 0
    }
    if locA.x < locB.x {return 1}
    return 3
}

func getFences(gardenData [][]rune, loc coordinate) []fence {
    var fences []fence
    cardinalCordinates, edgeOfMapCoordinates := getCardinalCoordinates(gardenData, loc)
    for _, coordinateToCheck := range cardinalCordinates {
        if gardenData[loc.y][loc.x] != gardenData[coordinateToCheck.y][coordinateToCheck.x] {
            fences = append(fences, fence{loc, getFenceDirection(loc, coordinateToCheck)})
        }
    }
    for _, edgeCoordinate := range edgeOfMapCoordinates {
        fences = append(fences, fence{loc, getFenceDirection(loc, edgeCoordinate)})
    }
    return fences
}

func getCardinalCoordinates(gardenData [][]rune, loc coordinate) ([]coordinate, []coordinate) {
    var cardinalCoordinates []coordinate
    var invalidCoordinates []coordinate
    coordinates := []coordinate{{loc.x - 1, loc.y}, {loc.x + 1, loc.y}, {loc.x, loc.y - 1}, {loc.x, loc.y + 1}}
    for _, c := range coordinates {
        if c.x < 0 || c.x >= len(gardenData[0]) || c.y < 0 || c.y >= len(gardenData) {
            invalidCoordinates = append(invalidCoordinates, c)
        } else {
            cardinalCoordinates = append(cardinalCoordinates, c)
        }
    }
    return cardinalCoordinates, invalidCoordinates
}

func mapGardenPlot(aPlot *plot, gardenData [][]rune, plotted [][]int, loc coordinate) {
    if plotted[loc.y][loc.x] == 1 {return} // only plot each space once
    aPlot.addSection(gardenData, loc)
    plotted[loc.y][loc.x] = 1

    coordinatesToCheck, _ := getCardinalCoordinates(gardenData, loc)
    for _, coordinateToCheck := range coordinatesToCheck {
        if aPlot.crop == gardenData[coordinateToCheck.y][coordinateToCheck.x] {
            mapGardenPlot(aPlot, gardenData, plotted, coordinateToCheck)
        }
    }
}

func createPlots(gardenData [][]rune) []plot {
    var plots []plot

    plotted := make([][]int, len(gardenData))
    for i := range plotted {
        plotted[i] = make([]int, len(gardenData[0])) 
    }
 
    for x := range len(gardenData[0]) {
        for y := range len(gardenData) {
            if plotted[y][x] == 0 {
                newPlot := createPlot(gardenData[y][x])
                mapGardenPlot(&newPlot, gardenData, plotted, coordinate{x, y})
                plots = append(plots, newPlot)
            }
        }
    }

    return plots
}

func computeFenceCosts(plots []plot) (int, int) {
    total, totalBulk := 0, 0
    for _, p := range plots {
        total += (p.fenceCost())
        totalBulk += (p.fenceBulkCost())
    }
    return total, totalBulk
}

func main() {
    gardenData := readInput("input.txt")
    plots := createPlots(gardenData)

    // Part1 and Part2
    totalFenceCost, totalFenceBulkCost := computeFenceCosts(plots)
    fmt.Printf("(Part 1) - Total fence cost: %d\n", totalFenceCost)
    fmt.Printf("(Part 2) - Total fence bulk cost: %d\n", totalFenceBulkCost)
}