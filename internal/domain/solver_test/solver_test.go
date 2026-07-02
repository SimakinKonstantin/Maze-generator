package solver_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solver"
	"github.com/stretchr/testify/assert"
)

func TestCountDistance(t *testing.T) {
	maze1 := [][]cell.Cell{
		{{X: 0, Y: 0, T: cell.Wall}, {X: 1, Y: 0, T: cell.Wall}, {X: 2, Y: 0, T: cell.Wall}, {X: 3, Y: 0, T: cell.Wall},
			{X: 4, Y: 0, T: cell.Wall}, {X: 5, Y: 0, T: cell.Wall}},
		{{X: 0, Y: 1, T: cell.Wall}, {X: 1, Y: 1, T: cell.Free}, {X: 2, Y: 1, T: cell.Wall}, {X: 3, Y: 1, T: cell.Free},
			{X: 4, Y: 1, T: cell.Free}, {X: 5, Y: 1, T: cell.Wall}},
		{{X: 0, Y: 2, T: cell.Wall}, {X: 1, Y: 2, T: cell.Free}, {X: 2, Y: 2, T: cell.Wall}, {X: 3, Y: 2, T: cell.Free},
			{X: 4, Y: 2, T: cell.Wall}, {X: 5, Y: 2, T: cell.Wall}},
		{{X: 0, Y: 3, T: cell.Wall}, {X: 1, Y: 3, T: cell.Free}, {X: 2, Y: 3, T: cell.Free}, {X: 3, Y: 3, T: cell.Free},
			{X: 4, Y: 3, T: cell.Free}, {X: 5, Y: 3, T: cell.Wall}},
		{{X: 0, Y: 4, T: cell.Wall}, {X: 1, Y: 4, T: cell.Free}, {X: 2, Y: 4, T: cell.Swamp}, {X: 3, Y: 4, T: cell.Wall},
			{X: 4, Y: 4, T: cell.Free}, {X: 5, Y: 4, T: cell.Wall}},
		{{X: 0, Y: 5, T: cell.Wall}, {X: 1, Y: 5, T: cell.Wall}, {X: 2, Y: 5, T: cell.Wall}, {X: 3, Y: 5, T: cell.Wall},
			{X: 4, Y: 5, T: cell.Wall}, {X: 5, Y: 5, T: cell.Wall}},
	}

	// Лабринт для теста 1 🟦 - начало, конец
	// ⬛⬛⬛⬛⬛⬛
	// ⬛⬜⬛⬜🟦⬛
	// ⬛⬜⬛⬜⬛⬛
	// ⬛⬜⬜⬜⬜⬛
	// ⬛🟦🟩⬛⬜⬛
	// ⬛⬛⬛⬛⬛⬛

	expectedDist1 := 6

	startX1, startY1, endX1, endY1 := 1, 4, 4, 1

	_, dijkstraLen1, _ := solver.DijkstraSolver{}.FindPath(maze1, startX1, startY1, endX1, endY1)
	assert.Equal(t, expectedDist1, dijkstraLen1, "При Дейкстре должен обойти болото и дойти до конца")

	_, astarLen1, _ := solver.AStarSolver{}.FindPath(maze1, startX1, startY1, endX1, endY1)
	assert.Equal(t, expectedDist1, astarLen1, "При A* должен обойти болото и дойти до конца")

	maze2 := [][]cell.Cell{
		{{X: 0, Y: 0, T: cell.Wall}, {X: 1, Y: 0, T: cell.Wall}, {X: 2, Y: 0, T: cell.Wall}, {X: 3, Y: 0, T: cell.Wall},
			{X: 4, Y: 0, T: cell.Wall}, {X: 5, Y: 0, T: cell.Wall}},
		{{X: 0, Y: 1, T: cell.Wall}, {X: 1, Y: 1, T: cell.Free}, {X: 2, Y: 1, T: cell.Wall}, {X: 3, Y: 1, T: cell.Free},
			{X: 4, Y: 1, T: cell.Free}, {X: 5, Y: 1, T: cell.Wall}},
		{{X: 0, Y: 2, T: cell.Wall}, {X: 1, Y: 2, T: cell.Free}, {X: 2, Y: 2, T: cell.Wall}, {X: 3, Y: 2, T: cell.Free},
			{X: 4, Y: 2, T: cell.Wall}, {X: 5, Y: 2, T: cell.Wall}},
		{{X: 0, Y: 3, T: cell.Wall}, {X: 1, Y: 3, T: cell.Free}, {X: 2, Y: 3, T: cell.Swamp}, {X: 3, Y: 3, T: cell.Swamp},
			{X: 4, Y: 3, T: cell.Free}, {X: 5, Y: 3, T: cell.Wall}},
		{{X: 0, Y: 4, T: cell.Wall}, {X: 1, Y: 4, T: cell.Free}, {X: 2, Y: 4, T: cell.Swamp}, {X: 3, Y: 4, T: cell.Wall},
			{X: 4, Y: 4, T: cell.Free}, {X: 5, Y: 4, T: cell.Wall}},
		{{X: 0, Y: 5, T: cell.Wall}, {X: 1, Y: 5, T: cell.Wall}, {X: 2, Y: 5, T: cell.Wall}, {X: 3, Y: 5, T: cell.Wall},
			{X: 4, Y: 5, T: cell.Wall}, {X: 5, Y: 5, T: cell.Wall}},
	}

	// Лабринт для теста 2 🟦 - начало, конец
	// ⬛⬛⬛⬛⬛⬛
	// ⬛⬜⬛⬜⬜⬛
	// ⬛⬜⬛⬜⬛⬛
	// ⬛⬜🟩🟩⬜⬛
	// ⬛🟦🟩⬛🟦⬛
	// ⬛⬛⬛⬛⬛⬛

	expectedDist2 := 13

	startX2, startY2, endX2, endY2 := 1, 4, 4, 4

	_, dijkstraLen2, _ := solver.DijkstraSolver{}.FindPath(maze2, startX2, startY2, endX2, endY2)
	assert.Equal(t, expectedDist2, dijkstraLen2, "При Дейкстре должен обойти нижнее болото и попасть в оставшиеся")

	_, astarLen2, _ := solver.AStarSolver{}.FindPath(maze2, startX2, startY2, endX2, endY2)
	assert.Equal(t, expectedDist2, astarLen2, "При A* должен должен обойти нижнее болото и попасть в оставшиеся")

	maze3 := [][]cell.Cell{
		{{X: 0, Y: 0, T: cell.Wall}, {X: 1, Y: 0, T: cell.Wall}, {X: 2, Y: 0, T: cell.Wall}, {X: 3, Y: 0, T: cell.Wall},
			{X: 4, Y: 0, T: cell.Wall}, {X: 5, Y: 0, T: cell.Wall}},
		{{X: 0, Y: 1, T: cell.Wall}, {X: 1, Y: 1, T: cell.Free}, {X: 2, Y: 1, T: cell.Swamp}, {X: 3, Y: 1, T: cell.Free},
			{X: 4, Y: 1, T: cell.Free}, {X: 5, Y: 1, T: cell.Wall}},
		{{X: 0, Y: 2, T: cell.Wall}, {X: 1, Y: 2, T: cell.Free}, {X: 2, Y: 2, T: cell.Swamp}, {X: 3, Y: 2, T: cell.Swamp},
			{X: 4, Y: 2, T: cell.Free}, {X: 5, Y: 2, T: cell.Wall}},
		{{X: 0, Y: 3, T: cell.Wall}, {X: 1, Y: 3, T: cell.Free}, {X: 2, Y: 3, T: cell.Free}, {X: 3, Y: 3, T: cell.Swamp},
			{X: 4, Y: 3, T: cell.Free}, {X: 5, Y: 3, T: cell.Wall}},
		{{X: 0, Y: 4, T: cell.Wall}, {X: 1, Y: 4, T: cell.Free}, {X: 2, Y: 4, T: cell.Swamp}, {X: 3, Y: 4, T: cell.Free},
			{X: 4, Y: 4, T: cell.Free}, {X: 5, Y: 4, T: cell.Wall}},
		{{X: 0, Y: 5, T: cell.Wall}, {X: 1, Y: 5, T: cell.Wall}, {X: 2, Y: 5, T: cell.Wall}, {X: 3, Y: 5, T: cell.Wall},
			{X: 4, Y: 5, T: cell.Wall}, {X: 5, Y: 5, T: cell.Wall}},
	}

	// Лабринт для теста 3 🟦 - начало, конец
	// ⬛⬛⬛⬛⬛⬛
	// ⬛⬜🟩⬜🟦⬛
	// ⬛🟦🟩🟩⬜⬛
	// ⬛⬜⬜🟩⬜⬛
	// ⬛⬜🟩⬜⬜⬛
	// ⬛⬛⬛⬛⬛⬛

	expectedDist3 := 8

	startX3, startY3, endX3, endY3 := 1, 2, 4, 1

	_, dijkstraLen3, _ := solver.DijkstraSolver{}.FindPath(maze3, startX3, startY3, endX3, endY3)
	assert.Equal(t, expectedDist3, dijkstraLen3, "При Дейкстре должен попасть только в одно болото")

	_, astarLen3, _ := solver.AStarSolver{}.FindPath(maze3, startX3, startY3, endX3, endY3)
	assert.Equal(t, expectedDist3, astarLen3, "При A* должен должен попасть только в одно болото")
}

