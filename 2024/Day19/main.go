package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput(path string) ([]string, []string) {
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var patterns []string
    for scanner.Scan() {
        currentLine := scanner.Text()
        trimmedLine := strings.TrimSpace(currentLine)
        if trimmedLine == "" {
            break
        }
        patterns = append(patterns, strings.Split(trimmedLine, ", ")...)
    }

    var designs []string
    for scanner.Scan() {
        currentLine := scanner.Text()
        designs = append(designs, strings.TrimRight(currentLine, "\n"))
    }

    return patterns, designs
}

func computePossibleCount(design string, patterns []string, knownDesigns map[string]int) int {
    count, inMap := knownDesigns[design]
    if inMap {
        return count
    }
    count = 0
    for _, pattern := range patterns {
        subDesign, found := strings.CutPrefix(design, pattern)
        if !found {
            continue
        }
        if subDesign == "" {
            count += 1
        } else { 
            count += computePossibleCount(subDesign, patterns, knownDesigns)
        }
    }
    knownDesigns[design] = count
    return count
}
    
func populateKnownDesigns(patterns []string, designs []string, knownDesigns map[string]int) {
    for _, design := range designs {
        computePossibleCount(design, patterns, knownDesigns)
    }
}

func getPossibleDesignCount(designs []string, knownDesigns map[string]int, countUnique bool) int {
    total := 0
    for _, design := range designs {
        count, _ := knownDesigns[design]
        if countUnique {
            total += count
        } else if count > 0 {
            total += 1
        }
    }
    return total
}

func main() {
    patterns, designs := parseInput("input.txt")
    knownDesigns := make(map[string]int)

    // Part1
    populateKnownDesigns(patterns, designs, knownDesigns)
    fmt.Printf("(Part 1) - Possible designs: %d\n", getPossibleDesignCount(designs, knownDesigns, false))

    // Part2
    fmt.Printf("(Part 2) - Possible unique design arrangements: %d\n", getPossibleDesignCount(designs, knownDesigns, true))
}