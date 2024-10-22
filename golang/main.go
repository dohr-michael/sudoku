package main

import (
	"github.com/dohr-michael/sudoku/sudoku"
	"log"
)

func main() {
	result, err := sudoku.Resolve("3-65-84-- 52------- -87----31 --3-1--8- 9--863--5 -5--9-6-- 13----25- -------74 --52-63--")
	if err != nil {
		log.Fatal(err)
	}
	println(result.String())
}
