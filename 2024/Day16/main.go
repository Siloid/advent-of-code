package main

import (
	"bufio"
	"fmt"
	"maze"
	"os"
	"strings"
)

func parseInput(path string) [][]rune {
	var layout [][]rune
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		layout = append(layout, []rune(strings.TrimRight(line, "\n")))
	}
	return layout
}

func main() {
	inputData := parseInput("input.txt")

	// Part1
	myMaze := maze.NewMaze(inputData)
	fmt.Printf("(Part 1) - Cheapest path cost: %d\n", myMaze.GetCheapestPath())
}
