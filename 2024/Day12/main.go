package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)


type plot struct {
    crop rune
    area int
    perimeter int
    sides int
}

func (p *plot) fenceCost()int {
    return p.area * p.perimeter
}

type coordinate struct {
    x int
    y int
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

func (p plot) String() string {
    return fmt.Sprintf("Plot: {crop: %s, a: %d, p: %d}\n", string(p.crop), p.area, p.perimeter)
}

func getPerimeter(gardenData [][]rune, loc coordinate) int {
    perimeter := 0
    var adjacentCrops []rune
    if loc.x > 0 {
        adjacentCrops = append(adjacentCrops, gardenData[loc.y][loc.x - 1])
    } else {perimeter += 1}
    if loc.x + 1 < len(gardenData[0]) {
        adjacentCrops = append(adjacentCrops, gardenData[loc.y][loc.x + 1])
    } else {perimeter += 1}
    if loc.y > 0 {
        adjacentCrops = append(adjacentCrops, gardenData[loc.y - 1][loc.x])
    } else {perimeter += 1}
    if loc.y + 1 < len(gardenData) {
        adjacentCrops = append(adjacentCrops, gardenData[loc.y + 1][loc.x])
    } else {perimeter += 1}
    for _, crop := range adjacentCrops {
        if crop != gardenData[loc.y][loc.x] {
            perimeter += 1
        }
    }
    return perimeter
}

func findGardenPlot(crop rune, gardenData [][]rune, plotted [][]int, loc coordinate) (int, int) {
    if gardenData[loc.y][loc.x] != crop || plotted[loc.y][loc.x] == 1 {
        return 0, 0
    }
    plotted[loc.y][loc.x] = 1
    area := 1
    perimeter := getPerimeter(gardenData, loc)
    // gotta be a better way to do this...
    var coordinatesToCheck []coordinate
    if loc.x > 0 {
        coordinatesToCheck = append(coordinatesToCheck, coordinate{loc.x - 1, loc.y})
    }
    if loc.x + 1 < len(gardenData[0]) {
        coordinatesToCheck = append(coordinatesToCheck, coordinate{loc.x + 1, loc.y})

    } 
    if loc.y > 0 {
        coordinatesToCheck = append(coordinatesToCheck, coordinate{loc.x, loc.y - 1})

    } 
    if loc.y + 1 < len(gardenData) {
        coordinatesToCheck = append(coordinatesToCheck, coordinate{loc.x, loc.y + 1})
    } 
    for _, coordinateToCheck := range coordinatesToCheck {
        additionalArea, additionalPermiter := findGardenPlot(crop, gardenData, plotted, coordinateToCheck)
        area += additionalArea
        perimeter += additionalPermiter
    }
    return area, perimeter
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
                area, perimeter := findGardenPlot(gardenData[y][x], gardenData, plotted, coordinate{x, y})
                newPlot := plot{gardenData[y][x], area, perimeter, 0}
                plots = append(plots, newPlot)
            }
        }
    }

    return plots
}

func sumFenceCost(plots []plot) int {
    total := 0
    for _, p := range plots {
        total += (p.fenceCost())
    }
    return total
}

func main() {
    gardenData := readInput("input.txt")
    plots := createPlots(gardenData)

    //fmt.Println(plots)
    // Part1
    totalFenceCost := sumFenceCost(plots)
    fmt.Printf("(Part 1) - Total fence cost: %d\n", totalFenceCost)
}