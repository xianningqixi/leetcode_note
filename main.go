package main

import (
	"fmt"
)

func main() {
	matrix := [][]string{{"1", "0", "1", "0", "0"}, {"1", "0", "1", "1", "1"}, {"1", "1", "1", "1", "1"}, {"1", "0", "0", "1", "0"}}
	result := maximalSquare(matrix)
	fmt.Printf("Hello and welcome, %+v!", result)
	return
}

func maximalSquare(matrix [][]string) int {
	x := len(matrix)
	y := len(matrix[0])
	n := min(x, y)
	result := 0
	for size := 1; size < n+1; size++ {
		check := false
		for i := 0; i < n-size+1; i++ {
			for j := 0; j < n-size+1; j++ {
				countCheck := true
				for count := 0; count <= size; count++ {
					if matrix[i+count][j+count] != "1" {
						countCheck = false
						break
					}
				}
				check = countCheck
				if countCheck {
					break
				}
			}
			if check {
				break
			}
		}
		if check {
			result = size * size
		} else {
			break
		}
	}

	return result
}
