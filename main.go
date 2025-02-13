package main

import (
	"fmt"
)

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	result := openLock(deadends, target)
	fmt.Printf("Hello and welcome, %+v!", result)
	return
}
