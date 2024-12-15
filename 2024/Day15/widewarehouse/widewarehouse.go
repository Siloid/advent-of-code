package widewarehouse

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
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

type widewarehouse struct {
    worker robot
    layout [][]rune
}

func ParseInput(path string) widewarehouse {
    var layout [][]rune
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    x, y, startX, startY := 0, 0, 0, 0
    for scanner.Scan() {
        currentLine := scanner.Text()
        if strings.TrimSpace(currentLine) == "" {
            break
        }
        x = 0
        var newLine string
        for _, char := range currentLine {
            if char == '#' {
                newLine = newLine + "##"
            } else if char == '.' {
                newLine = newLine + ".."
            } else if char == 'O' {
                newLine = newLine + "[]"
            } else if char == '@' {
                newLine = newLine + ".."
                startX, startY = x, y
            }
            x += 2
        }
        layout = append(layout, []rune(newLine))
        y += 1
    }

    var instructions []rune
    for scanner.Scan() {
        currentLine := scanner.Text()
        instructions = append(instructions, []rune(strings.TrimRight(currentLine, "\n"))...)
    }
    return widewarehouse{robot{startX, startY, instructions}, layout}
}

func canPush(layout [][]rune, fromX int, fromY int, direction rune) bool {
    yMod := 1
    if direction == up {
        yMod = -1
    }
    // always push in relation to left side of box
    if layout[fromY+yMod][fromX] == '#' || layout[fromY+yMod][fromX+1] == '#' {
        return false
    } else if layout[fromY+yMod][fromX] == '.' && layout[fromY+yMod][fromX+1] == '.' {
        return true
    } else if layout[fromY+yMod][fromX] == '[' {
        return canPush(layout, fromX, fromY+yMod, direction)
    } else if layout[fromY+yMod][fromX] == '.' && layout[fromY+yMod][fromX+1] == '[' {
        return canPush(layout, fromX+1, fromY+yMod, direction)
    } else if layout[fromY+yMod][fromX] == ']' && layout[fromY+yMod][fromX+1] == '.' {
        return canPush(layout, fromX-1, fromY+yMod, direction)
    } else if layout[fromY+yMod][fromX] == ']' && layout[fromY+yMod][fromX+1] == '[' {
        return canPush(layout, fromX-1, fromY+yMod, direction) && canPush(layout, fromX+1, fromY+yMod, direction)
    }
    return false // this shouldn't happen
}

func recursePush(layout [][]rune, fromX int, fromY int, direction rune, initialPush bool) {
    yMod := 1
    if direction == up {
        yMod = -1
    }
    // always pushing in relation to left side of box
    if layout[fromY+yMod][fromX] == '[' {
        recursePush(layout, fromX, fromY+yMod, direction, false)
    } else if layout[fromY+yMod][fromX] == '.' && layout[fromY+yMod][fromX+1] == '[' {
        recursePush(layout, fromX+1, fromY+yMod, direction, false)
    } else if layout[fromY+yMod][fromX] == ']' && layout[fromY+yMod][fromX+1] == '.' {
        recursePush(layout, fromX-1, fromY+yMod, direction, false)
    } else if layout[fromY+yMod][fromX] == ']' || layout[fromY+yMod][fromX+1] == '[' {
        recursePush(layout, fromX+1, fromY+yMod, direction, false)
        recursePush(layout, fromX-1, fromY+yMod, direction, false)
    }
    if !initialPush {
        layout[fromY+yMod][fromX], layout[fromY+yMod][fromX+1] = '[', ']'
        layout[fromY][fromX], layout[fromY][fromX+1] = '.', '.'
    } else {
        layout[fromY+yMod][fromX], layout[fromY+yMod][fromX+1] = '.', '.'
    }

}

func (r *robot) moveUp(layout [][]rune) {
    space := layout[r.y-1][r.x]
    if space == '.' {
        r.y -= 1
    } else if space == '[' {
        if canPush(layout, r.x, r.y, up) {
            recursePush(layout, r.x, r.y, up, true)
            r.y -= 1
        }
    } else if space == ']' {
        if canPush(layout, r.x-1, r.y, up) {
            recursePush(layout, r.x-1, r.y, up, true)
            r.y -= 1
        }
    }
}

func (r *robot) moveDown(layout [][]rune) {
    space := layout[r.y+1][r.x]
    if space == '.' {
        r.y += 1
    } else if space == '[' {
        if canPush(layout, r.x, r.y, down) {
            recursePush(layout, r.x, r.y, down, true)
            r.y += 1
        }
    } else if space == ']' {
        if canPush(layout, r.x-1, r.y, down) {
            recursePush(layout, r.x-1, r.y, down, true)
            r.y += 1
        }
    }
}

func (r *robot) moveLeft(layout [][]rune) {
    space := layout[r.y][r.x-1]
    if space == '.' {
        r.x -= 1
    } else if space == ']' {
        for x := r.x - 1; x > 0; x -= 2 {
            if layout[r.y][x] == '#' {
                return
            } else if layout[r.y][x] == '.' {
                for moveX := x; moveX < r.x; moveX++ {
                    layout[r.y][moveX] = layout[r.y][moveX+1]
                }
                r.x -= 1
                return
            }
        }
    }
}

func (r *robot) moveRight(layout [][]rune) {
    space := layout[r.y][r.x+1]
    if space == '.' {
        r.x += 1
    } else if space == '[' {
        for x := r.x + 1; x < len(layout[0]); x += 2 {
            if layout[r.y][x] == '#' {
                return
            } else if layout[r.y][x] == '.' {
                for moveX := x; moveX > r.x; moveX-- {
                    layout[r.y][moveX] = layout[r.y][moveX-1]
                }
                r.x += 1
                return
            }
        }
    }
}

func (wh widewarehouse) PrettyPrint() {
    for _, line := range wh.layout {
        fmt.Println(string(line))
    }
    fmt.Println("Robot Location: " + strconv.Itoa(wh.worker.x) + "," + strconv.Itoa(wh.worker.y))
}

func (wh *widewarehouse) RunSimulation() {
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

func (wh widewarehouse) GetBoxGPSSum() int {
    gpsSum := 0
    for y := range len(wh.layout) {
        for x := range len(wh.layout[0]) {
            if wh.layout[y][x] == '[' {
                gpsSum += (y * 100) + x
            }
        }
    }
    return gpsSum
}
