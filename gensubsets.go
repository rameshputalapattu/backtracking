package main

func genSubsets() {
	a := make([]interface{}, 0)
	data := []interface{}{1, 2, 3}
	backtrack(a, 0, data, createCandidatesSubsets, isasolutionSubsets, processSolutionSubsets,
		prunesearchSubsets,
		makeMoveSubsets, unmakeMoveSubsets, false)
}
