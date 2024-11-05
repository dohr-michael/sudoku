package sudoku

import (
	"errors"
	"sync"
)

func resolveMultiThread(grid *Grid, row, col int, fn func(*Grid)) *Grid {
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

	wg := new(sync.WaitGroup)
	wg.Add(len(availableValues))
	results := make(chan *Grid)

	next := func(grid *Grid, value, row, col, nextRow, nextCol int) {
		defer wg.Done()
		grid.UpdateValue(value, row, col)
		tmp := resolveMultiThread(grid, nextRow, nextCol, fn)
		if tmp != nil {
			results <- tmp
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for _, value := range availableValues {
		go next(grid.Copy(), value, row, col, nextRow, nextCol)
	}

	allResults := make([]*Grid, 0)
	for r := range results {
		allResults = append(allResults, r)
	}

	if len(allResults) > 0 {
		return allResults[0]
	}
	return nil
}

func resolve(grid *Grid, row, col int, fn func(*Grid)) *Grid {
	if grid.IsValid() {
		fn(grid)
	}
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
		tmp := resolve(gc, nextRow, nextCol, fn)
		if tmp != nil {
			return tmp
		}
	}
	return nil
}

func Resolve(input string, fn func(*Grid)) (*Grid, error) {
	base, err := NewGrid(input)
	if err != nil {
		return nil, err
	}
	if !base.IsValid() {
		return nil, errors.New("invalid sudoku grid")
	}
	fn(base)
	result := resolve(base, 1, 1, fn)
	if result == nil {
		return nil, errors.New("could not resolve sudoku grid")
	}
	return result, nil
}
