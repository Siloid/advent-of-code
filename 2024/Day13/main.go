package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func (cM *clawMachine) adjustPrize(adjustment int) {
    cM.prizeX += adjustment
    cM.prizeY += adjustment
}

func (cM clawMachine) findCheapestOption() int {
    /*
    * "Cramer's Rule" - Formula for solving this, look it up
    * I'm not 100% convinced this works in all cases
    * As I understand it, it returns the minimum number of presses which is not necessarily the cheapest, but it seemed to work for every case I tried /shrug
    */
    pressesA := int(((cM.prizeX * cM.buttonB.moveY) - (cM.buttonB.moveX * cM.prizeY))/((cM.buttonA.moveX * cM.buttonB.moveY) - (cM.buttonB.moveX * cM.buttonA.moveY)))
    pressesB := int(((cM.buttonA.moveX * cM.prizeY) - (cM.prizeX * cM.buttonA.moveY))/((cM.buttonA.moveX * cM.buttonB.moveY) - (cM.buttonB.moveX * cM.buttonA.moveY)))
    if ((cM.buttonA.moveX * pressesA) + (cM.buttonB.moveX * pressesB)) == cM.prizeX &&
       ((cM.buttonA.moveY * pressesA) + (cM.buttonB.moveY * pressesB)) == cM.prizeY {
        return (pressesA * cM.buttonA.cost) + (pressesB * cM.buttonB.cost)
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
    for _, machine := range machines {
        total += machine.findCheapestOption()
    }
    fmt.Printf("(Part 1) - Total coin cost: %d\n", total)

    // Part 2
    total = 0
    for _, machine := range machines {
        machine.adjustPrize(10000000000000)
        total += machine.findCheapestOption()
    }
    fmt.Printf("(Part 2) - Total coin cost: %d\n", total)
}