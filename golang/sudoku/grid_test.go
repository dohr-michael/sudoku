package sudoku

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var validGrid = "3-65-84-- 52------- -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--"
var invalidGrid1 = "3365-84-- 52------- -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--"
var invalidGrid2 = "3-65-84-- 52------- 387----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--"
var invalidGrid3 = "3-65-84-- 523------ -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--"

func grid(input string) (res *Grid) {
	res, err := NewGrid(input)
	if err != nil {
		println(err.Error())
	}
	return
}

func TestNewGrid(t *testing.T) {
	generateRow := func(input string, row int) []*Cell {
		result := make([]*Cell, 9)
		for i := 0; i < 9; i++ {
			value := 0
			if input[i] != '-' {
				value, _ = strconv.Atoi(string(input[i]))
			}
			result[i] = &Cell{
				value,
				value != 0,
				row,
				i + 1,
			}
		}
		return result
	}

	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantGrid *Grid
	}{
		{
			name:    "valid",
			input:   "123456789 123456789 123456789 --3456789 123456789 123456789 123456789 123456789 123456789",
			wantErr: false,
			wantGrid: &Grid{
				Rows: [][]*Cell{
					generateRow("123456789", 1),
					generateRow("123456789", 2),
					generateRow("123456789", 3),
					generateRow("--3456789", 4),
					generateRow("123456789", 5),
					generateRow("123456789", 6),
					generateRow("123456789", 7),
					generateRow("123456789", 8),
					generateRow("123456789", 9),
				},
			},
		},
		{
			name:    "invalid_number_of_rows",
			input:   "123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789",
			wantErr: true,
		},
		{
			name:    "invalid_number_of_cells",
			input:   "12345678 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789",
			wantErr: true,
		},
		{
			name:    "invalid_cell_value",
			input:   "123456789 12a456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := NewGrid(tt.input)
			if tt.wantErr {
				require.NotNil(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantGrid.Rows, res.Rows)
			}

			if err != nil {
				assert.Equal(t, tt.wantErr, true)
			}
		})
	}
}

func TestCellRegion(t *testing.T) {
	region := func(sx, sy int) []Cell {
		result := []Cell{
			{Row: sx, Col: sy},
			{Row: sx, Col: sy + 1},
			{Row: sx, Col: sy + 2},
			{Row: sx + 1, Col: sy},
			{Row: sx + 1, Col: sy + 1},
			{Row: sx + 1, Col: sy + 2},
			{Row: sx + 2, Col: sy},
			{Row: sx + 2, Col: sy + 1},
			{Row: sx + 2, Col: sy + 2},
		}

		return result
	}
	tests := []struct {
		name       string
		items      []Cell
		wantRegion int
	}{
		{
			name:       "1",
			items:      region(1, 1),
			wantRegion: 1,
		},
		{
			name:       "2",
			items:      region(1, 4),
			wantRegion: 2,
		},
		{
			name:       "3",
			items:      region(1, 7),
			wantRegion: 3,
		},
		{
			name:       "4",
			items:      region(4, 1),
			wantRegion: 4,
		},
		{
			name:       "5",
			items:      region(4, 4),
			wantRegion: 5,
		},
		{
			name:       "6",
			items:      region(4, 7),
			wantRegion: 6,
		},
		{
			name:       "7",
			items:      region(7, 1),
			wantRegion: 7,
		},
		{
			name:       "8",
			items:      region(7, 4),
			wantRegion: 8,
		},
		{
			name:       "9",
			items:      region(7, 7),
			wantRegion: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//res := tt.items.Region()
			for idx, i := range tt.items {
				assert.Equal(t, tt.wantRegion, i.Region(), fmt.Sprintf("idx: %d (r: %d,c: %d)", idx, i.Row, i.Col))
			}

		})
	}
}

func TestGridIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    *Grid
		wantBool bool
	}{
		{
			name:     "valid",
			input:    grid(validGrid),
			wantBool: true,
		},
		{
			name:     "invalid1",
			input:    grid(invalidGrid1),
			wantBool: false,
		},
		{
			name:     "invalid2",
			input:    grid(invalidGrid2),
			wantBool: false,
		},
		{
			name:     "invalid3",
			input:    grid(invalidGrid3),
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.input.IsValid()
			assert.Equal(t, tt.wantBool, output)
		})
	}
}

func BenchmarkGridIsValid(b *testing.B) {
	grids := []*Grid{
		grid(validGrid),
		grid(invalidGrid1),
	}

	for _, g := range grids {
		b.Run(fmt.Sprintf("Grid %p", g), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				g.IsValid()
			}
		})
	}
}

func TestGridCopy(t *testing.T) {
	tests := []struct {
		name     string
		input    *Grid
		expected *Grid
	}{
		{
			name:     "copy_valid_grid",
			input:    grid(validGrid),
			expected: grid(validGrid),
		},
		{
			name:     "copy_invalid_grid1",
			input:    grid(invalidGrid1),
			expected: grid(invalidGrid1),
		},
		{
			name:     "copy_invalid_grid2",
			input:    grid(invalidGrid2),
			expected: grid(invalidGrid2),
		},
		{
			name:     "copy_invalid_grid3",
			input:    grid(invalidGrid3),
			expected: grid(invalidGrid3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Copy()
			assert.EqualExportedValues(t, tt.expected, result)
		})
	}
}

func TestIsComplete(t *testing.T) {
	tests := []struct {
		name       string
		input      *Grid
		wantResult bool
	}{
		{
			name:       "complete_grid",
			input:      grid("123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789"),
			wantResult: true,
		},
		{
			name:       "incomplete_grid",
			input:      grid("123456789 123456789 123456789 ---456789 123456789 123456789 123456789 123456789 123456789"),
			wantResult: false,
		},
		{
			name:       "empty_grid",
			input:      grid("--------- --------- --------- --------- --------- --------- --------- --------- ---------"),
			wantResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsComplete()
			assert.Equal(t, tt.wantResult, result)
		})
	}
}

func TestAvailableValues(t *testing.T) {
	tests := []struct {
		name       string
		grid       *Grid
		row        int
		col        int
		wantValues []int
	}{
		{
			name:       "initialized_cell",
			grid:       grid(validGrid),
			row:        1,
			col:        1,
			wantValues: []int{3},
		},
		{
			name:       "uninitialized_cell",
			grid:       grid(validGrid),
			row:        1,
			col:        2,
			wantValues: []int{1, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValues := tt.grid.AvailableValues(tt.row, tt.col)
			sort.Ints(gotValues)
			sort.Ints(tt.wantValues)
			assert.Equal(t, tt.wantValues, gotValues)
		})
	}
}
func TestGridUpdateValue(t *testing.T) {
	tests := []struct {
		name     string
		grid     *Grid
		value    int
		row      int
		col      int
		expected string
	}{
		{
			name:     "update_uninitialized_cell",
			grid:     grid("--------- --------- --------- --------- --------- --------- --------- --------- ---------"),
			value:    5,
			row:      5,
			col:      5,
			expected: "--------- --------- --------- --------- ----5---- --------- --------- --------- ---------",
		},
		{
			name:     "update_initialized_cell",
			grid:     grid("--------- --------- --------- ----5---- --------- --------- --------- --------- ---------"),
			value:    9,
			row:      5,
			col:      5,
			expected: "--------- --------- --------- ----5---- ----9---- --------- --------- --------- ---------",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.grid.UpdateValue(tt.value, tt.row, tt.col)
			assert.Equal(t, tt.expected, strings.ReplaceAll(tt.grid.String(), "\n", " "))
		})
	}
}
