package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
    adv = 0
    bxl = 1
    bst = 2
    jnz = 3
    bxc = 4
    out = 5
    bdv = 6
    cdv = 7
)

type device struct {
    A int
    B int
    C int
    program []int
    instructionIndex int
}

func (d *device) adv(operand int) {
    // opcode 0
    d.A = int(d.A / int(math.Pow(2, float64(d.getComboOperand(operand)))))
}

func (d *device) bxl(operand int) {
    // opcode 1
    d.B = d.B ^ operand
}

func (d *device) bst(operand int) {
    // opcode 2
    d.B = d.getComboOperand(operand) % 8
}

func (d *device) jnz(operand int) bool {
    // opcode 3
    if d.A == 0 {
        return false
    }
    d.instructionIndex = operand
    return true
}

func (d *device) bxc(_ int) {
    // opcode 4
    d.B = d.B ^ d.C
}

func (d device) out(operand int) int {
    // opcode 5
    return d.getComboOperand(operand) % 8
}

func (d *device) bdv(operand int) {
    // opcode 6
    d.B = int(d.A / int(math.Pow(2, float64(d.getComboOperand(operand)))))
}

func (d *device) cdv(operand int) {
    // opcode 7
    d.C = int(d.A / int(math.Pow(2, float64(d.getComboOperand(operand)))))
}

func (d device) getComboOperand(operand int) int {
    switch operand {
    case 4:
        return d.A
    case 5:
        return d.B
    case 6:
        return d.C
    // case 7: invalid
    default: // 0, 1, 2, 3
        return operand
    }
}

func (d *device) runIt() []int {
    var output []int
    for d.instructionIndex < len(d.program) {
        opcode := d.program[d.instructionIndex]
        operand := d.program[d.instructionIndex + 1]
        indexMoved := false
        switch opcode {
        case adv:
            d.adv(operand)
        case bxl:
            d.bxl(operand)
        case bst:
            d.bst(operand)
        case jnz:
            indexMoved = d.jnz(operand)
        case bxc:
            d.bxc(operand)
        case out:
            output = append(output, d.out(operand))
        case bdv:
            d.bdv(operand)
        case cdv:
            d.cdv(operand)
        default:
            fmt.Println("Error: Invalid opcode: ", opcode)
        }
        if !indexMoved {
            d.instructionIndex += 2
        }

    }
    return output
}

func (d device) getOutputString() string {
    output := d.runIt()
    var outputStr string
    for i, num := range output {
        if i == 0 {
            outputStr = strconv.Itoa(num)
        } else {
            outputStr = outputStr + "," + strconv.Itoa(num)
        }
    }
    return outputStr
}

func (d device) getPossibleA(nextNumber int, possibleStartingA []int) []int {
    var nextPossibleStartingA []int
    for _, possibleA := range possibleStartingA {
        for i := range 8 {
            newA := (possibleA << 3) + i
            newDevice := device{newA, 0, 0, d.program, 0}
            result := newDevice.runIt()
            if result[0] == nextNumber {
                nextPossibleStartingA = append(nextPossibleStartingA, newA)
            }
        }
    }
    return nextPossibleStartingA
}

func (d device) findDuplicateA() int {
    possibleStartingA := []int{0, 1, 2, 3, 4, 5, 6, 7}
    for i := len(d.program) - 1; i >= 0; i-- {
        possibleStartingA = d.getPossibleA(d.program[i], possibleStartingA)
    }
    return possibleStartingA[0]
}

func parseInput(path string) device {
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    lineCount := 0
    A, B, C := 0,0,0
    var program []int
    for scanner.Scan() {
        currentLine := scanner.Text()
        lineCount += 1
        switch lineCount {
        case 1:
            valueStr := strings.Fields(strings.TrimSpace(currentLine))[2]
            A, _ = strconv.Atoi(valueStr)
        case 2:
            valueStr := strings.Fields(strings.TrimSpace(currentLine))[2]
            B, _ = strconv.Atoi(valueStr)
        case 3:
            valueStr := strings.Fields(strings.TrimSpace(currentLine))[2]
            C, _ = strconv.Atoi(valueStr)
        case 5:
            programStr := strings.Fields(strings.TrimSpace(currentLine))[1]
            valueStrs := strings.Split(programStr, ",")
            for _, valueStr := range valueStrs {
                value, _ := strconv.Atoi(valueStr)
                program = append(program, value)
            }
        default: // 4 empty line
            continue 
        }
    }
    return device{A, B, C, program, 0}
}

func main() {
    myDevice := parseInput("input.txt")
    
    // Part1
    fmt.Printf("(Part 1) - Device output: %s\n", myDevice.getOutputString())

    // Part2
    fmt.Printf("(Part 2) - Value of A which outputs copy: %d\n", myDevice.findDuplicateA())
}
