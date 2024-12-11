package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/dominikbraun/graph"
)

func parseInput(path string) [][]int {
    var inputData [][]int
    file, _ := os.Open(path)
    reader := bufio.NewReader(file)
    var currentLine []int
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            inputData = append(inputData, currentLine)
            break
        }
        if char == '\n' {
            appendCopy := make([]int, len(currentLine))
            copy(appendCopy, currentLine)
            inputData = append(inputData, appendCopy)
            currentLine = currentLine[:0]
        } else {
            value, _ := strconv.Atoi(string(char))
            currentLine = append(currentLine, value)
        }
    }
    return inputData
}

func parseMap(topoMap [][]int) (graph.Graph[string, string], []string, []string) {
    g := graph.New(graph.StringHash, graph.Directed())
    var trailheads []string
    var summits []string
    for x := range(len(topoMap[0])) {
        for y := range(len(topoMap)) {
            vertexAStr := string(strconv.Itoa(x)) + "," + string(strconv.Itoa(y))
            mapValue := topoMap[y][x]
            if mapValue == 0 {
                trailheads = append(trailheads, vertexAStr)
            } else if mapValue == 9 {
                summits = append(summits, vertexAStr)
            }
            g.AddVertex(vertexAStr)
            //Only look left and up, as we'll add the right and down edges when we add those vertices
            if x - 1 >= 0 {
                vertexBStr := string(strconv.Itoa(x-1)) + "," + string(strconv.Itoa(y))
                if isHikeable(mapValue, topoMap[y][x-1]) {
                    g.AddEdge(vertexBStr, vertexAStr)
                } else if isHikeable(topoMap[y][x-1], mapValue) {
                    g.AddEdge(vertexAStr, vertexBStr)
                }
            }
            if y - 1 >= 0 {
                vertexBStr := string(strconv.Itoa(x)) + "," + string(strconv.Itoa(y-1))
                if isHikeable(mapValue, topoMap[y-1][x]) {
                    g.AddEdge(vertexBStr, vertexAStr)
                } else if isHikeable(topoMap[y-1][x], mapValue) {
                    g.AddEdge(vertexAStr, vertexBStr)
                }
            }
        }
    }
    return g, trailheads, summits
}

func isHikeable(heightA int, heightB int) bool {
    // can only hike uphill and in steps of 1
    grade := heightA - heightB
    return grade == 1
}

func summitsReachable(gMap graph.Graph[string, string], trailhead string, summits []string, useRating bool) int {
    total := 0
    for _, summit := range summits {
        paths, _ := graph.AllPathsBetween(gMap, trailhead, summit)
        if useRating {
            total += len(paths)
        } else if len(paths) > 0 { //if not using rating, if any path exists we only add 1 for being able to reach that summit
            total += 1
        }
    }
    return total
}

func sumMapPaths(gMap graph.Graph[string, string], trailheads []string, summits []string, useRating bool) int {
    total := 0
    for _, trailhead := range trailheads {
        total += summitsReachable(gMap, trailhead, summits, useRating)
    }
    return total
}

func main() {
    mapData := parseInput("input.txt")
    mapGraph, trailheads, summits := parseMap(mapData)

    totalTrails := sumMapPaths(mapGraph, trailheads, summits, false)
    fmt.Printf("(Part 1) - total trails: %d\n", totalTrails)

    totalTrails = sumMapPaths(mapGraph, trailheads, summits, true)
    fmt.Printf("(Part 2) - total trails using rating: %d\n", totalTrails)
}