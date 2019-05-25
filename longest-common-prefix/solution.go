func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	max := len(strs[0])
	for i := 0; i < len(strs); i++ {
		if l := len(strs[i]); l < max {
			max = l
		}
	}
	var s []byte
	for cur := 0; cur < max; cur++ {
		c := strs[0][cur]
		diverge := false
		for i := 1; i < len(strs); i++ {
			if c != strs[i][cur] {
				diverge = true
				break
			}
		}
		if diverge {
			break
		} else {
			s = append(s, c)
		}
	}
	return string(s)
}