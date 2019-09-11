package main

import "log"

func solveSudokuPuzzles() {
	sudoku_1 := `
	000000012
	000035000
	000600070
	700000300
	000400800
	100000000
	000120000
	080000040
	050000600

	`

	sudoku_2 := `
	000600420
	000000000
	005049706
	840005100
	300080009
	006300075
	509130600
	000000000
	018002000

	`

	sudoku_3 := `
	530070000
	600195000
	098000060
	800060003
	400803001
	700020006
	060000280
	000419005
	000080079
	`

	sudoku_4 := `
	000000060
	730062500
	015800009
	090080006
	001506400
	300070080
	500008620
	007250031
	020000000

	`
	solveSudoku(sudoku_1)
	solveSudoku(sudoku_2)
	solveSudoku(sudoku_3)
	solveSudoku(sudoku_4)
}

func solveSudoku(sudoku string) {
	numCalls = 0
	numPrunes = 0
	board, err := parseBoard(sudoku)
	if err != nil {
		log.Fatal(err)
	}

	board.printBoard()
	isSolutionSku := func(a []interface{}, k int, data []interface{}) bool {
		return isasolutionSudoku(board, a, k, data)
	}

	processSolutionSku := func(a []interface{}, data []interface{}) {
		processSolutionSudoku(board, a, data)
	}

	createCandidatesSku := func(a []interface{}, k int, data []interface{}) []interface{} {
		return createCandidatesSudoku(board, a, k, data)
	}

	unmakeMoveSku := func(a []interface{}, k int, data []interface{}, candidate interface{}) {
		unmakeMove(board, a, k, data, candidate)
	}

	makeMoveSku := func(a []interface{}, k int, data []interface{}, candidate interface{}) {
		makeMove(board, a, k, data, candidate)
	}

	pruneSearchSku := func(a []interface{}, k int, data []interface{}) bool {
		return pruneSearchSudoku(board, a, k, data)

	}

	backtrack(nil, 0, nil, createCandidatesSku, isSolutionSku, processSolutionSku,
		pruneSearchSku, makeMoveSku, unmakeMoveSku, true)
}
