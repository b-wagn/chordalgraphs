package main

import (
	"github.com/yourbasic/graph"
)

// A struct representing a chordal graph in terms of its perfect elemination scheme,
// its inverse, and a adjacency list of all right neighbors in terms of the pes.
type ChordalGraph struct {
	n       int     //number of vertices
	pes     []int   //perfect elimination scheme:: position --> vertex
	pes_inv []int   //inverse of pes:: vertex --> position
	radj    [][]int //the adjacency lists, restricted to right neighbors
}

//This function checks if the given graph is really chordal, or to be precise, if the
//ordering graph.pes is really a perfect elimination scheme. This means that even for
//a chordal graph it can return false iff the given ordering is not a PES.
//If graph.pes is not a perfect elimination scheme, the second and third return parameter
//will contain a witness pair for that, i.e. a pair of vertices that are not adjecent but should be
//Note: it will not check if graph.pes and graph.pes_inv are really inverses.
//Takes linear time
func IsValid(graph *ChordalGraph) (bool,int,int) {
	//for each vertex v, careset[v] will be the set of vertices that have to be adjacent to it
	careset := make([][]int, graph.n)
	for v := 0; v < graph.n; v++ {
		careset[v] = make([]int, 0)
	}
	flags := make([]bool, graph.n) //needed for set comparison

	//iterate over ordering and build careset sets
	for i := 0; i < graph.n; i++ {
		v := graph.pes[i]
		rightNeighbors := graph.radj[v]

		//if v has any neighbors, we need to insert them into careset[u],
		//where u is the leftmost right neighbor of v
		//of course, we do not need to insert u itself
		if len(rightNeighbors) != 0 {
			u := rightNeighbors[0]
			for _, w := range rightNeighbors {
				if graph.pes_inv[w] < graph.pes_inv[u] {
					u = w
				}
			}
			for _, w := range rightNeighbors {
				if u != w {
					careset[u] = append(careset[u], w)
				}
			}
		}
		//check that careset is a subset of the right neighbors:
		//Precondition: All flags are zero
		//set all neighbors to true
		for _, u := range rightNeighbors {
			flags[u] = true
		}
		//check that all in careset have been set to true
		for _, u := range careset[v] {
			if !flags[u] {
				return false,u,v
			}
		}
		//Set all flags back to false, required for next iteration
		for _, u := range rightNeighbors {
			flags[u] = false
		}
	}
	return true,-1,-1
}

//This function takes a chordal graph graph and a perfect elimination scheme pes of this graph as input.
//It returns a representation of the graph as a ChordalGraph struct
//Takes linear time.
func ChordalGraphFromGraphAndPES(graph graph.Iterator, pes []int) *ChordalGraph {
	n := graph.Order()
	//build inverse
	pes_inv := make([]int, n)
	for i := 0; i < n; i++ {
		pes_inv[pes[i]] = i
	}
	//build right adjacency list
	radj := make([][]int, n)
	for v := 0; v < n; v++ {
		//count how many right neighbors we have
		numRightNeighbors := 0
		graph.Visit(v, func(w int, c int64) (skip bool) {
			if pes_inv[v] < pes_inv[w] {
				numRightNeighbors++
			}
			return
		})
		//make slice and add right neighbors to it
		radj[v] = make([]int, numRightNeighbors)
		i := 0
		graph.Visit(v, func(w int, c int64) (skip bool) {
			if pes_inv[v] < pes_inv[w] {
				radj[v][i] = w
				i++
			}
			return
		})
	}

	//build result
	res := ChordalGraph{
		n:       n,
		pes:     pes,
		pes_inv: pes_inv,
		radj:    radj,
	}
	return &res
}
