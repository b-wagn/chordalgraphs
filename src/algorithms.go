package main

//This function takes a chordal graph as input, which is represented as a ChordalGraph struct.
//It will return the leftmost vertex of the maximum clique. That is, if the function returns v then
//the maximum clique is gr.radj[v] + v
//Takes linear time.
func Clique(graph *ChordalGraph) int {
	currMaxV := 0
	currMaxClique := 0
	n := graph.n
	//iterate over PES from right to left and use greedy strategy
	for i := n - 1; i >= 0; i-- {
		v := graph.pes[i]
		rightNeighbors := graph.radj[v]
		if len(rightNeighbors) > currMaxClique {
			currMaxV = v
			currMaxClique = len(rightNeighbors)
		}
	}
	return currMaxV
}

//This function takes a chordal graph as input, which is represented as a ChordalGraph struct.
//It will compute an minimum coloring of the vertices, i.e. a mapping vertex --> color
//Takes linear time.
func Coloring(graph *ChordalGraph) []int {
	n := graph.n
	coloring := make([]int, n)
	//iterate over PES from right to left and use greedy strategy
	for i := n - 1; i >= 0; i-- {
		v := graph.pes[i]
		rightNeighbors := graph.radj[v]
		used := make([]bool, len(rightNeighbors)) //flags for the colors already used by neighbors
		for _, w := range rightNeighbors {
			if coloring[w] < len(rightNeighbors) {
				used[coloring[w]] = true
			}
		}
		found := false
		for j := 0; j < len(rightNeighbors); j++ {
			if !used[j] {
				found = true
				coloring[v] = j
				break
			}
		}
		if !found {
			coloring[v] = len(rightNeighbors)
		}
	}
	return coloring
}

//This function takes a chordal graph graph as input, which is represented as a ChordalGraph struct.
//It will return a minimum clique cover and a maximum independent set.
//Takes linear time.
func CliqueCoverAndIndSet(graph *ChordalGraph) ([]int, []int) {
	n := graph.n
	coloring := make([]int, n)
	indset := make([]int, 0)
	currColor := 1
	//iterate over PES from left to right and use greedy strategy
	for i := 0; i < n; i++ {
		v := graph.pes[i]
		rightNeighbors := graph.radj[v]
		if coloring[v] == 0 {
			coloring[v] = currColor
			indset = append(indset, v)
			for _, w := range rightNeighbors {
				coloring[w] = currColor
			}
			currColor++
		}
	}
	return coloring, indset
}
