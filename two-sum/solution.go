func twoSum(nums []int, target int) []int {
	// lookup table from nums array
	m := make(map[int]int)
	for i, num := range nums {
		q := target - num
		if j, match := m[q]; match {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}