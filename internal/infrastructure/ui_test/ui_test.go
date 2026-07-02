package ui_test

import (
	"io"
	"os"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/cell"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/ui"
	"github.com/stretchr/testify/assert"
)

func TestCreatePicture(t *testing.T) {
	newUI := ui.NewUI(os.Stdin, io.Discard)

	grid1 := [][]cell.Cell{
		{cell.Cell{X: 0, Y: 0, T: cell.Wall}, cell.Cell{X: 1, Y: 0, T: cell.Swamp}},
		{cell.Cell{X: 0, Y: 1, T: cell.Free}, cell.Cell{X: 1, Y: 1, T: cell.Wall}},
		{cell.Cell{X: 0, Y: 2, T: cell.Swamp}, cell.Cell{X: 1, Y: 2, T: cell.Free}},
	}

	pictureExpected1 := [][]rune{
		{'⬛', '🟩'},
		{'⬜', '⬛'},
		{'🟩', '⬜'},
	}

	picture1 := newUI.CreatePicture(grid1)
	assert.Equal(t, pictureExpected1, picture1, "Тест сопоставления типу ячейку символа")

	grid2 := [][]cell.Cell{}
	pictureExpected2 := [][]rune{}
	picture2 := newUI.CreatePicture(grid2)
	assert.Equal(t, pictureExpected2, picture2, "Тест пустой картинки")
}

func TestAddPathToPicture(t *testing.T) {
	newUI := ui.NewUI(os.Stdin, io.Discard)

	grid1 := [][]cell.Cell{
		{cell.Cell{X: 0, Y: 0, T: cell.Free}, cell.Cell{X: 1, Y: 0, T: cell.Free}, cell.Cell{X: 2, Y: 0, T: cell.Free}},
		{cell.Cell{X: 0, Y: 1, T: cell.Free}, cell.Cell{X: 1, Y: 1, T: cell.Free}, cell.Cell{X: 2, Y: 1, T: cell.Free}},
		{cell.Cell{X: 0, Y: 2, T: cell.Free}, cell.Cell{X: 1, Y: 2, T: cell.Free}, cell.Cell{X: 2, Y: 2, T: cell.Free}},
	}

	path1 := map[cell.Cell]cell.Cell{
		{X: 2, Y: 0, T: cell.Free}: {X: 2, Y: 1, T: cell.Free},
		{X: 2, Y: 1, T: cell.Free}: {X: 1, Y: 1, T: cell.Free},
		{X: 1, Y: 1, T: cell.Free}: {X: 0, Y: 1, T: cell.Free},
		{X: 0, Y: 1, T: cell.Free}: {X: 0, Y: 2, T: cell.Free},
	}

	picture1 := [][]rune{
		{'⬜', '⬜', '⬜'},
		{'⬜', '⬜', '⬜'},
		{'⬜', '⬜', '⬜'},
	}

	updatedPicture1 := newUI.AddPathToPicture(picture1, grid1, 0, 2, 2, 0, path1)
	pathPictureExpected1 := [][]rune{
		{'⬜', '⬜', '🟥'},
		{'🟥', '🟥', '🟥'},
		{'🟥', '⬜', '⬜'},
	}
	assert.Equal(t, pathPictureExpected1, updatedPicture1, "Тест добавления маршрута в разные стобцы\\строки")

	grid2 := [][]cell.Cell{
		{cell.Cell{X: 0, Y: 0, T: cell.Free}, cell.Cell{X: 1, Y: 0, T: cell.Free}, cell.Cell{X: 2, Y: 0, T: cell.Free}},
		{cell.Cell{X: 0, Y: 1, T: cell.Free}, cell.Cell{X: 1, Y: 1, T: cell.Free}, cell.Cell{X: 2, Y: 1, T: cell.Free}},
		{cell.Cell{X: 0, Y: 2, T: cell.Free}, cell.Cell{X: 1, Y: 2, T: cell.Free}, cell.Cell{X: 2, Y: 2, T: cell.Free}},
	}

	path2 := map[cell.Cell]cell.Cell{
		{X: 2, Y: 0, T: cell.Free}: {X: 2, Y: 1, T: cell.Free},
		{X: 2, Y: 1, T: cell.Free}: {X: 2, Y: 2, T: cell.Free},
	}

	picture2 := [][]rune{
		{'⬜', '⬜', '⬜'},
		{'⬜', '⬜', '⬜'},
		{'⬜', '⬜', '⬜'},
	}

	updatedPicture2 := newUI.AddPathToPicture(picture2, grid2, 2, 2, 2, 0, path2)
	pathPictureExpected2 := [][]rune{
		{'⬜', '⬜', '🟥'},
		{'⬜', '⬜', '🟥'},
		{'⬜', '⬜', '🟥'},
	}
	assert.Equal(t, pathPictureExpected2, updatedPicture2, "Тест добавления маршрута в один столбец")
}
