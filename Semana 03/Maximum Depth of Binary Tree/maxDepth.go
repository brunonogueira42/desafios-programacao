package main

import (
	"fmt"
)

// Definição da estrutura do nó da árvore
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Função para calcular a profundidade máxima
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// Função de exemplo para montar a árvore do exemplo 1: [3,9,20,null,null,15,7]
func buildExampleTree() *TreeNode {
	return &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
}

func main() {
	root := buildExampleTree()
	fmt.Println("Profundidade máxima:", maxDepth(root)) // Esperado: 3
}
