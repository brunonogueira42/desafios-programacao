package main

import "fmt"

// https://leetcode.com/problems/course-schedule/

// Complexidade Temporal: O(n + e)
// Complexidade Espacial: O(n)

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !canTake(i, prerequisites, visited, onPath) {
			return false
		}
	}

	return true
}

func canTake(course int, prerequisites [][]int, visited []bool, onPath []bool) bool {
	if visited[course] {
		return true
	}

	if onPath[course] {
		return false
	}

	onPath[course] = true

	for i := 0; i < len(prerequisites); i++ {
		if prerequisites[i][0] == course {
			if !canTake(prerequisites[i][1], prerequisites, visited, onPath) {
				return false
			}
		}
	}

	onPath[course] = false
	visited[course] = true

	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))                                         // Esperado: true
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))                                 // Esperado: false
	fmt.Println(canFinish(2, [][]int{{0, 1}}))                                         // Esperado: true
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))                 // Esperado: true
	fmt.Println(canFinish(8, [][]int{{1, 0}, {2, 6}, {1, 7}, {6, 4}, {7, 0}, {0, 5}})) // Esperado: true
}
