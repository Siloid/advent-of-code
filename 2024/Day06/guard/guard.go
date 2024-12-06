package guard

import (
	"slices"
	"strconv"
)

var DirectionMap = map[int][]int{
    0: []int{0, -1}, //up
    1: []int{1, 0}, //right
    2: []int{0, 1}, //down
    3: []int{-1, 0}, //left
}

type guard struct {
    floorplan [][]rune
    direction int
    currentX int
    currentY int
    visitedLocations []string
    visitedLocationsWithDirection []string
}

func NewGuard(floorplan [][]rune, startingX int, startingY int) *guard {
    newGuard := guard{floorplan: floorplan, currentX: startingX, currentY: startingY}
    newGuard.direction = 0 //Always starts facing up
    newGuard.visitedLocations = append(newGuard.visitedLocations, strconv.Itoa(startingX) + "," + strconv.Itoa(startingY))
    return &newGuard
}

func (g *guard) move() int {
    // returns 0 if the guard could move
    // returns 1 if not, i.e. edge of map
    // returns -1 if caught walking in a loop
    nextX := g.currentX + DirectionMap[g.direction][0]
    nextY := g.currentY + DirectionMap[g.direction][1]
    
    // Check for edge of map
    if nextX < 0 || nextX >= len(g.floorplan[0]) || nextY < 0 || nextY >= len(g.floorplan) {
        return 1
    }

    room := g.floorplan[nextY][nextX]
    switch room {
    case '#':
        g.turn()
    default:
        g.currentX = nextX
        g.currentY = nextY
        locationString := strconv.Itoa(g.currentX) + "," + strconv.Itoa(g.currentY)
        directionalLocationString := strconv.Itoa(g.currentX) + "," + strconv.Itoa(g.currentY) + "," + strconv.Itoa(g.direction)
        if !slices.Contains(g.visitedLocations, locationString) {
            g.visitedLocations = append(g.visitedLocations, locationString)
            g.visitedLocationsWithDirection = append(g.visitedLocationsWithDirection, directionalLocationString)
        } else {
            if slices.Contains(g.visitedLocationsWithDirection, directionalLocationString) {
                // We've walked this path before, infinite loop
                return -1
            } else {
                g.visitedLocationsWithDirection = append(g.visitedLocationsWithDirection, directionalLocationString)
            }
        }
    }
    return 0
}

func (g *guard) turn() {
    g.direction = (g.direction + 1) % 4
}

func (g *guard) Patrol() int {
    // returns the number of unique rooms visited or -1 if stuck in a loop
    patrolling := true
    for patrolling {
        moveResult := g.move()
        switch moveResult {
        case 1:
            patrolling = false 
        case -1:
            return -1
        default:
            continue
        }
    }
    return len(g.visitedLocations)
}

