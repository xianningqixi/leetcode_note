package main

import (
	"fmt"
	"strconv"
)

/*
func main() {
	nums := [][]int{{1, 2, 3}, {4, 0, 5}}
	result := slidingPuzzle(nums)
	fmt.Printf("Hello and welcome, %+v!", result)
	return
}
*/

func slidingPuzzle(board [][]int) int {
	target := "123450"
	start := ""
	//转换图为二叉树
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			start = start + strconv.Itoa(board[i][j])
		}
	}
	q := make([]string, 0)
	q = append(q, start)
	visited := make(map[string]bool)
	visited[start] = true
	step := 0
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			cur := q[0]
			q = q[1:]
			if cur == target {
				return step
			}
			for _, node := range getNeighbor(cur) {
				if !visited[node] {
					visited[node] = true
					q = append(q, node)
				}
			}
		}
		step++
	}
	return -1
}

func getNeighbor(cur string) []string {
	res := make([]string, 0)
	neighBorMap := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}
	for idx, value := range cur {
		if value == '0' {
			neighBors := neighBorMap[idx]
			for _, neighBor := range neighBors {
				res = append(res, swap(cur, idx, neighBor))
			}
		}
	}
	return res
}

func swap(choice string, x, y int) string {
	runeStr := []rune(choice)
	runeStr[x], runeStr[y] = runeStr[y], runeStr[x]
	return string(runeStr)
}
