package main

import (
	"fmt"
	"time"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Cell struct {
	y int
	x int
}

type Grid [][]int

func (grid Grid) printGrid() {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%d, ", grid[i][j])
		}
		fmt.Println()
	}
}

func createGrid(y int, x int) Grid {
	grid := make([][]int, y)
	for i := range grid {
		grid[i] = make([]int, x)
	}
	return grid
}

func (grid Grid) getCell(y, x int) int {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0]) {
		return grid[y][x]
	}
	return 0
}

func (grid Grid) isCellValid(y, x int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}

func (grid Grid) getLiveNeighbors(y int, x int) int {
	return grid.getCell(y - 1, x) + grid.getCell(y + 1, x) + grid.getCell(y, x - 1) + grid.getCell(y, x + 1) + grid.getCell(y + 1, x + 1) + grid.getCell(y + 1, x - 1) + grid.getCell(y - 1, x + 1) + grid.getCell(y - 1, x - 1)
}

func (grid Grid) draw(cells []Cell) {
	for _, cell := range cells {
		if grid.isCellValid(cell.y, cell.x) {
			grid[cell.y][cell.x] = 1
		}
	}
}

func (grid Grid) remove(cells []Cell) {
	for _, cell := range cells {
		if grid.isCellValid(cell.y, cell.x) {
			grid[cell.y][cell.x] = 0
		}
	}
}

func glider(y, x int) []Cell {
	glider := make([]Cell, 0, 5)
	glider = append(glider, Cell {y, x}, Cell {y + 1, x + 1}, Cell {y + 1, x + 2}, Cell {y, x + 2}, Cell {y - 1, x + 2})
	return glider
}

func worker_bee(y, x int) []Cell {
	bee := make([]Cell, 0, 3)
	bee = append(bee, Cell {y, x}, Cell {y, x + 1}, Cell {y, x + 2})
	return bee
}

func (grid Grid) update() {
	alive := make([]Cell, 0, 100)
	dead := make([]Cell, 0, 100)
	for i := range grid {
		for j := range grid[i] {
			n := grid.getLiveNeighbors(i, j)
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

	grid.draw(alive)
	grid.remove(dead)
}

func updateScreen(imd *imdraw.IMDraw, grid Grid) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				imd.Color = colornames.Blueviolet
				imd.Push(pixel.V(float64(j * 10), float64(i * 10)), pixel.V(float64(j * 10 + 10), float64(i * 10 + 10)))
				imd.Rectangle(0)
			}
		}
	}
}

func run() {
	grid := createGrid(80, 60)
	grid.draw(glider(10, 10))
	grid.draw(glider(40, 20))
	grid.draw(worker_bee(30, 30))

	cfg := pixelgl.WindowConfig {
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync: true,
	}

	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	for !win.Closed() {
		imd.Clear()
		updateScreen(imd, grid)
		win.Clear(colornames.Skyblue)
		imd.Draw(win)
		win.Update()
		time.Sleep(300000000)
		grid.update()
	}
}

func main() {
	pixelgl.Run(run)
}
