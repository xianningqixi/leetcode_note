package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	s := make([]int, 0)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return res
}