func TestCreatePath(t *testing.T) {
	maze1 := [][]cell.Cell{
		{{X: 0, Y: 0, T: cell.Wall}, {X: 1, Y: 0, T: cell.Wall}, {X: 2, Y: 0, T: cell.Wall}, {X: 3, Y: 0, T: cell.Wall},
			{X: 4, Y: 0, T: cell.Wall}, {X: 5, Y: 0, T: cell.Wall}},
		{{X: 0, Y: 1, T: cell.Wall}, {X: 1, Y: 1, T: cell.Free}, {X: 2, Y: 1, T: cell.Wall}, {X: 3, Y: 1, T: cell.Free},
			{X: 4, Y: 1, T: cell.Free}, {X: 5, Y: 1, T: cell.Wall}},
		{{X: 0, Y: 2, T: cell.Wall}, {X: 1, Y: 2, T: cell.Free}, {X: 2, Y: 2, T: cell.Free}, {X: 3, Y: 2, T: cell.Free},
			{X: 4, Y: 2, T: cell.Wall}, {X: 5, Y: 2, T: cell.Wall}},
		{{X: 0, Y: 3, T: cell.Wall}, {X: 1, Y: 3, T: cell.Wall}, {X: 2, Y: 3, T: cell.Wall}, {X: 3, Y: 3, T: cell.Free},
			{X: 4, Y: 3, T: cell.Free}, {X: 5, Y: 3, T: cell.Wall}},
		{{X: 0, Y: 4, T: cell.Wall}, {X: 1, Y: 4, T: cell.Free}, {X: 2, Y: 4, T: cell.Free}, {X: 3, Y: 4, T: cell.Free},
			{X: 4, Y: 4, T: cell.Free}, {X: 5, Y: 4, T: cell.Wall}},
		{{X: 0, Y: 5, T: cell.Wall}, {X: 1, Y: 5, T: cell.Wall}, {X: 2, Y: 5, T: cell.Wall}, {X: 3, Y: 5, T: cell.Wall},
			{X: 4, Y: 5, T: cell.Wall}, {X: 5, Y: 5, T: cell.Wall}},
	}

	// Лабринт для теста 1 🟦 - начало, конец
	// ⬛⬛⬛⬛⬛⬛
	// ⬛🟦⬛⬜⬜⬛
	// ⬛⬜⬜⬜⬛⬛
	// ⬛⬛⬛⬜⬜⬛
	// ⬛🟦⬜⬜⬜⬛
	// ⬛⬛⬛⬛⬛⬛

	startX1, startY1, endX1, endY1 := 1, 4, 1, 1

	expectedPath1 := []cell.Cell{
		{X: endX1, Y: endY1, T: cell.Free}, {X: 1, Y: 2, T: cell.Free}, {X: 2, Y: 2, T: cell.Free},
		{X: 3, Y: 2, T: cell.Free}, {X: 3, Y: 3, T: cell.Free}, {X: 3, Y: 4, T: cell.Free},
		{X: 2, Y: 4, T: cell.Free}, {X: startX1, Y: startY1, T: cell.Free},
	}

	// Проверка Дейкстры.
	unparsedDijkstraPath1, _, _ := solver.DijkstraSolver{}.FindPath(maze1, startX1, startY1, endX1, endY1)
	dijkstraPath1 := make([]cell.Cell, len(expectedPath1))

	cur := cell.Cell{X: endX1, Y: endY1, T: cell.Free}
	start := cell.Cell{X: startX1, Y: startY1, T: cell.Free}
	index := 0

	for cur != start {
		dijkstraPath1[index] = cur
		cur = unparsedDijkstraPath1[cur]
		index++
	}

	dijkstraPath1[index] = cur
	assert.Equal(t, expectedPath1, dijkstraPath1, "Дейкстра должен обойти стену и не застрять в тупике")

	// Проверка A*.
	unparsedAstarPath1, _, _ := solver.DijkstraSolver{}.FindPath(maze1, startX1, startY1, endX1, endY1)
	AstarPath1 := make([]cell.Cell, len(expectedPath1))

	cur = cell.Cell{X: endX1, Y: endY1, T: cell.Free}
	index = 0

	for cur != start {
		AstarPath1[index] = cur
		cur = unparsedAstarPath1[cur]
		index++
	}

	AstarPath1[index] = cur
	assert.Equal(t, expectedPath1, AstarPath1, "A* должен обойти стену и не застрять в тупике")

	maze2 := [][]cell.Cell{
		{{X: 0, Y: 0, T: cell.Wall}, {X: 1, Y: 0, T: cell.Wall}, {X: 2, Y: 0, T: cell.Wall}, {X: 3, Y: 0, T: cell.Wall},
			{X: 4, Y: 0, T: cell.Wall}, {X: 5, Y: 0, T: cell.Wall}},
		{{X: 0, Y: 1, T: cell.Wall}, {X: 1, Y: 1, T: cell.Free}, {X: 2, Y: 1, T: cell.Swamp}, {X: 3, Y: 1, T: cell.Wall},
			{X: 4, Y: 1, T: cell.Free}, {X: 5, Y: 1, T: cell.Wall}},
		{{X: 0, Y: 2, T: cell.Wall}, {X: 1, Y: 2, T: cell.Free}, {X: 2, Y: 2, T: cell.Swamp}, {X: 3, Y: 2, T: cell.Swamp},
			{X: 4, Y: 2, T: cell.Free}, {X: 5, Y: 2, T: cell.Wall}},
		{{X: 0, Y: 3, T: cell.Wall}, {X: 1, Y: 3, T: cell.Free}, {X: 2, Y: 3, T: cell.Swamp}, {X: 3, Y: 3, T: cell.Free},
			{X: 4, Y: 3, T: cell.Free}, {X: 5, Y: 3, T: cell.Wall}},
		{{X: 0, Y: 4, T: cell.Wall}, {X: 1, Y: 4, T: cell.Free}, {X: 2, Y: 4, T: cell.Swamp}, {X: 3, Y: 4, T: cell.Swamp},
			{X: 4, Y: 4, T: cell.Free}, {X: 5, Y: 4, T: cell.Wall}},
		{{X: 0, Y: 5, T: cell.Wall}, {X: 1, Y: 5, T: cell.Wall}, {X: 2, Y: 5, T: cell.Wall}, {X: 3, Y: 5, T: cell.Wall},
			{X: 4, Y: 5, T: cell.Wall}, {X: 5, Y: 5, T: cell.Wall}},
	}

	// Лабринт для теста 2 🟦 - начало, конец
	// ⬛⬛⬛⬛⬛⬛
	// ⬛⬜🟩⬛⬜⬛
	// ⬛⬜🟩🟩⬜⬛
	// ⬛⬜🟩⬜⬜⬛
	// ⬛🟦🟩🟩🟦⬛
	// ⬛⬛⬛⬛⬛⬛

	startX2, startY2, endX2, endY2 := 1, 4, 4, 4

	expectedPath2 := []cell.Cell{
		{X: endX2, Y: endY2, T: cell.Free}, {X: 4, Y: 3, T: cell.Free}, {X: 3, Y: 3, T: cell.Free},
		{X: 2, Y: 3, T: cell.Swamp}, {X: 1, Y: 3, T: cell.Free}, {X: startX2, Y: startY2, T: cell.Free},
	}

	// Проверка Дейкстры.
	unparsedDijkstraPath2, _, _ := solver.DijkstraSolver{}.FindPath(maze2, startX2, startY2, endX2, endY2)
	dijkstraPath2 := make([]cell.Cell, len(expectedPath2))

	cur = cell.Cell{X: endX2, Y: endY2, T: cell.Free}
	start = cell.Cell{X: startX2, Y: startY2, T: cell.Free}
	index = 0

	for cur != start {
		dijkstraPath2[index] = cur
		cur = unparsedDijkstraPath2[cur]
		index++
	}

	dijkstraPath2[index] = cur
	assert.Equal(t, expectedPath2, dijkstraPath2, "Дейкстра должен выбрать путь только через одно болото")

	// Проверка A*.
	unparsedAstarPath2, _, _ := solver.DijkstraSolver{}.FindPath(maze2, startX2, startY2, endX2, endY2)
	AstarPath2 := make([]cell.Cell, len(expectedPath2))

	cur = cell.Cell{X: endX2, Y: endY2, T: cell.Free}
	index = 0

	for cur != start {
		AstarPath2[index] = cur
		cur = unparsedAstarPath2[cur]
		index++
	}

	AstarPath2[index] = cur
	assert.Equal(t, expectedPath2, AstarPath2, "A* должен выбрать путь только через одно болото")
}
