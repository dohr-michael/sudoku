package sudoku

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestResolve(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name       string
		input      string
		wantResult string
		wantErr    error
	}{
		{
			name:    "Empty string",
			input:   "",
			wantErr: errors.New("invalid number of rows: 0"),
		},
		{
			name:    "Invalid grid",
			input:   "3365-84-- 52------- -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--",
			wantErr: errors.New("invalid sudoku grid"),
		},
		{
			name:       "Valid grid",
			input:      "534678912 672195348 198342567 859761423 426853791 713924856 961537284 287419635 345286179",
			wantResult: "534678912 672195348 198342567 859761423 426853791 713924856 961537284 287419635 345286179",
			wantErr:    nil,
		},
		{
			name:       "Valid grid 2",
			input:      "3-65-84-- 52------- -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--",
			wantResult: "316578492 529134768 487629531 263415987 974863125 851792643 138947256 692351874 745286319",
			wantErr:    nil,
		},
		/*{
			name:    "Unsolvable grid",
			input:   "3365-84-- 52------- -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--", // random unsolvable grid
			wantErr: errors.New("could not resolve sudoku grid"),
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGrid, err := Resolve(tt.input)
			if tt.wantErr == nil {
				assert.NoError(err)
				assert.NotNil(gotGrid)
				assert.Equal(tt.wantResult, strings.ReplaceAll(gotGrid.String(), "\n", " "))
			} else {
				assert.Error(err)
				assert.Equal(tt.wantErr.Error(), err.Error())
			}
		})
	}
}

// 7,085131
func BenchmarkResolve(b *testing.B) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Valid grid",
			input: "3-65-84-- 52------- -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--",
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Resolve(tt.input)
			}
		})
	}
}
