package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

type shard struct {
    isFile bool
    fileIndex int
    size int
    hasBeenMoved bool
}

// Common
func parseInput(path string) []int {
    var data []int
    file, _ := os.Open(path)
    reader := bufio.NewReader(file)
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            break
        }
        value, _ := strconv.Atoi(string(char))
        data = append(data, value)
    }
    file.Close()
    return data
}

func computeChecksum(data []int) int {
    total := 0
    for i, value := range data {
        if value == -1 {continue}
        total += (i * value)
    }
    return total
}

// Part1 functions
func uncompress(compressedData []int) []int {
    var uncompressedData []int
    for i, space := range compressedData {
        var toAdd int
        // even indices are files, odd are free space
        if i % 2 == 0 { // file
            toAdd = (i + 1) / 2 // file index 
        } else { // free space
            toAdd = -1 // '.' is represented as -1
        }
        for range space {
            uncompressedData = append(uncompressedData, toAdd)
        }
    }
    return uncompressedData
}

func defragment(data []int) []int {
    reverseIndex := len(data)
    for i, value := range data {
        if value != -1 {
            continue
        }
        reverseIndex -= 1
        for i < reverseIndex {
            if data[reverseIndex] != -1 {
                data[i], data[reverseIndex] = data[reverseIndex], data[i]
                break
            }
            reverseIndex -= 1
        }
    }
    return data
}

// Part2 functions
func convertToShards(compressedData []int) []shard {
    var shards []shard
    for i, size := range compressedData {
        if size == 0 {continue}
        // even indices are files, odd are free space
        if i % 2 == 0 { // file
            shards = append(shards, shard{true, (i + 1) / 2, size, false})
        } else { // free space
            shards = append(shards, shard{false, -1, size, false})
        }
    }
    return shards
}

func moveShard(shards []shard) (bool, []shard) {
    movedShard := false
    for i := len(shards) - 1; i >= 0; i-- {
        shardToMove := shards[i]
        if !shardToMove.isFile || shardToMove.hasBeenMoved {continue}
        for j, memShard := range shards {
            if j >= i {break}
            if !memShard.isFile && shardToMove.size <= memShard.size {
                // move file
                shardToMove.hasBeenMoved = true
                if shardToMove.size == memShard.size {
                    //swap
                    shards[i] = memShard
                    shards[j] = shardToMove
                } else {
                    // split memory and move
                    shards[i] = shard{false, -1, shardToMove.size, false}
                    shards[j] = shardToMove
                    shards = slices.Insert(shards, j+1, shard{false, -1, memShard.size - shardToMove.size, false})
                }                
                movedShard = true
                break
            }
        }
        if movedShard {break}
    }
    return movedShard, shards
}

func orderShards(shards []shard) []shard {
    movedShard, updatedShards := moveShard(shards)
    for movedShard {
        movedShard, updatedShards = moveShard(updatedShards)
    }
    return updatedShards
}

func convertShardsToIntSlice(shards []shard) []int {
    var intSlice []int
    for _, shard := range shards {
        for range shard.size {
            intSlice = append(intSlice, shard.fileIndex)
        }
    }
    return intSlice
}

// main
func main() {
    compressedData := parseInput("input.txt")

    // Part1
    uncompressedData := uncompress(compressedData)
    defragmentedData := defragment(uncompressedData)
    checksum := computeChecksum(defragmentedData)
    fmt.Printf("(Part 1) - checksum: %d\n", checksum)

    // Part2
    shards := convertToShards(compressedData)
    orderedShards := orderShards(shards)
    intSlice := convertShardsToIntSlice(orderedShards)
    checksum = computeChecksum(intSlice)
    fmt.Printf("(Part 2) - checksum w/ whole file defragmenting: %d\n", checksum)
}