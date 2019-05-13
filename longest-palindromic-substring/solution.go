func longestPalindrome(s string) string {
	rs := []rune(s)

	if len(rs) == 1 {
		return string(rs)
	}

	l := len(rs)

	maxLen := 0
	maxString := []rune("")

	for middle, _ := range rs {
		// palindrom of odd length (single char in the middle)
		maxLen, maxString = expand(middle, middle, rs, l, maxLen, maxString)
		// palindrom of even length (two chars in the middle)
		if middle+1 != l {
			maxLen, maxString = expand(middle, middle+1, rs, l, maxLen, maxString)
		}
	}
	return string(maxString)
}

func expand(below, above int, rs []rune, l, maxLen int, maxString []rune) (int, []rune) {

	if rs[below] == rs[above] {
		for below-1 >= 0 && above+1 < l {
			if rs[below-1] != rs[above+1] {
				break
			} else {
				below--
				above++
			}
		}
	} else {
		return maxLen, maxString
	}

	if newL := above - below + 1; newL > maxLen {
		maxLen = newL
		maxString = rs[below : above+1]
	}
	return maxLen, maxString
}