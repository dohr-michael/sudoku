package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

// Cell represents a cell in a Sudoku grid, holding its value, initialization state, row, and column.
type Cell struct {
	Value         int
	IsInitialized bool
	Row           int
	Col           int
}

// Region calculates the region number of the cell in a 9x9 Sudoku grid.
func (c *Cell) Region() int {
	return ((c.Col-1)/3 + 1) + 3*((c.Row-1)/3)
}

// Copy creates a deep copy of the Cell instance and returns a pointer to the new Cell.
func (c *Cell) Copy() *Cell {
	return &Cell{
		Value:         c.Value,
		IsInitialized: c.IsInitialized,
		Row:           c.Row,
		Col:           c.Col,
	}
}

// Grid represents a 9x9 Sudoku grid containing cells organized into rows, columns, and regions.
type Grid struct {
	Rows    [][]*Cell
	cols    [][]*Cell
	regions [][]*Cell
}

// NewGrid creates a new 9x9 Sudoku grid from the given string of initial values.
// The string should contain 9 lines of 9 characters each, separated by spaces.
// Valid characters are digits (1-9) and '-' or '0' for empty cells.
// Returns a pointer to the Grid and an error if the input is invalid.
func NewGrid(initialValues string) (*Grid, error) {
	initialValuesRows := strings.Fields(initialValues)
	if len(initialValuesRows) != 9 {
		return nil, fmt.Errorf("invalid number of rows: %d", len(initialValuesRows))
	}
	for idx, row := range initialValuesRows {
		if len(row) != 9 {
			return nil, fmt.Errorf("invalid number of cells (at line %d): %d", idx+1, len(row))
		}
	}
	rows := make([][]*Cell, 9)
	for i := 0; i < 9; i++ {
		cols := make([]*Cell, 9)
		for j := 0; j < 9; j++ {
			current := string(initialValuesRows[i][j])
			value := 0
			if current != "-" && current != "0" {
				v, err := strconv.Atoi(current)
				if err != nil {
					return nil, fmt.Errorf("invalid number (at cell %d,%d), %e", i+1, j+1, err)
				}
				value = v
			}
			cols[j] = &Cell{Value: value, IsInitialized: value != 0, Row: i + 1, Col: j + 1}
		}
		rows[i] = cols
	}
	res := &Grid{Rows: rows}
	res.prepare()
	return res, nil
}

// prepare organizes the cells in the grid into regions and columns for easy access, and stores them in the Grid struct.
func (g *Grid) prepare() {
	regions := make([][]*Cell, 9)
	cols := make([][]*Cell, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := g.Rows[i][j]
			if regions[cell.Region()-1] == nil {
				regions[cell.Region()-1] = make([]*Cell, 0)
			}
			regions[cell.Region()-1] = append(regions[cell.Region()-1], cell)
			if cols[cell.Col-1] == nil {
				cols[cell.Col-1] = make([]*Cell, 0)
			}
			cols[cell.Col-1] = append(cols[cell.Col-1], cell)
		}
	}
	g.regions = regions
	g.cols = cols
}

// Copy creates a deep copy of the Grid, duplicating its Rows and all contained Cells. Returns a pointer to the new Grid.
func (g *Grid) Copy() *Grid {
	result := &Grid{Rows: make([][]*Cell, 9)}
	for i := 0; i < 9; i++ {
		result.Rows[i] = make([]*Cell, 9)
		for j := 0; j < 9; j++ {
			result.Rows[i][j] = g.Rows[i][j].Copy()
		}
	}
	result.prepare()
	return result
}

// IsValid checks if the Sudoku grid is valid by ensuring no duplicate non-zero values exist in any row, column, or region.
func (g *Grid) IsValid() bool {
	check := func(items []*Cell) bool {
		exists := make(map[int]bool)
		for _, cell := range items {
			if cell.Value == 0 {
				continue
			}
			if exists[cell.Value] {
				return false
			}
			exists[cell.Value] = true
		}
		return true
	}
	for i := 0; i < 9; i++ {
		if !check(g.Rows[i]) || !check(g.cols[i]) || !check(g.regions[i]) {
			return false
		}
	}
	return true
}

// IsComplete checks if all cells in the 9x9 Sudoku grid have non-zero values. Returns true if complete, otherwise false.
func (g *Grid) IsComplete() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if g.Rows[i][j].Value == 0 {
				return false
			}
		}
	}
	return true
}

func (g *Grid) AvailableValues(row, col int) []int {
	cell := g.Rows[row-1][col-1]
	if cell.IsInitialized {
		return []int{cell.Value}
	}
	values := make([]bool, 9)
	readValues := func(items []*Cell) {
		for _, c := range items {
			if c.Value > 0 {
				values[c.Value-1] = true
			}
		}
	}
	for i := 0; i < 9; i++ {
		readValues(g.Rows[cell.Row-1])
		readValues(g.cols[cell.Col-1])
		readValues(g.regions[cell.Region()-1])
	}
	result := make([]int, 0)
	for idx, value := range values {
		if !value {
			result = append(result, idx+1)
		}
	}
	return result
}

func (g *Grid) UpdateValue(value, row, col int) {
	cell := g.Rows[row-1][col-1]
	if !cell.IsInitialized {
		cell.Value = value
	}
}

func (g *Grid) String() string {
	rows := make([]string, 9)
	for i := 0; i < 9; i++ {
		row := make([]string, 9)
		for j := 0; j < 9; j++ {
			cell := g.Rows[i][j]
			if cell.Value == 0 {
				row[j] = "-"
			} else {
				row[j] = strconv.Itoa(cell.Value)
			}
		}
		rows[i] = strings.Join(row, "")
	}
	return strings.Join(rows, "\n")
}
