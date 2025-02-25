package main

/*
func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	result := openLock(deadends, target)
	fmt.Printf("Hello and welcome, %+v!", result)
	return
}
*/

func openLock(deadends []string, target string) int {
	deadMap := make(map[string]bool)
	for _, dead := range deadends {
		deadMap[dead] = true
	}
	start := "0000"
	if deadMap[start] {
		return -1
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
			for _, neighbor := range getNeighbors(cur, deadMap) {
				if !visited[neighbor] {
					q = append(q, neighbor)
					visited[neighbor] = true
				}
			}
		}
		step++
	}
	return -1
}

func getNeighbors(s string, deadMap map[string]bool) []string {
	neighborMap := map[string][]string{
		"0": {"9", "1"},
		"1": {"0", "2"},
		"2": {"1", "3"},
		"3": {"2", "4"},
		"4": {"3", "5"},
		"5": {"4", "6"},
		"6": {"5", "7"},
		"7": {"6", "8"},
		"8": {"7", "9"},
		"9": {"8", "0"},
	}
	runeStr := []rune(s)
	res := make([]string, 0)
	for idx, value := range runeStr {
		neighBors := neighborMap[string(value)]
		for _, neighBor := range neighBors {
			tempRune := make([]rune, len(runeStr))
			copy(tempRune, runeStr)
			tempRune[idx] = []rune(neighBor)[0]
			if deadMap[string(tempRune)] {
				continue
			} else {
				res = append(res, string(tempRune))
			}
		}
	}
	return res
}
