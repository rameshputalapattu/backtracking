package main

import (
	"fmt"
)

func isasolutionSubsets(a []interface{}, k int, data []interface{}) bool {
	return len(a) == len(data)
}

func prunesearchSubsets(a []interface{}, k int, data []interface{}) bool {
	return false
}

func processSolutionSubsets(a []interface{}, data []interface{}) {
	soln := make([]interface{}, 0, len(data))

	for idx, elem := range a {
		if elem.(bool) {
			soln = append(soln, data[idx])
		}

	}
	fmt.Println(soln)

}

func createCandidatesSubsets(a []interface{}, k int, data []interface{}) []interface{} {
	candidates := []interface{}{true, false}
	return candidates

}
