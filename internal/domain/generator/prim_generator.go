package generator

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"
)

type PrimMazeGenerator struct {
}

func (pmg PrimMazeGenerator) Generate(width, height int) [][]cell.Cell {
	maze := make([][]cell.Cell, height)

	// Создается лабиринт, полностью состоящий из стен.
	for i := 0; i < height; i++ {
		maze[i] = make([]cell.Cell, width)
		for j := 0; j < width; j++ {
			maze[i][j] = cell.Cell{X: j, Y: i, T: cell.Wall}
		}
	}

	x, y := pmg.getRandomXY(width, height)

	maze[y][x].T = getRandomFloorType()

	// Инициализация слайса ячеек, которые будут становится проходами.
	var cellsToClear []cell.Cell

	isOccupied := make(map[cell.Cell]bool)

	// Добавление соседних от освобожденной ячеек для дальнейшего освобождения.
	if x+2 < width-1 {
		cellsToClear = append(cellsToClear, maze[y][x+2])
		isOccupied[maze[y][x+2]] = true
	}

	if x-2 > 0 {
		cellsToClear = append(cellsToClear, maze[y][x-2])
		isOccupied[maze[y][x-2]] = true
	}

	if y+2 < height-1 {
		cellsToClear = append(cellsToClear, maze[y+2][x])
		isOccupied[maze[y+2][x]] = true
	}

	if y-2 > 0 {
		cellsToClear = append(cellsToClear, maze[y-2][x])
		isOccupied[maze[y-2][x]] = true
	}

	for len(cellsToClear) > 0 {
		// Выбираем случайную точку очищаем ее и удаляем из списка.
		randomIndex, _ := getRandomInt(0, len(cellsToClear))

		// Освобождаем ячейку.
		maze[cellsToClear[randomIndex].Y][cellsToClear[randomIndex].X].T = getRandomFloorType()
		y, x = cellsToClear[randomIndex].Y, cellsToClear[randomIndex].X
		cellsToClear = append(cellsToClear[:randomIndex], cellsToClear[randomIndex+1:]...)

		// Для новой свободной клетки находится сосед и они объединяются.
		directions := []int{down, up, left, right}
		for len(directions) > 0 {
			// Выбирается рандомное направление, которое будет очищаться.
			randomInd, _ := getRandomInt(0, len(directions))

			randomDirection := directions[randomInd]
			pmg.clearRandNeighbour(randomDirection, y, x, maze, &directions)

			if len(directions) <= 1 {
				break
			}

			directions = append(directions[:randomInd], directions[randomInd+1:]...)
		}

		// Добавление следующей ячейки.
		pmg.appendCellsToClear(y, x, maze, isOccupied, &cellsToClear)
	}

	return maze
}

// Получить случайнуя ячейку для очистки в зависимости от размеров.
func (PrimMazeGenerator) getRandomXY(width, height int) (x, y int) {
	if width%2 == 0 {
		x, _ = getRandomInt(1, width/2)
		x *= 2
	} else {
		x, _ = getRandomInt(0, height/2)
		x = x*2 + 1
	}

	if height%2 == 0 {
		y, _ = getRandomInt(1, height/2)
		y *= 2
	} else {
		y, _ = getRandomInt(0, height/2)
		y = y*2 + 1
	}

	return x, y
}

func (pmg PrimMazeGenerator) clearRandNeighbour(randomDirection, y, x int, maze [][]cell.Cell, directions *[]int) {
	switch randomDirection {
	case down:
		if y-2 >= 0 {
			if maze[y-2][x].T != cell.Wall {
				maze[y-1][x].T = getRandomFloorType()

				// Нужно, чтобы лабиринт мог иметь больше одного решения.
				pmg.randNilDirections(directions)
			}
		}
	case up:
		if y+2 < len(maze) {
			if maze[y+2][x].T != cell.Wall {
				maze[y+1][x].T = getRandomFloorType()

				pmg.randNilDirections(directions)
			}
		}
	case left:
		if x-2 >= 0 {
			if maze[y][x-2].T != cell.Wall {
				maze[y][x-1].T = getRandomFloorType()

				pmg.randNilDirections(directions)
			}
		}
	case right:
		if x+2 < len(maze[0]) {
			if maze[y][x+2].T != cell.Wall {
				maze[y][x+1].T = getRandomFloorType()

				pmg.randNilDirections(directions)
			}
		}
	}
}

func (PrimMazeGenerator) randNilDirections(directions *[]int) {
	rand, _ := getRandomInt(1, 4)
	if rand != 3 {
		*directions = (*directions)[:0]
	}
}

func (PrimMazeGenerator) appendCellsToClear(y, x int, maze [][]cell.Cell, isOccupied map[cell.Cell]bool, cellsToClear *[]cell.Cell) {
	// Добавление следующей ячейки.
	if y-2 > 0 {
		if maze[y-2][x].T == cell.Wall {
			if !isOccupied[maze[y-2][x]] {
				*cellsToClear = append(*cellsToClear, maze[y-2][x])
				isOccupied[maze[y-2][x]] = true
			}
		}
	}

	if y+2 < len(maze)-1 {
		if maze[y+2][x].T == cell.Wall {
			if !isOccupied[maze[y+2][x]] {
				*cellsToClear = append(*cellsToClear, maze[y+2][x])
				isOccupied[maze[y+2][x]] = true
			}
		}
	}

	if x-2 > 0 {
		if maze[y][x-2].T == cell.Wall {
			if !isOccupied[maze[y][x-2]] {
				*cellsToClear = append(*cellsToClear, maze[y][x-2])
				isOccupied[maze[y][x-2]] = true
			}
		}
	}

	if x+2 < len(maze[0])-1 {
		if maze[y][x+2].T == cell.Wall {
			if !isOccupied[maze[y][x+2]] {
				*cellsToClear = append(*cellsToClear, maze[y][x+2])
				isOccupied[maze[y][x+2]] = true
			}
		}
	}
}
