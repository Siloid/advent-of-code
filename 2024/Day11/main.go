package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)


type stone struct {
    value int
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

func blinkNTimes(stoneLine []stone, timesToBlink int) []stone {
    for range timesToBlink {
        // work from the right and append to the right to ensure correct ordering
        for i := len(stoneLine) -1 ; i >= 0 ; i-- {
            newStoneValue := stoneLine[i].blink()
            if newStoneValue != -1 {
                newStone := stone{newStoneValue}
                stoneLine = slices.Insert(stoneLine, i + 1, newStone)
            }
        }
    }
    return stoneLine
}

func createStones(values []int) []stone {
    var stoneLine []stone
    for _, value := range values {
        newStone := stone{value}
        stoneLine = append(stoneLine, newStone)
    }
    return stoneLine
}

func main() {
    stoneValues := parseInput("input.txt")
    stoneLine := createStones(stoneValues)
    
    // Part1
    stoneLine = blinkNTimes(stoneLine, 25)
    fmt.Printf("(Part 1) - Stone total after 25 blinks: %d\n", len(stoneLine))
}