package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Board represents a sudoko board
type Board struct {
	Squares   [][]int
	FreeCells int
}

type move struct {
	x     int
	y     int
	value int
}

func (b *Board) setFreeCells() {
	for i := 0; i < len(b.Squares); i++ {
		for j := 0; j < len(b.Squares[i]); j++ {
			if b.Squares[i][j] == 0 {
				b.FreeCells++
			}
		}
	}
}

func computeminXminYVals(b Board) (int, int, int) {
	var minX, minY int
	numPosVals := 9
	for i := 0; i < len(b.Squares); i++ {
		for j := 0; j < len(b.Squares[i]); j++ {
			if b.Squares[i][j] == 0 {
				currentPosVals := len(possibleValues(&b, i, j))
				if currentPosVals < numPosVals {
					numPosVals = currentPosVals
					minX = i
					minY = j
				}
			}
		}
	}
	return minX, minY, numPosVals

}

func (b Board) nextFreeSquare() (int, int, int) {
	// for i := 0; i < len(b.Squares); i++ {
	// 	for j := 0; j < len(b.Squares[i]); j++ {
	// 		if b.Squares[i][j] == 0 {
	// 			return i, j
	// 		}
	// 	}
	// }
	//return 0, 0
	minX, minY, minMoves := computeminXminYVals(b)
	return minX, minY, minMoves

}

func (b Board) printBoard() {
	fmt.Println("Free cells:", b.FreeCells)
	for _, row := range b.Squares {
		fmt.Println(row)
	}
}

func makeMove(b *Board, a []interface{}, k int, data []interface{}, candidate interface{}) {
	currentMove := candidate.(move)
	b.Squares[currentMove.x][currentMove.y] = currentMove.value
	b.FreeCells--
	//b.printBoard()

}

func unmakeMove(b *Board, a []interface{}, k int, data []interface{}, candidate interface{}) {
	currentMove := candidate.(move)
	b.Squares[currentMove.x][currentMove.y] = 0
	b.FreeCells++
}

func pruneSearchSudoku(b *Board, a []interface{}, k int, data []interface{}) bool {
	// _, _, minPosVals := computeminXminYVals(*b)
	// if minPosVals == 0 {
	// 	numPrunes++
	// 	return true
	// }
	return false
	//return b.FreeCells == 0
}

func processSolutionSudoku(b *Board, a []interface{}, data []interface{}) {
	b.printBoard()
}

func isasolutionSudoku(b *Board, a []interface{}, k int, data []interface{}) bool {
	numCalls++
	return b.FreeCells == 0
}

func possibleValues(b *Board, x, y int) []interface{} {
	var vals []int
	var pvals []interface{}
	rowSet := make(map[int]struct{})
	colSet := make(map[int]struct{})
	sectorSet := make(map[int]struct{})

	for i := 0; i < 9; i++ {
		if i == y {
			continue
		}
		val := b.Squares[x][i]
		if val != 0 {
			rowSet[val] = struct{}{}
		}

	}

	for i := 0; i < 9; i++ {
		if i == x {
			continue
		}
		val := b.Squares[i][y]
		if val != 0 {
			colSet[val] = struct{}{}
		}

	}

	sectorXtopLeft := (x / 3) * 3
	sectorYtopLeft := (y / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val := b.Squares[i+sectorXtopLeft][j+sectorYtopLeft]
			if val != 0 {
				sectorSet[val] = struct{}{}
			}
		}

	}

	for i := 1; i <= 9; i++ {
		_, inRow := rowSet[i]
		_, inCol := colSet[i]
		_, inSector := sectorSet[i]

		if !inRow && !inCol && !inSector {
			vals = append(vals, i)
		}
	}

	for _, val := range vals {
		pvals = append(pvals, move{x, y, val})
	}

	return pvals
}

func createCandidatesSudoku(b *Board, a []interface{}, k int, data []interface{}) []interface{} {

	x, y, minMoves := b.nextFreeSquare()
	if minMoves == 0 {
		return nil
	}

	return possibleValues(b, x, y)
}

func parseBoard(puzzle string) (*Board, error) {

	puzzleCleaned := strings.ReplaceAll(puzzle, "\n", "")
	puzzleCleaned = strings.Trim(puzzleCleaned, "\n")
	puzzleCleaned = strings.ReplaceAll(puzzleCleaned, "\t", "")
	boardEntriesStrArr := strings.Split(puzzleCleaned, "")
	boardEntries := make([][]int, 9)
	for i := 0; i < len(boardEntries); i++ {
		boardEntries[i] = make([]int, 9)
		for j := 0; j < len(boardEntries[i]); j++ {
			entry, err := strconv.Atoi(boardEntriesStrArr[i*9+j])
			if err != nil {
				return nil, err
			}
			boardEntries[i][j] = entry

		}
	}

	board := Board{Squares: boardEntries}
	board.setFreeCells()
	return &board, nil

}
