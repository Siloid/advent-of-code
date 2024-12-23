package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput(path string) []*buyer {
    var buyers []*buyer
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        currentLine := scanner.Text()
        value, _ := strconv.Atoi(currentLine)
        buyers = append(buyers, &buyer{value, 0, make(map[string]int)})
    }
    return buyers
}

func mult64(secret int) int {
    return secret * 64
}

func mult2048(secret int) int {
    return secret * 2048
}

func mix(value int, secret int) int {
    return value ^ secret
}

func prune(secret int) int {
    return secret % 16777216
}

func div(secret int) int {
    return int(secret / 32)
}

func computeNextSecret(secret int) int {
    // step1
    result := mult64(secret)
    secret = mix(result, secret)
    secret = prune(secret)

    // step2
    result = div(secret)
    secret = mix(result, secret)
    secret = prune(secret)

    // step3
    result = mult2048(secret)
    secret = mix(result, secret)
    secret = prune(secret)

    return secret
}

type buyer struct {
    startingSecret int
    endingSecret int
    exchangeMap map[string]int
}

func getChangeKey(fourth int, third int, second int, first int) string {
    var changeKey string
    for _, value := range []int{fourth, third, second, first} {
        changeKey += strconv.Itoa(value)
    }
    return changeKey
}

func getMaxBananaSale(buyers []*buyer) int {
    bestTotal := 0
    for i := -9; i < 10; i++ {
        for j := -9; j < 10; j++ {
            for k := -9; k < 10; k++ {
                for l := -9; l < 10; l++ {
                    changeKey := getChangeKey(i,j,k,l)
                    total := 0
                    for _, buyer := range buyers {
                        value, inMap := buyer.exchangeMap[changeKey]
                        if inMap {
                            total += value
                        }
                    }
                    if total > bestTotal {
                        bestTotal = total
                    }
                }
            }
        }
    }
    return bestTotal
}


func main() {
    buyers := parseInput("./input.txt")
    total := 0
    for _, buyer := range buyers {
        secret := buyer.startingSecret
        fourth, third, second, first, previousLast := 0, 0, 0, 0, buyer.startingSecret % 10
        for i := range 2000 {
            secret = computeNextSecret(secret)
            lastDigit := secret % 10
            fourth = third
            third = second
            second = first
            first = lastDigit - previousLast
            if i >= 3 { 
                changeKey := getChangeKey(fourth, third, second, first)
                _, inMap := buyer.exchangeMap[changeKey]
                if !inMap {
                    buyer.exchangeMap[changeKey] = lastDigit
                }
            }
            previousLast = lastDigit
        }
        buyer.endingSecret = secret
    }
    for _, buyer := range buyers {
        total += buyer.endingSecret
    }
    fmt.Printf("(Part 1) - Total secret values after 2000 iterations: %d\n", total)
    fmt.Printf("(Part 2) - Total bananas: %d\n", getMaxBananaSale(buyers))

}