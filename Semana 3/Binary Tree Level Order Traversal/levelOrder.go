package main

import "fmt"

// https://leetcode.com/problems/binary-tree-level-order-traversal/

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(n)

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelValues)
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Exemplo 1: [3,9,20,null,null,15,7]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}

	result1 := levelOrder(root1)
	fmt.Println(result1) // Esperado: [[3], [9, 20], [15, 7]]

	// Exemplo 2: [1]
	root2 := &TreeNode{Val: 1}

	result2 := levelOrder(root2)
	fmt.Println(result2) // Esperado: [[1]]

	// Exemplo 3: []
	var root3 *TreeNode

	result3 := levelOrder(root3)
	fmt.Println(result3) // Esperado: []
}
