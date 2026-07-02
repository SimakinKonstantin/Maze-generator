package solver

import (
	"math"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"
)

type AStarSolver struct {
}

func (as AStarSolver) FindPath(maze [][]cell.Cell, startX, startY, endX, endY int) (path map[cell.Cell]cell.Cell, dist int, err error) {
	areValid := checkStartFinishValidity(maze, startX, startY, endX, endY)
	if !areValid {
		return nil, -1, ValidityError{"между запрашиваемыми ячейками нельзя найти путь"}
	}

	start := cell.Cell{X: startX, Y: startY, T: maze[startY][startX].T}
	end := cell.Cell{X: endX, Y: endY, T: maze[endY][endX].T}

	toProcess := priorityQueue{}
	toProcess.Push(priorityQueueElem{Cell: start, Priority: 0})

	// Ключ - текущая ячейка, значение - откуда пришли в нее.
	path = make(map[cell.Cell]cell.Cell)

	// Расстояние от начальной ячейки до текущей.
	gValue := make(map[cell.Cell]int)

	for toProcess.Len() != 0 {
		// Получаем точку с наименьшим приоритетом - расстоянием от start до точки.
		curCell, err := toProcess.Pop()
		if err != nil {
			return nil, -1, err
		}

		if curCell.Cell == end {
			return path, curCell.Priority, nil
		}

		// Все возможные шаги из текущей ячейки.
		stepsToNeighbour := make([][]int, 4)
		stepsToNeighbour[0] = []int{0, 1}
		stepsToNeighbour[1] = []int{0, -1}
		stepsToNeighbour[2] = []int{1, 0}
		stepsToNeighbour[3] = []int{-1, 0}

		for i := 0; i < len(stepsToNeighbour); i++ {
			y := stepsToNeighbour[i][0] + curCell.Cell.Y
			x := stepsToNeighbour[i][1] + curCell.Cell.X

			// Получаем одного из 4 соседних ячеек.
			neighbour := maze[y][x]
			neighbourCurGValue := gValue[curCell.Cell] + neighbour.T

			// Проверка, что сосед лежит внутри лабиринта и не стена.
			if (y > 0 && y < len(maze)-1) && (x > 0 && x < len(maze[0])-1) && (neighbour.T != cell.Wall) {
				if _, ok := gValue[neighbour]; !ok || neighbourCurGValue < gValue[neighbour] {
					gValue[neighbour] = neighbourCurGValue
					fValue := neighbourCurGValue + as.heuristic(neighbour, end)
					toProcess.Push(priorityQueueElem{Cell: neighbour, Priority: fValue})
					path[neighbour] = curCell.Cell
				}
			}
		}
	}

	return make(map[cell.Cell]cell.Cell), -1, nil
}

func (AStarSolver) heuristic(start, finish cell.Cell) int {
	return int(math.Abs(float64(finish.X-start.X)) + math.Abs(float64(finish.Y-start.Y)))
}
