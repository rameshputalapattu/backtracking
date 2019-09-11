package main

import "sort"

func solveSubsetSum() {
	data := []interface{}{-7, -3, -2, 5, 8}
	sort.Slice(data, func(i, j int) bool {
		return data[i].(int) < data[j].(int)
	})
	a := make([]interface{}, 0)
	targetSum := 0
	isSolutionfuncSS := func(a []interface{}, k int, data []interface{}) bool {
		return isasolutionSubSetSum(a, k, data, targetSum)
	}
	pruneSearchSS := func(a []interface{}, k int, data []interface{}) bool {
		return prunesearchSubSetSum(a, k, data, targetSum)
	}
	backtrack(a, 0, data,
		createCandidatesSubSetSum,
		isSolutionfuncSS,
		processSolutionSubSetSum,
		pruneSearchSS,
		makeMoveSubsetSum,
		unmakeMoveSubsetSum, false)

}
