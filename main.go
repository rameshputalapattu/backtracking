package main

func main() {
	//data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data := []interface{}{10, 20, 30, 40, 50, 60}
	a := make([]interface{}, 0)
	//backtrack(a, 0, data, createCandidatesPerm, isasolutionperm, processSolutionperm,prunesearchperm)
	//backtrack(a, 0, data, createCandidatesSubsets, isasolutionSubsets, processSolutionSubsets,prunesearchSubsets)
	//backtrack(a, 0, nil, createCandidatesSumPart, isasolutionSumPart, processSolutionSumPart, prunesearchSumPart)
	backtrack(a, 0, data, createCandidatesSubSetSum, isasolutionSubSetSum, processSolutionSubSetSum, prunesearchSubSetSum)
}
