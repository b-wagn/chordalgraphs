package main

import (
	"github.com/yourbasic/graph"
	"math/rand"
)

//returns a graph G(n,p), randomly sampled in the Erdoes-Renyi model with n vertices,
//i.e. every edge is taken with probability p.
func RandomGraph(n int, p float32, seed int64) *graph.Mutable {
	rand.Seed(seed)
	g := graph.New(n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if rand.Float32() <= p {
				g.AddBoth(i, j)
			}
		}
	}
	return g
}
