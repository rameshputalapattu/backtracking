package main

import (
	"fmt"
)

func isasolutionSumPart(a []interface{}, k int, data []interface{}) bool {
	var total = 0
	for _, elem := range a {
		total += elem.(int)
	}
	return total == 19
}

func prunesearchSumPart(a []interface{}, k int, data []interface{}) bool {
	var total = 0
	for _, elem := range a {
		total += elem.(int)
		if total > 19 {
			return true
		}
	}
	return false
}

func processSolutionSumPart(a []interface{}, data []interface{}) {

	cost := 0
	for _, elem := range a {
		if elem.(int) == 3 {
			cost += 10
		}

		if elem.(int) == 4 {
			cost += 15
		}
	}

	fmt.Println("soln=", a, "cost=", cost)

}

func createCandidatesSumPart(a []interface{}, k int, data []interface{}) []interface{} {
	candidates := []interface{}{3, 4}
	return candidates

}
