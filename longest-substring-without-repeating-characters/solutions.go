func lengthOfLongestSubstring(s string) int {
	// use sliding window

	// work on runes instead of bytes
	// avoids problems with indexing
	rs := []rune(s)

	// store where each char is in the window
	m := make(map[rune]int)

	max := 0
	windowStart := 0

	for i, c := range rs {
		if j, found := m[c]; found {
			if windowSize := i - windowStart; windowSize > max {
				max = windowSize
			}
			// delete everything between windowStart and (including) j
			for k := windowStart; k <= j; k++ {
				delete(m, rs[k])
			}
			// new window starts after the duplicate
			windowStart = j + 1
		}
		m[c] = i
	}
	// check if the latest window is larger
	if windowSize := len(rs) - windowStart; windowSize > max {
		max = windowSize
	}
	return max
}