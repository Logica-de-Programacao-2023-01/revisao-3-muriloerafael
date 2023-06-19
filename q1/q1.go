package q1

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for j, num := range nums {
		complement := target - num
		if i, ok := numMap[complement]; ok {
			return []int{i, j}
		}

		numMap[num] = j
	}

	return nil
}
