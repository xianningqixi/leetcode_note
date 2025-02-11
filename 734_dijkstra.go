package main

import (
	"math"
)

/*
func main() {
	var times [][]int
	times = [][]int{
		{2, 1, 1}, {2, 3, 1}, {3, 4, 1},
	}
	result := networkDelayTime(times, 4, 2)
	fmt.Printf("Hello and welcome, %d!", result)
	return
}
*/

var inf = math.MaxInt / 2

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]int, n)

	for idx, _ := range graph {
		graph[idx] = make([]int, n)
		for j := range graph[idx] {
			graph[idx][j] = inf
		}
	}
	for _, value := range times {
		graph[value[0]-1][value[1]-1] = value[2]
	}
	dis, _ := dijkstra(graph, k-1)
	maxTime := 0
	for idx, value := range dis {
		if value == inf && idx != k-1 {
			return -1
		}
		if value > maxTime && value != inf {
			maxTime = value
		}
	}
	return maxTime
}

func dijkstra(graph [][]int, start int) ([]int, []int) {
	startList := make([]int, len(graph))
	dis := make([]int, len(graph))
	path := make([]int, len(graph))
	// 初始化
	for idx, value := range graph {
		if idx == start {
			dis = value
			break
		}
	}
	for idt, value := range dis {
		if value != inf {
			path[idt] = start
		} else {
			path[idt] = -1
		}
	}

	startList[start] = 1
	count := 1
	for count < len(graph) {
		nextNode := getNextNode(startList, dis)
		startList[nextNode] = 1
		for idy, value := range graph[nextNode] {
			if startList[idy] == 0 && (dis[idy] > dis[nextNode]+value) {
				dis[idy] = dis[nextNode] + value
				path[idy] = nextNode
			}
		}
		count++
	}
	return dis, path
}

func getNextNode(startList []int, dis []int) int {
	minValue := inf + 1
	minTag := -1
	for idx, value := range dis {
		if value < minValue && startList[idx] == 0 {
			minValue = value
			minTag = idx
		}
	}
	return minTag
}
