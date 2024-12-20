package warehouse

import (
    "bufio"
    "os"
    "strings"
)

const (
    up    = '^'
    down  = 'v'
    left  = '<'
    right = '>'
)

type robot struct {
    x            int
    y            int
    instructions []rune
}

type warehouse struct {
    worker robot
    layout [][]rune
}

func getStartingPosition(layout [][]rune) (int, int) {
    for y := range len(layout) {
        for x := range len(layout[0]) {
            if layout[y][x] == '@' {
                layout[y][x] = '.'
                return x, y
            }
        }
    }
    return 0, 0 // should never occur
}

func ParseInput(path string) warehouse {
    var layout [][]rune
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        currentLine := scanner.Text()
        if strings.TrimSpace(currentLine) == "" {
            break
        }
        layout = append(layout, []rune(strings.TrimRight(currentLine, "\n")))
    }
    startX, startY := getStartingPosition(layout)

    var instructions []rune
    for scanner.Scan() {
        currentLine := scanner.Text()
        instructions = append(instructions, []rune(strings.TrimRight(currentLine, "\n"))...)
    }
    return warehouse{robot{startX, startY, instructions}, layout}
}

func (r *robot) moveUp(layout [][]rune) {
    space := layout[r.y-1][r.x]
    if space == '.' {
        r.y -= 1
    } else if space == 'O' {
        // attempt to shove boxes
        for y := r.y - 2; y >= 0; y-- {
            if layout[y][r.x] == '.' { // shove boxes
                layout[r.y-1][r.x] = '.'
                layout[y][r.x] = 'O'
                r.y -= 1
                return
            } else if layout[y][r.x] == '#' { // hit wall, can't shove
                return
            }
        }
    }
}

func (r *robot) moveDown(layout [][]rune) {
    space := layout[r.y+1][r.x]
    if space == '.' {
        r.y += 1
    } else if space == 'O' {
        // attempt to shove boxes
        for y := r.y + 2; y < len(layout); y++ {
            if layout[y][r.x] == '.' { // shove boxes
                layout[r.y+1][r.x] = '.'
                layout[y][r.x] = 'O'
                r.y += 1
                return
            } else if layout[y][r.x] == '#' { // hit wall, can't shove
                return
            }
        }
    }
}

func (r *robot) moveLeft(layout [][]rune) {
    space := layout[r.y][r.x-1]
    if space == '.' {
        r.x -= 1
    } else if space == 'O' {
        // attempt to shove boxes
        for x := r.x - 2; x >= 0; x-- {
            if layout[r.y][x] == '.' { // shove boxes
                layout[r.y][r.x-1] = '.'
                layout[r.y][x] = 'O'
                r.x -= 1
                return
            } else if layout[r.y][x] == '#' { // hit wall, can't shove
                return
            }
        }
    }
}

func (r *robot) moveRight(layout [][]rune) {
    space := layout[r.y][r.x+1]
    if space == '.' {
        r.x += 1
    } else if space == 'O' {
        // attempt to shove boxes
        for x := r.x + 2; x < len(layout[0]); x++ {
            if layout[r.y][x] == '.' { // shove boxes
                layout[r.y][r.x+1] = '.'
                layout[r.y][x] = 'O'
                r.x += 1
                return
            } else if layout[r.y][x] == '#' { // hit wall, can't shove
                return
            }
        }
    }
}

func (wh *warehouse) RunSimulation() {
    for _, step := range wh.worker.instructions {
        if step == up {
            wh.worker.moveUp(wh.layout)
        } else if step == down {
            wh.worker.moveDown(wh.layout)
        } else if step == left {
            wh.worker.moveLeft(wh.layout)
        } else if step == right {
            wh.worker.moveRight(wh.layout)
        }
    }
}

func (wh warehouse) GetBoxGPSSum() int {
    gpsSum := 0
    for y := range len(wh.layout) {
        for x := range len(wh.layout[0]) {
            if wh.layout[y][x] == 'O' {
                gpsSum += (y * 100) + x
            }
        }
    }
    return gpsSum
}
