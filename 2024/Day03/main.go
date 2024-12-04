package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(path string) string {
    fileData, _ := os.ReadFile(path)

    memory := string(fileData)
    return memory
}

func getMultiple(multiplier string) int {
    // We've already found the valid data, so know we'll get exactly 2 numbers
    numberRegexp := regexp.MustCompile(`\d+`)
    numbers := numberRegexp.FindAllString(multiplier, -1)

    firstNumber, _ := strconv.Atoi(numbers[0])
    secondNumber, _ := strconv.Atoi(numbers[1])

    return firstNumber * secondNumber
}

func sumAllMultipliers(multipliers []string) int {
    multiplierTotal := 0
    for _, multiplier := range multipliers {
        multiplierTotal += getMultiple(multiplier)
    }
    return multiplierTotal
} 

func getAllMultipliers(memory string) []string {
    mulRegexp := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
    multipliers := mulRegexp.FindAllString(memory, -1)
    return multipliers
}

func getAllDoMemorySegments(memory string) []string {
    // A regex such as `do\(\).*?(don't\(\)|$)` should handle this but it wasn't and I didn't 
    // feel like spending more time on it... therefore... behold
    // Splitting on do() and then don't() allows us to keep the first set of each (the do's)
    // and discard the rest (the dont's)
    var doMemorySegments []string
    doSplit := "do()"
    segments := strings.Split(memory, doSplit)
    for _, segment := range segments {
        dontSplit := "don't()"
        doSegment := strings.Split(segment, dontSplit)
        doMemorySegments = append(doMemorySegments, doSegment[0])
    }
    return doMemorySegments
}

func main() {
    corruptedMemory := parseInput("./input.txt")

    // Part 1
    multipliers := getAllMultipliers(corruptedMemory)
    multiplierTotal := sumAllMultipliers(multipliers)
    fmt.Printf("(Part 1) - Sum of Multipliers: %d\n", multiplierTotal)

    // Part 2
    multiplierTotal = 0
    doMemorySegments := getAllDoMemorySegments(corruptedMemory)
    for _, memorySegement := range doMemorySegments {
        multipliers = getAllMultipliers(memorySegement)
        multiplierTotal += sumAllMultipliers(multipliers)
    }
    fmt.Printf("(Part 2) - Sum of all \"do()\" Multipliers: %d\n", multiplierTotal)
}