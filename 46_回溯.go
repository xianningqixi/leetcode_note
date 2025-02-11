package main

/*
算法模版
result = []
def backtrack(路径, 选择列表):
	if 满足结束条件:
	result.add(路径)
	return

for 选择 in 选择列表:
	做选择
	backtrack(路径, 选择列表)
	撤销选择
*/

/*
func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Printf("Hello and welcome, %+v!", result)
	return
}
*/

func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))
	backTrack(nums, used, track, &res)
	return res
}

// 路径：记录在 track 中
// 选择列表：nums 中不存在于 track 的那些元素
// 结束条件：nums 中的元素全都在 track 中出现
func backTrack(nums []int, used []bool, track []int, res *[][]int) {
	// 回溯结束条件
	if len(track) == len(nums) {
		// 深拷贝
		tempTrack := make([]int, len(track))
		copy(tempTrack, track)
		*res = append(*res, tempTrack)
		return
	}
	for i := range nums {
		if used[i] {
			continue
		}
		//做选择
		used[i] = true
		track = append(track, nums[i])
		//进入下一层
		backTrack(nums, used, track, res)
		//撤销选择
		used[i] = false
		track = track[:len(track)-1]
	}
}
