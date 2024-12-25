package day06

import (
	"bufio"
	"log"
	"os"
)

func solvePart2(path string) int {
	m, pos, dir := parseMap(path)

	visited := run(m, pos, dir)

	var count int
	for obstacle := range visited {
		if testObstacle(m, obstacle, pos, dir) {
			count++
		}
	}
	return count
}

type status struct {
	Pos P
	Dir Direction
}

func testObstacle(m map[P]struct{}, obstacle P, pos P, dir Direction) bool {

	// impossible if the guard is there
	if obstacle == pos {
		return false
	}

	// impossible if there is already an obstacle there
	if _, ok := m[obstacle]; ok {
		return false
	}

	m[obstacle] = struct{}{}
	defer func() {
		delete(m, obstacle)
	}()

	visited := map[status]struct{}{
		status{pos, dir}: {},
	}

	for {
		next := pos.Next(dir)

		if next.x < 0 || next.x >= Size || next.y < 0 || next.y >= Size {
			return false
		}

		if _, ok := m[next]; ok {
			dir = (dir + 1) % 4
		} else {
			pos = next
		}

		st := status{pos, dir}
		if _, ok := visited[st]; ok {
			return true
		}
		visited[st] = struct{}{}
	}
}

const Size = 130

type Direction uint8

const (
	N Direction = iota
	E
	S
	W
)

type P struct{ x, y int }

func (p P) Next(d Direction) P {
	switch d {
	case N:
		return P{p.x - 1, p.y}
	case E:
		return P{p.x, p.y + 1}
	case S:
		return P{p.x + 1, p.y}
	case W:
		return P{p.x, p.y - 1}
	}
	return P{0, 0}
}

func parseMap(path string) (map[P]struct{}, P, Direction) {
	var pos P
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	m := map[P]struct{}{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			switch c {
			case '#':
				m[P{i, j}] = struct{}{}
			case '^':
				pos = P{i, j}
			}
		}
	}
	return m, pos, N
}

func run(m map[P]struct{}, pos P, dir Direction) map[P]struct{} {
	visited := map[P]struct{}{
		pos: {},
	}

	for {
		next := pos.Next(dir)

		if next.x < 0 || next.x >= Size || next.y < 0 || next.y >= Size {
			return visited
		}

		if _, ok := m[next]; ok {
			dir = (dir + 1) % 4
		} else {
			pos = next
			visited[pos] = struct{}{}
		}
	}
}
