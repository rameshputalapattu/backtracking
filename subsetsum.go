package main

import (
	"fmt"
)

func isasolutionSubSetSum(a []interface{}, k int, data []interface{}) bool {
	var total = 0
	for idx, elem := range a {
		if elem.(bool) {
			total += data[idx].(int)
		}

		if total == 60 {
			return true
		}
	}

	return false
}

func prunesearchSubSetSum(a []interface{}, k int, data []interface{}) bool {
	if len(a) == len(data) {
		return true
	}

	return false
}

func processSolutionSubSetSum(a []interface{}, data []interface{}) {
	soln := make([]interface{}, 0, len(data))

	for idx, elem := range a {
		if elem.(bool) {
			soln = append(soln, data[idx])
		}

	}
	fmt.Println(soln)

}

func createCandidatesSubSetSum(a []interface{}, k int, data []interface{}) []interface{} {
	candidates := []interface{}{true, false}
	return candidates

}
