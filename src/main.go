package main

import (
	"os"
	"fmt"
	"strconv"
	"time"
)

func runAlgorithmsTest(seed int64, n int, numberofinstances int) (float64,float64,float64) {
	var av_clique int64 = 0
	var av_coloring int64 = 0
	var av_ccindset int64 = 0
	for i := 0; i < numberofinstances; i++ {
		graph := RandomIntervalGraph(n,20*n,seed+int64(i))
		order := LexBFS(graph)
		cg := ChordalGraphFromGraphAndPES(graph, order)
		start_clique := time.Now()
		Clique(cg)
		elapsed_clique := time.Since(start_clique)
		av_clique += elapsed_clique.Nanoseconds()

		start_coloring := time.Now()
		Coloring(cg)
		elapsed_coloring := time.Since(start_coloring)
		av_coloring+= elapsed_coloring.Nanoseconds()

		start_ccindset := time.Now()
		CliqueCoverAndIndSet(cg)
		elapsed_ccindset := time.Since(start_ccindset)
		av_ccindset += elapsed_ccindset.Nanoseconds()

	}
	avg_clique := float64(av_clique)/(1000.0 * float64(numberofinstances)) //return in millisecs
	avg_coloring := float64(av_coloring)/(1000.0 * float64(numberofinstances))
	avg_ccindset := float64(av_ccindset)/(1000.0 * float64(numberofinstances))
	return avg_clique,avg_coloring,avg_ccindset
}



func runLexBFSTest(seed int64, n int, numberofinstances int) float64 {
	var p float32 = 0.1
	var av int64 = 0
	for i := 0; i < numberofinstances; i++ {
		graph := RandomGraph(n, p, seed+int64(i))
		start := time.Now()
		LexBFS(graph)
		elapsed := time.Since(start)
		av += elapsed.Nanoseconds()
	}
	avg := float64(av)/(1000.0 * float64(numberofinstances)) //return in millisecs
	return avg
}

func runValidityCheckTest(seed int64, n int, numberofinstances int) float64 {
	var p float32 = 0.1
	var av int64 = 0
	for i := 0; i < numberofinstances; i++ {
		graph := RandomGraph(n, p, seed+int64(i))
		order := LexBFS(graph)
		cg := ChordalGraphFromGraphAndPES(graph, order)
		start := time.Now()
		IsValid(cg)
		elapsed := time.Since(start)
		av += elapsed.Nanoseconds()
	}
	avg := float64(av)/(1000.0 * float64(numberofinstances)) //return in millisecs
	return avg
}

func runRecognitionTest(seed int64, n int, numberofinstances int) float64 {
	var p float32 = 0.1
	var av int64 = 0
	for i := 0; i < numberofinstances; i++ {
		graph := RandomGraph(n, p, seed+int64(i))
		start := time.Now()
		order := LexBFS(graph)
		cg := ChordalGraphFromGraphAndPES(graph, order)
		IsValid(cg)
		elapsed := time.Since(start)
		av += elapsed.Nanoseconds()
	}
	avg := float64(av)/(1000.0 * float64(numberofinstances)) //return in millisecs
	return avg
}

func main() {
	
	argc := len(os.Args[1:])
	if argc < 3 {
		fmt.Println("Usage: <seed> <numberofvertices> <numberofinstances>")
		return
	}

	seed, _ := strconv.ParseInt(os.Args[1], 10, 0)
	n, _ := strconv.Atoi(os.Args[2])
	numberofinstances, _ := strconv.Atoi(os.Args[3])

	fmt.Println(runRecognitionTest(seed, n, numberofinstances))
	fmt.Println(runLexBFSTest(seed, n, numberofinstances))
	fmt.Println(runValidityCheckTest(seed, n, numberofinstances))
	res1,res2,res3 := runAlgorithmsTest(seed, n, numberofinstances)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}
