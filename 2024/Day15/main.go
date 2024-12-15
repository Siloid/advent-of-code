package main

import (
    "fmt"
    "warehouse"
    "widewarehouse"
)

// run with "go run ." not "go run main.go"
func main() {
    // Part1
    wh := warehouse.ParseInput("./input.txt")
    wh.RunSimulation()
    fmt.Printf("(Part 1) - Box GPS Sum: %d\n", wh.GetBoxGPSSum())

    // Part2
    wWh := widewarehouse.ParseInput("./input.txt")
    wWh.RunSimulation()
    fmt.Printf("(Part 2) - Box GPS Sum Wide Warehouse: %d\n", wWh.GetBoxGPSSum())
}
