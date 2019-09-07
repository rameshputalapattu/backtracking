package main

import (
	"log"
)

func main() {
	//data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//data := []interface{}{10, 20, 30, 40, 50, 60}
	// data := []interface{}{-7, -3, -2, 5, 8}
	// sort.Slice(data, func(i, j int) bool {
	// 	return data[i].(int) < data[j].(int)
	// })
	// a := make([]interface{}, 0)
	// targetSum := 0
	// isSolutionfuncSS := func(a []interface{}, k int, data []interface{}) bool {
	// 	return isasolutionSubSetSum(a, k, data, targetSum)
	// }
	// pruneSearchSS := func(a []interface{}, k int, data []interface{}) bool {
	// 	return prunesearchSubSetSum(a, k, data, targetSum)
	// }

	// //backtrack(a, 0, data, createCandidatesPerm, isasolutionperm, processSolutionperm,prunesearchperm)
	// //backtrack(a, 0, data, createCandidatesSubsets, isasolutionSubsets, processSolutionSubsets,prunesearchSubsets)
	// //backtrack(a, 0, nil, createCandidatesSumPart, isasolutionSumPart, processSolutionSumPart, prunesearchSumPart)
	// backtrack(a, 0, data, createCandidatesSubSetSum, isSolutionfuncSS, processSolutionSubSetSum, pruneSearchSS)

	sudoku := `
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
		pruneSearchSku, makeMoveSku, unmakeMoveSku)

}
