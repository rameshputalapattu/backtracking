package main

import (
	"fmt"
)

func isasolutionSubSetSum(a []interface{}, k int, data []interface{}, targetSum int) bool {
	var total = 0
	for idx, elem := range a {
		if elem.(bool) {
			total += data[idx].(int)
		}

	}

	if total == targetSum && any(a) {
		return true
	}

	return false
}

func any(a []interface{}) bool {
	for _, elem := range a {
		if elem.(bool) {
			return true
		}
	}

	return false
}

func prunesearchSubSetSum(a []interface{}, k int, data []interface{}, targetSum int) bool {
	if len(a) == len(data) {
		return true
	}
	var total = 0
	for idx, elem := range a {
		if elem.(bool) {
			total += data[idx].(int)
		}

	}

	if total > targetSum {
		return true
	}

	if total+data[len(a)].(int) > targetSum {
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
