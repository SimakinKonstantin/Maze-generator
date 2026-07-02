package solver

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"

type DijkstraSolver struct {
}

// В функциях принимаем MAX_INT за бесконечность, так как даже в массиве 1000*1000 максимальная длина пути если все
// ячейки - болота == 5 * 1000 * 1000 << MAX_INT.
const inf = int(^uint(0) >> 1)

func (ds DijkstraSolver) FindPath(maze [][]cell.Cell, startX, startY, endX, endY int) (path map[cell.Cell]cell.Cell, dist int, err error) {
	areValid := checkStartFinishValidity(maze, startX, startY, endX, endY)
	if !areValid {
		return nil, -1, ValidityError{"между запрашиваемыми ячейками нельзя найти путь"}
	}

	start := cell.Cell{X: startX, Y: startY, T: maze[startY][startX].T}
	end := cell.Cell{X: endX, Y: endY, T: maze[endY][endX].T}
	label, isFinalLabel, parents := ds.initSlices(maze, start)

	// Как parents, только используется в случае, если алгоритм упрется в тупик.
	curParents := make(map[cell.Cell]cell.Cell)
	stepsLeft := len(isFinalLabel) - 1
	current := start

	// Тело алгоритма.
	for stepsLeft > 0 {
		// Ищем всех соседей, которые не стены и еще нужно обрабатывать.
		ds.findNeighbours(current, maze, label, isFinalLabel, parents)

		// Ищем наименьший путь в данной итерации.
		minLabel := inf
		notValidLabelCell := cell.Cell{X: -1, Y: -1, T: -1}
		minLabelCell := notValidLabelCell

		// Выбираем наименьший путь.
		for k, v := range label {
			if v < minLabel && !isFinalLabel[k] {
				minLabel = v
				minLabelCell = k
			}
		}

		if minLabelCell != notValidLabelCell {
			isFinalLabel[minLabelCell] = true

			// Обновляется при переходе.
			curParents[minLabelCell] = current
			current = minLabelCell
			stepsLeft--
		} else {
			// Если тупик, то возвращаемся обратно
			current = curParents[current]
		}
	}

	return parents, label[end], nil
}

func (DijkstraSolver) findNeighbours(current cell.Cell, maze [][]cell.Cell, label map[cell.Cell]int,
	isFinalLabel map[cell.Cell]bool, parents map[cell.Cell]cell.Cell) {
	steps := []int{1, -1}
	for i := range steps {
		if current.X+steps[i] < len(maze[i]) && current.X+steps[i] >= 0 {
			// Если сосед - не стена, неизвестна окончательное расстояние до него.
			if maze[current.Y][current.X+steps[i]].T != cell.Wall && !isFinalLabel[maze[current.Y][current.X+steps[i]]] {
				// Если из текущей точки расстояние получается меньше, чем сохраненное расстояние до соседа, то обновляем его
				if label[maze[current.Y][current.X]]+maze[current.Y][current.X+steps[i]].T < label[maze[current.Y][current.X+steps[i]]] {
					label[maze[current.Y][current.X+steps[i]]] = label[maze[current.Y][current.X]] + maze[current.Y][current.X+steps[i]].T
					parents[maze[current.Y][current.X+steps[i]]] = current
				}
			}
		}

		if current.Y+steps[i] < len(maze) && current.Y+steps[i] >= 0 {
			if maze[current.Y+steps[i]][current.X].T != cell.Wall && !isFinalLabel[maze[current.Y+steps[i]][current.X]] {
				if label[maze[current.Y][current.X]]+maze[current.Y+steps[i]][current.X].T < label[maze[current.Y+steps[i]][current.X]] {
					label[maze[current.Y+steps[i]][current.X]] = label[maze[current.Y][current.X]] + maze[current.Y+steps[i]][current.X].T
					parents[maze[current.Y+steps[i]][current.X]] = current
				}
			}
		}
	}
}

// Выполняет инициализацию вспомогательных слайсов для алгоритма Дейкстры.
func (DijkstraSolver) initSlices(maze [][]cell.Cell, start cell.Cell) (label map[cell.Cell]int,
	isFinalLabel map[cell.Cell]bool, parents map[cell.Cell]cell.Cell) {
	// Расстояние до вершин.
	label = make(map[cell.Cell]int)

	// Является ли расстояние конечным - минимальным возможным
	isFinalLabel = make(map[cell.Cell]bool)

	// Ключ - текущая веришна, значение - её родитель.
	// Используется, когда нужно просчитать финальный путь от старта до финиша.
	parents = make(map[cell.Cell]cell.Cell)

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j].T != cell.Wall {
				label[maze[i][j]] = inf
				isFinalLabel[maze[i][j]] = false
				parents[maze[i][j]] = cell.Cell{X: start.X, Y: start.Y, T: start.T}
			}
		}
	}

	label[start] = 0
	isFinalLabel[start] = true
	parents[start] = start

	return label, isFinalLabel, parents
}
