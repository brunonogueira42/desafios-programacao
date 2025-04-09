package main

import "fmt"

// https://leetcode.com/problems/merge-k-sorted-lists/

// Complexidade Temporal:
// Complexidade Espacial:

func mergeKLists(lists []*ListNode) *ListNode {
	var finalList *ListNode = nil

	for i := 0; i < len(lists); i++ {
		node := lists[i]
		for node != nil {
			newNode := &ListNode{Val: node.Val}

			if finalList == nil || newNode.Val < finalList.Val {
				newNode.Next = finalList
				finalList = newNode
			} else {
				actualNode := finalList
				for actualNode.Next != nil && actualNode.Next.Val < newNode.Val {
					actualNode = actualNode.Next
				}
				newNode.Next = actualNode.Next
				actualNode.Next = newNode
			}

			node = node.Next
		}
	}

	return finalList
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Exemplo 1: [[1,4,5],[1,3,4],[2,6]]
	// Esperado: [1,1,2,3,4,4,5,6]
	l1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := &ListNode{2, &ListNode{6, nil}}

	result1 := mergeKLists([]*ListNode{l1, l2, l3})
	for result1 != nil {
		fmt.Print(result1.Val)
		if result1.Next != nil {
			fmt.Print(" -> ")
		}
		result1 = result1.Next
	}
	fmt.Println()

	// Exemplo 2: []
	// Esperado: []
	result2 := mergeKLists([]*ListNode{})
	for result2 != nil {
		fmt.Print(result2.Val)
		if result2.Next != nil {
			fmt.Print(" -> ")
		}
		result2 = result2.Next
	}
	fmt.Println()

	// Exemplo 3: [[]]
	// Esperado: []
	result3 := mergeKLists([]*ListNode{nil})
	for result3 != nil {
		fmt.Print(result3.Val)
		if result3.Next != nil {
			fmt.Print(" -> ")
		}
		result3 = result3.Next
	}
	fmt.Println()
}
