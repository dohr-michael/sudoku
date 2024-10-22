package sudoku

import (
	"errors"
)

func resolve(grid *Grid, row, col int) *Grid {
	if grid.IsComplete() && grid.IsValid() {
		return grid
	}
	if row > 9 || col > 9 {
		return nil
	}
	availableValues := grid.AvailableValues(row, col)
	nextRow, nextCol := row, col
	if col < 9 {
		nextCol = col + 1
	} else {
		nextRow = row + 1
		nextCol = 1
	}
	for _, value := range availableValues {
		gc := grid.Copy()
		gc.UpdateValue(value, row, col)
		tmp := resolve(gc, nextRow, nextCol)
		if tmp != nil {
			return tmp
		}
	}
	return nil
}

func Resolve(input string) (*Grid, error) {
	base, err := NewGrid(input)
	if err != nil {
		return nil, err
	}
	if !base.IsValid() {
		return nil, errors.New("invalid sudoku grid")
	}
	result := resolve(base, 1, 1)
	if result == nil {
		return nil, errors.New("could not resolve sudoku grid")
	}
	return result, nil
}
