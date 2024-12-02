package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxSafeDiff = 3

func parseInput(path string) ([][]int) {
    var reportData [][]int
    
    file, _ := os.Open(path)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        currentLine := scanner.Text()
        levels := strings.Fields((currentLine))

        var report []int
        for _, value := range levels {
            level, _ := strconv.Atoi(value)
            report = append(report, level)
        }
        reportData = append(reportData, report)
    }

    return reportData
}

func removeLevel(report []int, index int) []int {
    reportCopy := make([]int, len(report))
    copy(reportCopy, report)
    return append(reportCopy[:index], reportCopy[index+1:]...)
}

func isReportSafe(report []int, dampenerEnabled bool) bool {
    if dampenerEnabled {
        // If removing any single level makes it safe then it's safe
        for i, _ := range report {
            if isReportSafe(removeLevel(report, i), false) {
                return true
            }
        }
        return false
    }

    ascending := true
    if report[0] > report[1] {
        ascending = false
    }
    currentLevel := report[0]
    for _, nextLevel := range report[1:] {
        var diff int
        if ascending {
            diff = nextLevel - currentLevel
        } else {
            diff = currentLevel - nextLevel
        }
        if diff > maxSafeDiff || diff <= 0 {
            return false
        }
        currentLevel = nextLevel
    }
    return true
}

func main() {
    reports := parseInput("./input.txt")
    
    // Part1
    var safeCounter int = 0
    for _, report := range reports {
        if isReportSafe(report, false) {
            safeCounter += 1
        }
    }
    fmt.Printf("(Part 1) - Total Safe Reports: %d\n", safeCounter)

    // Part2
    safeCounter = 0
    for _, report := range reports {
        if isReportSafe(report, true) {
            safeCounter += 1
        }
    }
    fmt.Printf("(Part 2) - Total Safe Reports w/ Dampener: %d\n", safeCounter)
}