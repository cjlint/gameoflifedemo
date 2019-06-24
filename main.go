package main

import (
	"bufio"
	"os"
	"time"
)

type point struct {
	x int
	y int
}

type grid struct {
	maxX, maxY int
	// "marks" is a grid that helps calculate how many live cells are adjacent to each space.
	// this method is simple but prevents us from simulating an "infinite" 2d plane
	marks [][]int
}

func initializeGrid(maxX, maxY int) *grid {
	marks := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		marks[i] = make([]int, maxY)
		for j := 0; j < maxY; j++ {
			marks[i][j] = 0
		}
	}
	return &grid{maxX, maxY, marks}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (g *grid) markSurrounding(p point) {
	if p.x < 0 || p.y < 0 || p.x >= g.maxX || p.y >= g.maxY {
		return
	}
	// min/max functions keep us within grid bounds
	for i := max(0, p.x-1); i < min(g.maxX, p.x+2); i++ {
		for j := max(0, p.y-1); j < min(g.maxY, p.y+2); j++ {
			if i == p.x && j == p.y {
				// live cells are not adjacent to themselves
				continue
			}
			g.marks[i][j]++
		}
	}
}

// go doesn't have a built-in "set" type so map[point]bool replaces that
func (g *grid) getNextLiveCells(liveCells map[point]bool) map[point]bool {
	// mark all the neighbors of live cells
	for cell := range liveCells {
		g.markSurrounding(cell)
	}
	// now look at every point in grid and apply these rules:
	// if >3 marks OR <2 marks, cell is dead in next generation
	// if =3 marks, cell is alive in next generation
	// if =2 marks, cell is alive in next generation IF it was also alive in this generation
	nextLiveCells := make(map[point]bool)
	for i := 0; i < g.maxX; i++ {
		for j := 0; j < g.maxY; j++ {
			currentPoint := point{i, j}
			if g.marks[i][j] == 3 || (g.marks[i][j] == 2 && liveCells[currentPoint]) {
				nextLiveCells[currentPoint] = true
			}
			// clean up marks for next time
			g.marks[i][j] = 0
		}
	}

	return nextLiveCells
}

func convertInputLineToPoints(y int, line string) []point {
	return []point{}
}

func main() {
	// TODO get these from command line
	maxX := 10
	maxY := 10
	interval := 500

	scanner := bufio.NewScanner(os.Stdin)
	liveCells := map[point]bool{}
	currentY := 0
	for scanner.Scan() {
		for _, cell := range convertInputLineToPoints(currentY, scanner.Text()) {
			liveCells[cell] = true
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	g := initializeGrid(maxX, maxY)
	for {
		// TODO print to screen
		liveCells = g.getNextLiveCells(liveCells)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
