package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type button struct {
    moveX int
    moveY int
    cost int
}

type clawMachine struct {
    buttonA button
    buttonB button
    prizeX int
    prizeY int
}

func (cM clawMachine) findCheapestOption(maxPresses int) int {
    var possibleCosts []int
    for numAPresses := range maxPresses + 1 {
        remainingX := cM.prizeX - (cM.buttonA.moveX * numAPresses)
        if remainingX < 0 {break}
        if remainingX % cM.buttonB.moveX == 0 {
            numBPresses := remainingX/cM.buttonB.moveX
            if numBPresses > maxPresses {continue}
            if cM.prizeY - ((cM.buttonA.moveY * numAPresses) + (cM.buttonB.moveY * numBPresses)) == 0 {
                possibleCosts = append(possibleCosts, (cM.buttonA.cost * numAPresses) + (cM.buttonB.cost * numBPresses))
            }
        }
    }
    sort.Ints(possibleCosts)
    if len(possibleCosts) > 0 {
        return possibleCosts[0]
    }
    return 0
} 

func parseInput(path string) []clawMachine {
    var machines []clawMachine
    file, _ := os.Open(path)
    scanner := bufio.NewScanner(file)
    walker := 0
    aX, aY, bX, bY := 0, 0, 0, 0
    for scanner.Scan() {
        currentLine := scanner.Text()
        numberRegexp := regexp.MustCompile(`\d+`)
        if walker == 0 { // Button A
            numbers := numberRegexp.FindAllString(currentLine, -1)
            aX, _ = strconv.Atoi(numbers[0])
            aY, _ = strconv.Atoi(numbers[1])
        } else if walker == 1 { // Button B
            numbers := numberRegexp.FindAllString(currentLine, -1)
            bX, _ = strconv.Atoi(numbers[0])
            bY, _ = strconv.Atoi(numbers[1])
        } else if walker == 2 { // Prize
            numbers := numberRegexp.FindAllString(currentLine, -1)
            x, _ := strconv.Atoi(numbers[0])
            y, _ := strconv.Atoi(numbers[1])
            machines = append(machines, clawMachine{button{aX, aY, 3}, button{bX, bY, 1}, x, y})
        } else { // empty line
            walker = 0
            continue
        }
        walker += 1
    }
    file.Close()
    return machines
}

func main() {
    machines := parseInput("./input.txt")

    // Part 1
    total := 0
    maxPresses := 100
    for _, machine := range machines {
        total += machine.findCheapestOption(maxPresses)
    }
    fmt.Printf("(Part 1) - Total coin cost: %d\n", total)
}