package generator

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"

type BinaryMazeGenerator struct {
}

func (bmg BinaryMazeGenerator) Generate(width, height int) [][]cell.Cell {
	maze := make([][]cell.Cell, height)

	// Создается лабиринт, полностью состоящий из стен.
	for i := 0; i < height; i++ {
		maze[i] = make([]cell.Cell, width)
		for j := 0; j < width; j++ {
			maze[i][j] = cell.Cell{X: j, Y: i, T: cell.Wall}
		}
	}

	for i := 1; i < height-1; i += 2 {
		for j := 1; j < width-1; j += 2 {
			// Получаем направление освобождения ячеек.
			direction := bmg.getDirection(i, j, height, width)

			floorType := getRandomFloorType()

			switch direction {
			case down:
				// Нужно, чтобы получался лабиринт с несколькими решениями.
				var maxA int
				if i+4 < height {
					maxA = 4
				} else {
					maxA = 2
				}

				for a := 0; a < maxA; a++ {
					maze[i+a][j].T = floorType
				}

			case right:
				// Нужно, чтобы получался лабиринт с несколькими решениями.
				var maxA int
				if j+4 < width {
					maxA = 4
				} else {
					maxA = 2
				}

				for a := 0; a < maxA; a++ {
					maze[i][j+a].T = floorType
				}

			case self:
				maze[i][j].T = floorType
			}
		}
	}

	return maze
}

func (BinaryMazeGenerator) getDirection(column, row, height, width int) int {
	var direction int

	switch {
	case column+2 == height && row+2 == width:
		direction = self
	case (column+2 == height && height%2 != 0) || (column+3 == height && height%2 == 0):
		direction = right
	case (row+2 == width && width%2 != 0) || (row+3 == width && width%2 == 0):
		direction = down
	default:
		directions := []int{down, right, self}
		randomInd, _ := getRandomInt(0, 2)
		direction = directions[randomInd]
	}

	return direction
}
