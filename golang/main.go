package main

import (
	"fmt"
	"github.com/dohr-michael/sudoku/sudoku"
	"github.com/pterm/pterm"
	"log"
	"time"
)

func printGrid(area *pterm.AreaPrinter, grid *sudoku.Grid) {
	// Create a default box with specified padding
	// paddedBox := pterm.DefaultBox //.WithLeftPadding(1).WithRightPadding(1).WithTopPadding(1).WithBottomPadding(1)
	panels := make([][]pterm.Panel, 9)
	for i, row := range grid.Rows {
		panels[i] = make([]pterm.Panel, 9)
		for j, cell := range row {
			value := fmt.Sprintf("%d", cell.Value)
			if cell.Value == 0 {
				value = "-"
			}
			if cell.IsInitialized {
				value = pterm.LightRed(value)
			}

			panels[i][j] = pterm.Panel{Data: value}
		}
	}
	content, _ := pterm.DefaultPanel.WithPadding(0).WithBottomPadding(0).WithSameColumnWidth(true).WithPanels(panels).Srender()
	area.Update(content)
	time.Sleep(50 * time.Millisecond)
}

func main() {
	area, _ := pterm.DefaultArea. /*.WithFullscreen().WithCenter()*/ Start()
	defer area.Stop()
	result, err := sudoku.Resolve("3-65-84-- 52------- -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--",
		func(grid *sudoku.Grid) {
			//printGrid(area, grid)
		})
	if err != nil {
		log.Fatal(err)
	}
	printGrid(area, result)
}
