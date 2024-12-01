package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(path string) ([]int, []int, map[int]int) {
    var leftLocationIDs []int
    var rightLocationIDs []int
    similarityMap := make(map[int]int)

    file, _ := os.Open(path)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        currentLine := scanner.Text()
        locations := strings.Fields((currentLine))

        leftLocation, _ := strconv.Atoi(locations[0])
        leftLocationIDs = append(leftLocationIDs, leftLocation)

        rightLocation, _ := strconv.Atoi(locations[1])
        rightLocationIDs = append(rightLocationIDs, rightLocation)

        count, inMap := similarityMap[rightLocation]
        if inMap {
            similarityMap[rightLocation] = count + 1
        } else {
            similarityMap[rightLocation] = 1
        }
    }

    sort.Ints(leftLocationIDs)
    sort.Ints(rightLocationIDs)

    return leftLocationIDs, rightLocationIDs, similarityMap
}

func absDiff(a int, b int) int {
    if a < b {
        return b - a
    }
    return a - b
}

func similarity(id int, similarityMap map[int]int) int {
    similityMultiplier, inMap := similarityMap[id]
    if inMap {
        return similityMultiplier * id
    } else {
        return 0
    }
}

func main() {
    leftIDs, rightIDs, similarityMap := parseInput("./input.txt")

    var totalDifference int = 0
    var totalSimilarity int = 0
    for index, id := range leftIDs {
        // Part 1
        totalDifference += absDiff(id, rightIDs[index])

        // Part 2
        totalSimilarity += similarity(id, similarityMap)
    }

    fmt.Printf("(Part 1) - Total Difference: %d\n", totalDifference)
    fmt.Printf("(Part 2) - Total Similarity: %d\n", totalSimilarity)
}
