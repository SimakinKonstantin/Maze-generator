package solver

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"

type Solverable interface {
	FindPath([][]cell.Cell, int, int, int, int) (map[cell.Cell]cell.Cell, int)
}

func checkStartFinishValidity(maze [][]cell.Cell, startX, startY, endX, endY int) bool {
	// Проверка на выход из массива.
	if startX < 0 || startY < 0 {
		return false
	}

	if endX >= len(maze[0]) || endY >= len(maze) {
		return false
	}

	// Проверка на валидность стартовой ячейки.
	if maze[startY][startX].T == cell.Wall {
		return false
	}

	if maze[endY][endX].T == cell.Wall {
		return false
	}

	return true
}
