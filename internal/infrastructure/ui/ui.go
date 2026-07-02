package ui

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"
)

type UI struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewUI(reader io.Reader, writer io.Writer) *UI {
	return &UI{reader: bufio.NewReader(reader), writer: bufio.NewWriter(writer)}
}

// Запрос ввода целого числа в отрезке [min, max], с сообщением message.
func (ui UI) InputInt(message string, minVal, maxVal int) (int, error) {
	for {
		_, err := fmt.Fprint(ui.writer, message)
		if err != nil {
			return -1, OutputError{"ошибка вывода при запросе целого числа", err}
		}

		err = ui.writer.Flush()
		if err != nil {
			return -1, OutputError{"ошибка вывода при запросе целого числа", err}
		}

		byteValue, _, err := ui.reader.ReadLine()
		if err != nil {
			return -1, InputError{"ошибка ввода при запросе целого числа", err}
		}

		// Если значение не является числом, то запрашиваем новое.
		intValue, err := strconv.Atoi(string(byteValue))
		if err != nil {
			continue
		}

		if intValue >= minVal && intValue <= maxVal {
			return intValue, nil
		}
	}
}

func (ui UI) ChooseAlg() (string, error) {
	var res string

	for {
		_, err := fmt.Fprint(ui.writer, `Выберите способ генерации лабиринта:
a) Модификация алгоритма бинарного дерева.
b) Модификация алгоритма Прима.
`)
		if err != nil {
			return "", OutputError{"ошибка вывода при выборе алгоритма генерации", err}
		}

		err = ui.writer.Flush()
		if err != nil {
			return "", OutputError{"ошибка вывода при выборе алгоритма генерации", err}
		}

		byteRes, _, err := ui.reader.ReadLine()
		if err != nil {
			return "", InputError{"ошибка ввода при выборе алгоритма генерации", err}
		}

		res = string(byteRes)
		if res == "a" || res == "b" {
			break
		}
	}

	return res, nil
}

// Запрашивает координаты X, Y ячейки с соответствующим текстом.
func (ui UI) InputCell(messageWithX, messageWithY string) (x, y int, err error) {
	for {
		_, err = fmt.Fprint(ui.writer, messageWithX)
		if err != nil {
			return -1, -1, OutputError{"ошибка вывода при запросе координаты x", err}
		}

		err = ui.writer.Flush()
		if err != nil {
			return -1, -1, OutputError{"ошибка вывода при запросе координаты x", err}
		}

		xByte, _, err := ui.reader.ReadLine()
		if err != nil {
			return -1, -1, InputError{"ошибка при вводе координаты x", err}
		}

		x, err = strconv.Atoi(string(xByte))
		if err == nil {
			break
		}
	}

	for {
		_, err = fmt.Fprint(ui.writer, messageWithY)
		if err != nil {
			return -1, -1, OutputError{"ошибка вывода при запросе координаты y", err}
		}

		err = ui.writer.Flush()
		if err != nil {
			return -1, -1, OutputError{"ошибка вывода при запросе координаты y", err}
		}

		yByte, _, err := ui.reader.ReadLine()
		if err != nil {
			return -1, -1, InputError{"ошибка при вводе координаты y", err}
		}

		y, err = strconv.Atoi(string(yByte))
		if err == nil {
			break
		}
	}

	return x, y, nil
}

// Формирует изображение лабиринта из матрицы его ячеек.
// В полученном изображении лабиринта.
func (UI) CreatePicture(maze [][]cell.Cell) [][]rune {
	picture := make([][]rune, len(maze))

	// Т.к. матрица ячеек лабиринта перевернута относительно Oy.
	for i := 0; i < len(maze); i++ {
		picture[i] = make([]rune, len(maze[i]))

		for j := 0; j < len(maze[i]); j++ {
			switch maze[i][j].T {
			case cell.Wall:
				picture[i][j] = '⬛'
			case cell.Free:
				picture[i][j] = '⬜'
			case cell.Swamp:
				picture[i][j] = '🟩'
			}
		}
	}

	return picture
}

// Создает новое изображение с маршрутом, не изменяет исходное изображение.
func (ui UI) AddPathToPicture(picture [][]rune, maze [][]cell.Cell, startX, startY, endX, endY int, path map[cell.Cell]cell.Cell) [][]rune {
	newPicture := make([][]rune, len(picture))

	for i := range picture {
		newPicture[i] = make([]rune, len(picture[i]))
		copy(newPicture[i], picture[i])
	}

	end := cell.Cell{X: endX, Y: endY, T: maze[endY][endX].T}
	start := cell.Cell{X: startX, Y: startY, T: maze[startY][startX].T}

	cur := end

	for cur != start {
		newPicture[cur.Y][cur.X] = '🟥'
		cur = path[cur]
	}

	newPicture[start.Y][start.X] = '🟥'

	return newPicture
}

// Рисует сформированное изображение.
func (ui UI) DrawPicture(maze [][]rune, message string) error {
	_, err := fmt.Fprintln(ui.writer, message)
	if err != nil {
		return OutputError{"ошибка при выводе сообщения к рисунку", err}
	}

	err = ui.writer.Flush()
	if err != nil {
		return OutputError{"ошибка при выводе сообщения к рисунку", err}
	}

	for i := len(maze) - 1; i > -1; i-- {
		for j := 0; j < len(maze[i]); j++ {
			_, err = fmt.Fprintf(ui.writer, "%c", maze[i][j])
			if err != nil {
				return OutputError{"ошибка при выводе элемента рисунка", err}
			}

			err = ui.writer.Flush()
			if err != nil {
				return OutputError{"ошибка при выводе элемента рисунка", err}
			}
		}

		_, err = fmt.Fprint(ui.writer, "\n")
		if err != nil {
			return OutputError{"ошибка при переходе на следующую строку при выводе рисунка", err}
		}

		err = ui.writer.Flush()
		if err != nil {
			return OutputError{"ошибка при переходе на следующую строку при выводе рисунка", err}
		}
	}

	_, err = fmt.Fprintln(ui.writer)
	if err != nil {
		return OutputError{"ошибка при переходе на следующую строку при выводе рисунка", err}
	}

	return nil
}
