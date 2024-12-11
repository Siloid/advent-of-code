package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)


type stone struct {
    value int // value of stone
    count int // number of stones with same value
}

func (s *stone) blink() int { 
    // returns a value of a new right-side stone to create or -1 if no new stone
    //intValue, _ := strconv.Atoi(s.value)
    // rule 1
    if s.value == 0 {
        s.value = 1
        return -1
    }
    // rule 2
    valueString := strconv.Itoa(s.value)
    if len(valueString) % 2 == 0 {
        middleIndex := len(valueString) / 2
        newStoneValue, _ := strconv.Atoi(valueString[middleIndex:])
        s.value, _ = strconv.Atoi(valueString[:middleIndex])
        return newStoneValue
    }
    // rule 3
    s.value = s.value * 2024
    return -1
}

func parseInput(path string) []int {
    var stoneValues []int
    fileData, _ := os.ReadFile(path)
    intStrings := strings.Fields(string(fileData))
    for _, intString := range(intStrings) {
        value, _ := strconv.Atoi(intString)
        stoneValues = append(stoneValues, value)
    }
    return stoneValues
}

func blinkNTimesAndCount(stoneLine []stone, timesToBlink int) int { //[]stone {
    for range timesToBlink {
        // work from the right and append to the right to ensure correct ordering
        for i := len(stoneLine) -1 ; i >= 0 ; i-- {
            newStoneValue := stoneLine[i].blink()
            if newStoneValue != -1 {
                newStone := stone{newStoneValue, stoneLine[i].count}
                stoneLine = slices.Insert(stoneLine, i + 1, newStone)
            }
        }
        // reduce the stones
        for i := len(stoneLine) -1 ; i >= 0 ; i-- {
            for j, _ := range stoneLine {
                if j >= i {break}
                // if two stones have the same value, sum the counts and remove 1 stone
                if stoneLine[i].value == stoneLine[j].value {
                    stoneLine[j].count += stoneLine[i].count
                    stoneLine = slices.Delete(stoneLine, i, i+1)
                    break
                }
            }
        }
    }
    totalStones := 0
    for _, aStone := range stoneLine {
        totalStones += aStone.count
    }
    return totalStones
}

func createStoneLines(values []int) [][]stone {
    // each stone get it's own line as they don't interfere with one another
    // this way we can more easily parallelize in the next step
    var stoneLines [][]stone
    for _, value := range values {
        newStoneLine := []stone{stone{value, 1}}
        stoneLines = append(stoneLines, newStoneLine)
    }
    return stoneLines
}

func runIt(stoneLines [][]stone, timesToBlink int) int {
    var wg sync.WaitGroup
    results := make([]int, len(stoneLines))
    for index, stoneLine := range stoneLines {
        wg.Add(1)
        go func(i int, sl []stone) {
            defer wg.Done()
            results[i] = blinkNTimesAndCount(sl, timesToBlink)
        }(index, stoneLine)
    }
    wg.Wait()

    totalStones := 0
    for _, stoneCount := range results {
        totalStones += stoneCount
    }
    return totalStones
}

func main() {
    stoneValues := parseInput("input.txt")

    // Part1
    stoneLines := createStoneLines(stoneValues)
    totalStones := runIt(stoneLines, 25)
    fmt.Printf("(Part 1) - Stone total after 25 blinks: %d\n", totalStones)

    // Part2
    stoneLines = createStoneLines(stoneValues)
    totalStones = runIt(stoneLines, 75)
    fmt.Printf("(Part 2) - Stone total after 75 blinks: %d\n", totalStones)
}