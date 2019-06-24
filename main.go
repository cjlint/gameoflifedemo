package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	tm "github.com/buger/goterm"
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
	// line should be a series of 0s and 1s, like "00010011010"
	result := make([]point, len(line))
	for i, c := range line {
		if c == '1' {
			result = append(result, point{i, y})
		}
	}
	return result
}

func main() {
	maxXPtr := flag.Int("maxX", 10, "grid X boundary")
	maxYPtr := flag.Int("maxY", 10, "grid Y boundary")
	intervalPtr := flag.Int("interval", 500, "time between game ticks (in milliseconds)")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	liveCells := map[point]bool{}
	currentY := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			fmt.Println("breaking")
			break
		}
		for _, cell := range convertInputLineToPoints(currentY, line) {
			liveCells[cell] = true
		}
		currentY++
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Print("liveCells", liveCells)

	g := initializeGrid(*maxXPtr, *maxYPtr)
	for {
		fmt.Println("main loop")
		tm.Clear()
		tm.MoveCursor(1, 1)
		// flip X and Y to make it appear normal on command line
		for i := 0; i < g.maxY; i++ {
			for j := 0; j < g.maxX; j++ {
				if liveCells[point{j, i}] {
					tm.Print(tm.Color("O", tm.GREEN))
				} else {
					tm.Print(tm.Color("O", tm.RED))
				}
			}
			tm.Println()
		}
		tm.Flush()
		liveCells = g.getNextLiveCells(liveCells)
		time.Sleep(time.Duration(*intervalPtr) * time.Millisecond)
	}
}
