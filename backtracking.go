package main

import "fmt"

//import "fmt"

var numCalls int
var numPrunes int

type createCandidatesfunc func(a []interface{}, k int, data []interface{}) []interface{}

type isaSolutionfunc func(a []interface{}, k int, data []interface{}) bool

type processSolutionfunc func(a []interface{}, data []interface{})

type pruneSearchfunc func(a []interface{}, k int, data []interface{}) bool

type makeMovefunc func(a []interface{}, k int, data []interface{}, candidate interface{})

type unmakeMovefunc func(a []interface{}, k int, data []interface{}, candidate interface{})

func printIndent(indent int) {
	for i := 1; i <= indent; i++ {
		fmt.Print(" ")
	}
}

func backtrack(a []interface{}, k int, data []interface{},
	createCandidates createCandidatesfunc,
	isasolution isaSolutionfunc,
	processSolution processSolutionfunc,
	prunesearch pruneSearchfunc,
	makemove makeMovefunc,
	unmakemove unmakeMovefunc,
	stopEarly bool,
) bool {

	//printIndent(len(a))
	//fmt.Println("k=", k, ",a=", a)

	if isasolution(a, k, data) {

		fmt.Println("total calls:", numCalls)
		fmt.Println("total prunes:", numPrunes)
		processSolution(a, data)
		return true
	}

	if prunesearch(a, k, data) {

		return false
	}

	k = k + 1

	candidates := createCandidates(a, k, data)
	for _, candidate := range candidates {
		//fmt.Println("k=", k, "candidate=", candidate, "a=", a)
		makemove(a, k, data, candidate)
		if backtrack(append(a, candidate),
			k,
			data,
			createCandidates,
			isasolution,
			processSolution,
			prunesearch,
			makemove, unmakemove, stopEarly) {
			if stopEarly {
				return true
			}

		}
		unmakemove(a, k, data, candidate)

	}
	return false

}
