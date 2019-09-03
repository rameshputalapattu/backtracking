package main

import (
	"fmt"
)

func isasolutionperm(a []interface{}, k int, data []interface{}) bool {
	return len(a) == len(data)
}

func processSolutionperm(a []interface{}, data []interface{}) {
	soln := make([]interface{}, 0, len(data))

	for _, elem := range a {
		soln = append(soln, data[elem.(int)])
	}
	fmt.Println(soln)

}

func prunesearchperm(a []interface{}, k int, data []interface{}) bool {
	return false
}

func createCandidatesPerm(a []interface{}, k int, data []interface{}) []interface{} {
	inPerm := make([]bool, len(data))

	for i := 0; i < k-1; i++ {
		inPerm[a[i].(int)] = true
	}

	candidates := make([]interface{}, 0, len(data))
	for idx, incl := range inPerm {
		if !incl {
			candidates = append(candidates, idx)
		}

	}
	return candidates

}
