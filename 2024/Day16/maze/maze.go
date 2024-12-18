package maze

import (
	"strconv"
)

const (
	left  = 0
	right = 1
	up    = 2
	down  = 3
)

type room struct {
	x         int
	y         int
	cost      int // cheapest cost to reach this room
	neighbors []*room
}

func (r *room) connect(neighbor *room) {
	r.neighbors = append(r.neighbors, neighbor)
}

func (r *room) walk(cost int, fromRoom *room) {
	if r.cost < cost && r.cost != -1 {
		return
	}

	r.cost = cost
	for _, nextRoom := range r.neighbors {
		if nextRoom == fromRoom {
			continue
		}
		newCost := r.cost + calculateCost(fromRoom, r, nextRoom)
		nextRoom.walk(newCost, r)
	}
}

func getRoomKey(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func getDirection(fromRoom *room, toRoom *room) int {
	if fromRoom.x == toRoom.x {
		if fromRoom.y < toRoom.y {
			return down
		}
		return up
	}
	if fromRoom.x < toRoom.x {
		return right
	}
	return left
}

func calculateCost(fromRoom *room, currentRoom *room, nextRoom *room) int {
	var currentDirection int
	if fromRoom == nil { // starting room
		currentDirection = right
	} else {
		currentDirection = getDirection(fromRoom, currentRoom)
	}
	newDirection := getDirection(currentRoom, nextRoom)
	if newDirection == currentDirection {
		return 1 // continue in current direction, 1 step
	}
	return 1001 // 1000 + 1, one turn plus a step
}

type Maze struct {
	layout map[string]*room
	start  string
	end    string
}

func (m *Maze) connectRooms(roomKey1 string, roomKey2 string) {
	m.layout[roomKey1].connect(m.layout[roomKey2])
	m.layout[roomKey2].connect(m.layout[roomKey1])
}

func (m *Maze) computeCosts() {
	m.layout[m.start].walk(0, nil)
}

func (m *Maze) GetCheapestPath() int {
	return m.layout[m.end].cost
}

func NewMaze(layout [][]rune) Maze {
	var newMaze Maze
	newMaze.layout = make(map[string]*room)
	for y, row := range layout {
		for x, char := range row {
			if char == '#' {
				continue
			} else if char == 'S' {
				newMaze.start = getRoomKey(x, y)
			} else if char == 'E' {
				newMaze.end = getRoomKey(x, y)
			}
			newRoom := room{x, y, -1, []*room{}}
			newMaze.layout[getRoomKey(x, y)] = &newRoom
			if layout[y][x-1] == '.' || layout[y][x-1] == 'S' || layout[y][x-1] == 'E' {
				newMaze.connectRooms(getRoomKey(x, y), getRoomKey(x-1, y))
			}
			if layout[y-1][x] == '.' || layout[y-1][x] == 'S' || layout[y-1][x] == 'E' {
				newMaze.connectRooms(getRoomKey(x, y), getRoomKey(x, y-1))
			}
		}
	}
	newMaze.computeCosts()
	return newMaze
}
