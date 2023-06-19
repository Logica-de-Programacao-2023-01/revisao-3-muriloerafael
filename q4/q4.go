package q4

func SingleNumber(nums []int) int {
	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num]++
	}

	for i, num := range nums {
		if numMap[num] == 1 {
			return i
		}
	}

	return -1
}
