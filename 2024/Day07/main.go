package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(path string) []string {
    file, _ := os.Open(path)
    scanner := bufio.NewScanner(file)
    var equationStrings []string
    for scanner.Scan() {
        equationStrings = append(equationStrings, scanner.Text())
    }
    file.Close()
    return equationStrings
}

func existsPossible(target int, currentValue int, operator rune, numbers []int, enableConcatenation bool) bool {
    nextNumber, numbers := numbers[0], numbers[1:]
    if operator == '+' {
        currentValue += nextNumber
    } else if operator == '*' {
        currentValue *= nextNumber
    } else {
        if !enableConcatenation {
            return false
        }
        currentValueString := strconv.Itoa(currentValue) + strconv.Itoa(nextNumber)
        currentValue, _ = strconv.Atoi(currentValueString)
    }
    if len(numbers) == 0 {
        if currentValue == target {
            return true
        }
        return false
    }

    return existsPossible(target, currentValue, '+', numbers, enableConcatenation) ||
           existsPossible(target, currentValue, '*', numbers, enableConcatenation) ||
           existsPossible(target, currentValue, '|', numbers, enableConcatenation)
}

func isValid(target int, numbers []int, enableConcatenation bool) bool {
    value, numbers := numbers[0], numbers[1:]

    return existsPossible(target, value, '+', numbers, enableConcatenation) ||
           existsPossible(target, value, '*', numbers, enableConcatenation) ||
           existsPossible(target, value, '|', numbers, enableConcatenation)
}

func parseEquation(equation string) (int, []int) {
    parts := strings.Split(equation, ":")
    target, _ := strconv.Atoi(parts[0])
    numberStrings := strings.Fields(parts[1])
    var numbers []int
    for _, numberString := range numberStrings {
        number, _ := strconv.Atoi(numberString)
        numbers = append(numbers, number)
    }
    return target, numbers
}

func sumValidEquations(equations []string, enableConcatenation bool) int {
    validEquationSum := 0
    for _, equation := range equations {
        target, numbers := parseEquation(equation)
        if isValid(target, numbers, enableConcatenation) {
            validEquationSum += target
        }
    }
    return validEquationSum
}

func main() {
    equations := parseInput("./input.txt")
    
    // Part1
    total := sumValidEquations(equations, false)
    fmt.Printf("(Part 1) - valid equation sum: %d\n", total)

    // Part2
    total = sumValidEquations(equations, true)
    fmt.Printf("(Part 2) - valid equation sum w/ concatenation: %d\n", total)
}