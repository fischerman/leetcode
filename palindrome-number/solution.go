import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	rs := []rune(s)
	length := len(rs)

	for i := 0; i < length/2; i++ {
		if rs[i] != rs[length-i-1] {
			return false
		}
	}
	return true
}