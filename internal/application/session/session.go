package session

import (
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generator"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solver"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/ui"
)

func Run() error {
	newUI := ui.NewUI(os.Stdin, os.Stdout)

	height, err := newUI.InputInt("Введите высоту лабиринта (в ячейках): ", 7, 100)
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	width, err := newUI.InputInt("Введите ширину лабиринта (в ячейках): ", 7, 100)
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	algo, err := newUI.ChooseAlg()
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	var mazeGrid [][]cell.Cell
	if algo == "a" {
		mazeGrid = generator.BinaryMazeGenerator{}.Generate(height, width)
	} else if algo == "b" {
		mazeGrid = generator.PrimMazeGenerator{}.Generate(height, width)
	}

	mazePicture := newUI.CreatePicture(mazeGrid)

	err = newUI.DrawPicture(mazePicture, "Сгенерированный лабиринт:")
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	startX, startY, err := newUI.InputCell("Введите координату X начала: ", "Введите координату Y начала: ")
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	finishX, finishY, err := newUI.InputCell("Введите координату X конца: ", "Введите координату Y конца: ")
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	// Решение задачи через алгоритм Дейкстры.
	dijkstraResult, _, err := solver.DijkstraSolver{}.FindPath(mazeGrid, startX, startY, finishX, finishY)
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	mazePictureDijkstra := newUI.AddPathToPicture(mazePicture, mazeGrid, startX, startY, finishX, finishY, dijkstraResult)

	err = newUI.DrawPicture(mazePictureDijkstra, "Решение с помощью алгоритма Дейкстры:")
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	// Решение задачи через алгоритм A*.
	AstarResult, _, err := solver.AStarSolver{}.FindPath(mazeGrid, startX, startY, finishX, finishY)
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	mazePictureAStar := newUI.AddPathToPicture(mazePicture, mazeGrid, startX, startY, finishX, finishY, AstarResult)

	err = newUI.DrawPicture(mazePictureAStar, "Решение с помощью алгоритма A*:")
	if err != nil {
		return AppError{"ошибка сессии", err}
	}

	return nil
}
