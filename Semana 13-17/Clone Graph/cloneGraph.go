package main

import "fmt"

// https://leetcode.com/problems/clone-graph/

// Complexidade Temporal: O(n + m)
// Complexidade Espacial: O(n + m)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	cloned := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if c, ok := cloned[n]; ok {
			return c
		}

		copy := &Node{Val: n.Val}
		cloned[n] = copy

		for _, neighbor := range n.Neighbors {
			copy.Neighbors = append(copy.Neighbors, dfs(neighbor))
		}

		return copy
	}

	return dfs(node)
}

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	cloned := cloneGraph(n1)
	fmt.Printf("Cloned Node 1: %v\n", cloned.Val)
}
