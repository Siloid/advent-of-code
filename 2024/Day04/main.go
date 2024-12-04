package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func parseInput(path string) ([][]rune) {
    var inputData [][]rune
    file, _ := os.Open(path)
    reader := bufio.NewReader(file)
    var currentLine []rune
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            inputData = append(inputData, currentLine)
            break
        }
        if char == '\n' {
            appendCopy := make([]rune, len(currentLine))
            copy(appendCopy, currentLine)
            inputData = append(inputData, appendCopy)
            currentLine = currentLine[:0]
        } else {
            currentLine = append(currentLine, char)
        }
    }

    return inputData
}

func padInput(input [][]rune) [][]rune {
    // by padding the input with empty spaces around the edges we can treat all nodes the same
    // without having to worry about walking off of the array
    inputLength := len(input)
    inputWidth := len(input[0])
    emptyLine := make([]rune, inputWidth + 6)
    var emptyRune rune
    paddedInput := make([][]rune, inputLength + 6)
    for row := range paddedInput {
        paddedInput[row] = make([]rune, inputWidth + 6)
    }
    for y := range inputLength + 5 {
        if y < 3 || y >= inputLength + 3 {
            paddedInput[y] = emptyLine
            continue
        }
        for x := range inputWidth + 5 {
            if x < 3 || x >= inputWidth + 3 {
                paddedInput[y][x] = emptyRune
                continue
            }
            paddedInput[y][x] = input[y - 3][x - 3]
        }
    }
    return paddedInput
}

func findXmasCountFromIndex(wordSearch [][]rune, x int, y int) int {
    xmasCount := 0
    
    // Bruce force check all directions from an X
    var words []string
    words = append(words, string([]rune{wordSearch[y][x], wordSearch[y-1][x], wordSearch[y-2][x], wordSearch[y-3][x]})) // up
    words = append(words, string([]rune{wordSearch[y][x], wordSearch[y+1][x], wordSearch[y+2][x], wordSearch[y+3][x]})) // down
    words = append(words, string([]rune{wordSearch[y][x], wordSearch[y][x-1], wordSearch[y][x-2], wordSearch[y][x-3]})) // left
    words = append(words, string([]rune{wordSearch[y][x], wordSearch[y][x+1], wordSearch[y][x+2], wordSearch[y][x+3]})) // right
    words = append(words, string([]rune{wordSearch[y][x], wordSearch[y-1][x-1], wordSearch[y-2][x-2], wordSearch[y-3][x-3]})) // up left
    words = append(words, string([]rune{wordSearch[y][x], wordSearch[y-1][x+1], wordSearch[y-2][x+2], wordSearch[y-3][x+3]})) // up right
    words = append(words, string([]rune{wordSearch[y][x], wordSearch[y+1][x-1], wordSearch[y+2][x-2], wordSearch[y+3][x-3]})) // down left
    words = append(words, string([]rune{wordSearch[y][x], wordSearch[y+1][x+1], wordSearch[y+2][x+2], wordSearch[y+3][x+3]})) // down right
    
    for _, word := range words {
        if word == "XMAS" {
            xmasCount += 1
        }
    }

    return xmasCount
}

func findXmasCount(wordSearch [][]rune) int {
    xmasCount := 0
    for y := range len(wordSearch) {
        for x := range len(wordSearch[0]) {
            if wordSearch[y][x] == 'X' {
                xmasCount += findXmasCountFromIndex(wordSearch, x, y)
            }
        }
    }
    return xmasCount
}

func isIndexCrossMas(wordSearch [][]rune, x int, y int) bool {
    crossA := string([]rune{wordSearch[y-1][x-1], wordSearch[y][x], wordSearch[y+1][x+1]})
    crossB := string([]rune{wordSearch[y-1][x+1], wordSearch[y][x], wordSearch[y+1][x-1]})
    if (crossA == "MAS" || crossA == "SAM") && 
       (crossB == "MAS" || crossB == "SAM") {
        return true
    }
    return false
}

func findCrossMasCount(wordSearch [][]rune) int {
    xmasCount := 0
    for y := range len(wordSearch) {
        for x := range len(wordSearch[0]) {
            if wordSearch[y][x] == 'A' {
                if isIndexCrossMas(wordSearch, x, y) {
                    xmasCount += 1
                }
            }
        }
    }
    return xmasCount
}

func main() {
    wordSearch := parseInput("./input.txt")
    paddedWordSearch := padInput(wordSearch)

    // Part1
    xmasCount := findXmasCount(paddedWordSearch)
    fmt.Printf("(Part 1) - XMAS Count: %d\n", xmasCount)

    // Part2
    crossMasCount := findCrossMasCount(paddedWordSearch)
    fmt.Printf("(Part 2) - X-MAS Count: %d\n", crossMasCount)
}