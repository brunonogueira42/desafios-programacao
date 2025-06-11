package main

import "fmt"

// https://leetcode.com/problems/path-sum/

// Complexidade Temporal:
// Complexidade Espacial:

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Exemplo: root = [5,4,8,11,null,13,4,7,2,null,null,null,1]
	root := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 1},
			},
		},
	}

	targetSum := 22
	fmt.Println(hasPathSum(root, targetSum)) // Output: true
}
