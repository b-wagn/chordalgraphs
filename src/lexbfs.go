package main

import (
	"fmt"
	"github.com/yourbasic/graph"
)

type lexBFSSet struct {
	prev *lexBFSSet //lexBFSSets are chained in a doubly linked list
	next *lexBFSSet
	hol  *lexBFSNode //head of list of nodes within that set
	flag bool        //stores if this set was already split up during one lexBFS step
}

type lexBFSNode struct {
	next   *lexBFSNode //nodes inside a set are chained in a doubly linked list
	prev   *lexBFSNode
	set    *lexBFSSet //the set the node belongs to
	vertex int        //the vertex the node represents
}

//creates a new dummy node to init a list
func newDummyLexBFSNode() *lexBFSNode {
	node := &lexBFSNode{
		next:   nil,
		prev:   nil,
		set:    nil,
		vertex: -1,
	}
	node.next = node
	node.prev = node
	return node
}

//creates a new dummy set to init a queue
func newDummyLexBFSSet() *lexBFSSet {
	dummyNode := newDummyLexBFSNode()
	set := &lexBFSSet{
		flag: false,
		next: nil,
		prev: nil,
		hol:  dummyNode,
	}
	set.next = set
	set.prev = set
	set.hol.set = set
	return set
}

//inserts a new node representing the given vertex into the list after the given node.
//After calling this function, the new node is accessible using node.next
func insertNewNodeAfter(node *lexBFSNode, vertex int) {
	newNode := &lexBFSNode{
		next:   node.next,
		prev:   node,
		set:    node.set,
		vertex: vertex,
	}
	node.next.prev = newNode
	node.next = newNode
}

//inserts a new set into the list after the given set.
//After calling this function, the new node is accessible using set.next
func insertNewSetAfter(set *lexBFSSet) {
	dummyNode := newDummyLexBFSNode()
	newSet := &lexBFSSet{
		flag: false,
		next: set.next,
		prev: set,
		hol:  dummyNode,
	}
	newSet.hol.set = newSet
	set.next.prev = newSet
	set.next = newSet
}

//removes the given node from its list
func removeNode(node *lexBFSNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

//removes the given set from the list
func removeSet(set *lexBFSSet) {
	set.prev.next = set.next
	set.next.prev = set.prev
}

//prints the queue used in the lexBFS algorithm.
//for debugging.
func printStructure(queue *lexBFSSet) {
	for set := queue.next; set != queue; set = set.next {
		fmt.Printf("[Set: {%v}", &set)
		for node := set.hol.next; node != set.hol; node = node.next {
			fmt.Printf("<%v>", node.vertex)
		}
		fmt.Printf("]<-->")
	}
}

//lexicographic BFS on the given graph.
//returns the traversal order in reverse
//which is a PES iff graph is chordal
func LexBFS(graph graph.Iterator) []int {
	n := graph.Order()
	order := make([]int, n)
	inserted := make([]bool, n) //flags indicating if a vertex is already inserted into the order

	nodes := make([]*lexBFSNode, n) //we need to go from the vertex id to its node in the list

	//insert all vertices into a single set in queue
	//which stores sets of nodes, simulates the labels
	queue := newDummyLexBFSSet()
	insertNewSetAfter(queue)
	for i := 0; i < n; i++ {
		insertNewNodeAfter(queue.next.hol, i)
		nodes[i] = queue.next.hol.next
	}

	for i := n - 1; i >= 0; i-- {
		//we need this to clean up: In the fixlist we store all sets that may have become empty in this iteration
		fixlist := make([]*lexBFSSet, 0)

		//pop the first vertex from the first set
		v := queue.next.hol.next.vertex
		fixlist = append(fixlist, queue.next)
		removeNode(queue.next.hol.next)

		//add this vertex to the order
		order[i] = v
		inserted[v] = true

		//for each neighbor which is not yet inserted, we need to split sets (i.e. update its label)
		graph.Visit(v, func(w int, c int64) (skip bool) {
			if !inserted[w] {
				nodeW := nodes[w]
				if !nodeW.set.flag {
					//we need to split the set, i.e. insert a new set in front of it
					insertNewSetAfter(nodeW.set.prev)
					nodeW.set.flag = true
					//insert the set into fixlist
					fixlist = append(fixlist, nodeW.set)

				}
				//remove nodeW from its current set
				removeNode(nodeW)
				//add a new node representing w to the new set
				newSet := nodeW.set.prev
				insertNewNodeAfter(newSet.hol, w)
				nodes[w] = newSet.hol.next
			}
			return
		})
		//we need to remove empty sets, which are contained in fixlist
		for _, set := range fixlist {
			set.flag = false
			if set.hol.next == set.hol {
				removeSet(set)
			}
		}
	}
	return order
}
