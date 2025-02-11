package main

/*
func main() {
	result := fib(8)
	fmt.Printf("Hello and welcome, %d!", result)
	return
}
*/

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	fib1 := 1
	fib2 := 0
	for i := 2; i <= n; i++ {
		result := fib1 + fib2
		fib2 = fib1
		fib1 = result
	}
	return fib1
}
