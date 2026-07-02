package generator

import (
	"crypto/rand"
	"math/big"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"
)

// Интерфейс, определяющий сигнатуру функции генерации лабиринта.
type Generatable interface {
	Generate(width, height int) [][]cell.Cell
}

// Возвращает случайное число в полуинтервале [lower_bound, upper_bound).
func getRandomInt(lowerBound, upperBound int) (int, error) {
	diff := big.NewInt(int64(upperBound - lowerBound))
	randomValue, err := rand.Int(rand.Reader, diff)

	return int(randomValue.Int64()) + lowerBound, err
}

// Используется для определения частоты появления болот.
// Вероятность того, что клетка - болото определяется, как 1/swampK+1.
const swampK = 4

// Получает тип поверхности в зависимости от swampK.
func getRandomFloorType() int {
	randomInt, _ := getRandomInt(0, swampK+1)
	if randomInt == swampK {
		return cell.Swamp
	}

	return cell.Free
}

// Определяет направления выбора ячеек при генерации лабиринтов.
const (
	up = iota
	right
	down
	left
	self
)
