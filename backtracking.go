package main

//import "fmt"

type createCandidatesfunc func(a []interface{}, k int, data []interface{}) []interface{}

type isaSolutionfunc func(a []interface{}, k int, data []interface{}) bool

type processSolutionfunc func(a []interface{}, data []interface{})

type pruneSearchfunc func(a []interface{}, k int, data []interface{}) bool

func backtrack(a []interface{}, k int, data []interface{},
	createCandidates createCandidatesfunc,
	isasolution isaSolutionfunc,
	processSolution processSolutionfunc,
	prunesearch pruneSearchfunc,
) {
	if isasolution(a, k, data) {
		processSolution(a, data)
		return
	}

	k = k + 1

	if prunesearch(a, k, data) {
		return
	}

	candidates := createCandidates(a, k, data)
	for _, candidate := range candidates {
		//fmt.Println("k=", k, "candidate=", candidate, "a=", a)
		backtrack(append(a, candidate), k, data, createCandidates, isasolution, processSolution, prunesearch)
	}

}
