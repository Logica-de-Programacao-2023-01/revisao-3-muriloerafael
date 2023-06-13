package q4

func SingleNumber(nums []int) int {
	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num] += 1
	}

	for key, value := range numMap {
		if value == 1 {
			return key
		}
	}

	return 0
}
