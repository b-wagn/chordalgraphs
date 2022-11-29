package main

import (
	"math/rand"
	"github.com/yourbasic/graph"
)


// generates a list of n uniformly random intervals 
// (given as a list of int pairs) in the range [0,bound)
func randomIntervals(n int, bound int, seed int64) ([]int,[]int)  {
	rand.Seed(seed)

	left := make([]int,n)
	right := make([]int,n)

	for i := 0; i < n; i++ {
		l := rand.Intn(bound)
		r := rand.Intn(bound-l)+l
		left[i] = l
		right[i] = r
	}
	return left,right
}

func intervalIntersection(al int, ar int, bl int, br int) bool {
	if ar <= bl || br <= al {
		return false
	}
	return true
}


// generates a random interval graph (interval graphs are chordal)
func RandomIntervalGraph(n int, bound int, seed int64) *graph.Mutable {
	left,right := randomIntervals(n, bound, seed)
	g := graph.New(n)

	for u := 0; u < n; u++ {
		for v := u+1; v < n; v++ {
			if intervalIntersection(left[u],right[u],left[v],right[v]) {
				g.AddBoth(u,v)
			}
		}
	}
	return g
}


// generates the complement of a random interval graph
func RandomIntervalGraphComplement(n int, bound int, seed int64) *graph.Mutable {
	left,right := randomIntervals(n, bound, seed)
	g := graph.New(n)

	for u := 0; u < n; u++ {
		for v := u+1; v < n; v++ {
			if !intervalIntersection(left[u],right[u],left[v],right[v]) {
				g.AddBoth(u,v)
			}
		}
	}
	return g
}

