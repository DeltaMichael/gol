package main

import (
	"fmt"
)

type Cell struct {
	y int
	x int
}

func printGrid(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%d, ", grid[i][j])
		}
		fmt.Println()
	}
}

func createGrid(y int, x int) [][]int {
	grid := make([][]int, y)
	for i := range grid {
		grid[i] = make([]int, x)
	}
	return grid
}

func getCell(y, x int, grid[][]int) int {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0]) {
		return grid[y][x]
	}
	return 0
}

func isCellValid(y, x int, grid[][]int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}

func getLiveNeighbors(y int, x int, grid [][]int) int {
	return getCell(y - 1, x, grid) + getCell(y + 1, x, grid) + getCell(y, x - 1, grid) + getCell(y, x + 1, grid) + getCell(y + 1, x + 1, grid) + getCell(y + 1, x - 1, grid) + getCell(y - 1, x + 1, grid) + getCell(y - 1, x - 1, grid)
}

func draw(cells []Cell, grid [][]int) {
	for _, cell := range cells {
		if isCellValid(cell.y, cell.x, grid) {
			grid[cell.y][cell.x] = 1
		}
	}
}

func remove(cells []Cell, grid [][]int) {
	for _, cell := range cells {
		if isCellValid(cell.y, cell.x, grid) {
			grid[cell.y][cell.x] = 0
		}
	}
}

func glider(y, x int) []Cell {
	glider := make([]Cell, 0, 5)
	glider = append(glider, Cell {y, x}, Cell {y + 1, x + 1}, Cell {y + 1, x + 2}, Cell {y, x + 2}, Cell {y - 1, x + 2})
	return glider
}

func update(grid[][]int) {
	alive := make([]Cell, 0, 100)
	dead := make([]Cell, 0, 100)

	for i := range grid {
		for j := range grid[i] {
			n := getLiveNeighbors(i, j, grid)
			if grid[i][j] == 1 && n < 2 {
				dead = append(dead, Cell {int(i), int(j)})
			}
			if grid[i][j] == 1 && n > 3 {
				dead = append(dead, Cell {int(i), int(j)})
			}
			if grid[i][j] == 1 && n >= 2 && n <= 3 {
				alive = append(alive, Cell {int(i), int(j)})
			}
			if grid[i][j] == 0 && n == 3 {
				alive = append(alive, Cell {int(i), int(j)})
			}
		}
	}

	draw(alive, grid)
	remove(dead, grid)
}

func main() {
	grid := createGrid(20, 20)
	draw(glider(10, 10), grid)
	printGrid(grid)
	update(grid)
	printGrid(grid)

}
