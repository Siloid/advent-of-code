package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
)

const (
    mapSizeX = 71
    mapSizeY = 71
)

type xy struct {
    x int
    y int
}

func parseInput(path string) []xy {
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var bytes []xy
    for scanner.Scan() {
        currentLine := scanner.Text()
        xyStrs := strings.Split(strings.TrimSpace(currentLine), ",")
        x, _ := strconv.Atoi(xyStrs[0])
        y, _ := strconv.Atoi(xyStrs[1])
        bytes = append(bytes, xy{x, y})
    }
    return bytes
}

func getVertexString(x int, y int) string {
    return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func createGraphAfterXFallenBytes(bytes []xy, fallenByteCount int) (graph.Graph[string, string], string, string) {
    layout := make([][]int, mapSizeY)
    for i := range layout {
        layout[i] = make([]int, mapSizeX)
    }
    for i := range fallenByteCount {
        fallingByte := bytes[i]
        layout[fallingByte.y][fallingByte.x] = 1
    }

    g := graph.New(graph.StringHash)
    start := xy{0, 0}
    end := xy{mapSizeX - 1, mapSizeY - 1}
    for x := range(len(layout[0])) {
        for y := range(len(layout)) {
            if layout[y][x] == 1 {
                continue
            }
            vertexStr := getVertexString(x, y)
            g.AddVertex(vertexStr)
            //Only look left and up, as we'll add the right and down edges when we add those vertices
            if x - 1 >= 0 {
                adjacentVertexStr := getVertexString(x - 1, y)
                g.AddEdge(vertexStr, adjacentVertexStr)
            }
            if y - 1 >= 0 {
                adjacentVertexStr := getVertexString(x, y - 1)
                g.AddEdge(vertexStr, adjacentVertexStr)
            }
        }
    }
    return g, getVertexString(start.x, start.y), getVertexString(end.x, end.y)
}

func findFirstBlockingByte(bytes []xy) xy {
    min, max := 0, len(bytes) - 1
    for min != max - 1 {
        current := int((min + max) / 2)
        g, start, end := createGraphAfterXFallenBytes(bytes, current)
        path, _ := graph.ShortestPath(g, start, end)
        if len(path) != 0 {
            min = current
        } else {
            max = current
        }
    }
    return bytes[max - 1]
}

func main() {
    bytes := parseInput("input.txt")

    // Part1
    space, start, end := createGraphAfterXFallenBytes(bytes, 1024)
    shortestPath, _ := graph.ShortestPath(space, start, end)
    fmt.Printf("(Part 1) - Fewest steps after 1024 fallen bytes: %d\n", len(shortestPath) - 1) // -1 because they wanted steps, so don't count starting vertex

    // Part2
    blockingByte := findFirstBlockingByte(bytes)
    fmt.Printf("(Part 2) - First blocking byte: %d,%d\n", blockingByte.x, blockingByte.y)
}