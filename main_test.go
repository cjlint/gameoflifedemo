package main

import (
	"reflect"
	"testing"
)

func Test_initializeGrid(t *testing.T) {
	type args struct {
		maxX int
		maxY int
	}
	tests := []struct {
		name string
		args args
		want *grid
	}{
		{"empty grid doesn't fail", args{0, 0}, &grid{0, 0, [][]int{}}},
		{"non-empty grid initializes to all 0s", args{5, 5}, &grid{5, 5, [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}}},
		{"non-square grid is fine", args{3, 4}, &grid{3, 4, [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeGrid(tt.args.maxX, tt.args.maxY); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_markSurrounding(t *testing.T) {
	tests := []struct {
		name          string
		p             point
		expectedMarks [][]int
	}{
		{"mark surrounding 8", point{1, 1}, [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}},
		{"mark surrounding 5 from north edge", point{0, 1}, [][]int{
			{1, 0, 1},
			{1, 1, 1},
			{0, 0, 0},
		}},
		{"mark surrounding 5 from east edge", point{1, 2}, [][]int{
			{0, 1, 1},
			{0, 1, 0},
			{0, 1, 1},
		}},
		{"mark surrounding 5 from south edge", point{2, 1}, [][]int{
			{0, 0, 0},
			{1, 1, 1},
			{1, 0, 1},
		}},
		{"mark surrounding 5 from west edge", point{1, 0}, [][]int{
			{1, 1, 0},
			{0, 1, 0},
			{1, 1, 0},
		}},
		{"mark surrounding 3 from northeast corner", point{0, 2}, [][]int{
			{0, 1, 0},
			{0, 1, 1},
			{0, 0, 0},
		}},
		{"mark surrounding 3 from southeast corner", point{2, 2}, [][]int{
			{0, 0, 0},
			{0, 1, 1},
			{0, 1, 0},
		}},
		{"mark surrounding 3 from southwest corner", point{2, 0}, [][]int{
			{0, 0, 0},
			{1, 1, 0},
			{0, 1, 0},
		}},
		{"mark surrounding 3 from northwest corner", point{0, 0}, [][]int{
			{0, 1, 0},
			{1, 1, 0},
			{0, 0, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := initializeGrid(3, 3)
			g.markSurrounding(tt.p)
			if !reflect.DeepEqual(g.marks, tt.expectedMarks) {
				t.Errorf("grid.markSurrounding() = %v, want %v", g.marks, tt.expectedMarks)
			}
		})
	}
}

func Test_grid_markSurrounding_increments_existing_marks(t *testing.T) {
	g := initializeGrid(3, 3)
	g.markSurrounding(point{1, 1})
	g.markSurrounding(point{1, 1})
	g.markSurrounding(point{2, 2})
	expectedMarks := [][]int{
		{2, 2, 2},
		{2, 1, 3},
		{2, 3, 2},
	}
	if !reflect.DeepEqual(g.marks, expectedMarks) {
		t.Errorf("calling markSurrounding twice = %v, want %v", g.marks, expectedMarks)
	}
}

func Test_grid_getNextLiveCells(t *testing.T) {
	tests := []struct {
		name string
		seed map[point]bool
		want map[point]bool
	}{
		{"empty", map[point]bool{}, map[point]bool{}},
		{"full", map[point]bool{
			point{0, 0}: true,
			point{0, 1}: true,
			point{0, 2}: true,
			point{1, 0}: true,
			point{1, 1}: true,
			point{1, 2}: true,
			point{2, 0}: true,
			point{2, 1}: true,
			point{2, 2}: true,
		}, map[point]bool{
			point{0, 0}: true,
			point{0, 2}: true,
			point{2, 0}: true,
			point{2, 2}: true,
		}},
		{"line", map[point]bool{
			point{1, 0}: true,
			point{1, 1}: true,
			point{1, 2}: true,
		}, map[point]bool{
			point{0, 1}: true,
			point{1, 1}: true,
			point{2, 1}: true,
		}},
		{"plus sign", map[point]bool{
			point{1, 0}: true,
			point{1, 1}: true,
			point{1, 2}: true,
			point{0, 1}: true,
			point{2, 1}: true,
		}, map[point]bool{
			point{0, 0}: true,
			point{0, 1}: true,
			point{0, 2}: true,
			point{1, 0}: true,
			point{1, 2}: true,
			point{2, 0}: true,
			point{2, 1}: true,
			point{2, 2}: true,
		}},
		{"ring", map[point]bool{
			point{0, 0}: true,
			point{0, 1}: true,
			point{0, 2}: true,
			point{1, 0}: true,
			point{1, 2}: true,
			point{2, 0}: true,
			point{2, 1}: true,
			point{2, 2}: true,
		}, map[point]bool{
			point{0, 0}: true,
			point{0, 2}: true,
			point{2, 0}: true,
			point{2, 2}: true,
		}},
		{"out of bounds are ignored", map[point]bool{
			point{-1, 0}: true,
			point{-1, 1}: true,
			point{-1, 2}: true,
		}, map[point]bool{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := initializeGrid(3, 3)
			if got := g.getNextLiveCells(tt.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grid.getNextLiveCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
