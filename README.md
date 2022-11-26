# chordalgraphs
This repository contains some algorithms for generating, recognizing, and solving problems on chordal graphs in Go.

The following algorithms are contained:
  * Algorithms for random graph generation (file _randomgraph.go_)
  * Algorithms for random interval graph generation. Recall that interval graphs are chordal. (file _randomintervalgraph.go_)
  * A struct to represent chordal graphs by a graph and a perfect elimination scheme (file _chordalgraph.go_)
  * Algorithms for solving problems on chordal graphs. Contains Clique, Clique Cover, Independent Set, Coloring. (file _algorithms.go_)
  * Algorithms for chordal graph recognition, i.e. LexBFS and checking whether a given ordering is a perfect elimination scheme (files _chordalgraph.go_ and _lexbfs.go_)


# Example Usage

First, generate a random interval graph:
```
graph := RandomIntervalGraphComplement(n,20*n,seed)"
```
Then, run a LexBFS on it to determine a (candidate) perfect elimation scheme:
``` 
order := LexBFS(graph)
```  
Next, parse the order into a chordal graph structure. 
Then check whether this is really a perfect elimination scheme (thereby checking whether the graph is chordal):
```
cg := ChordalGraphFromGraphAndPES(graph, order)
IsValid(cg)
``` 
Run an algorithm on it, e.g. determine a max clique:
``` 
v: IsValid(cg)
```  


