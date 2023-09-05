func longestPalindrome(s string) string {
	var res string
	for index := 0; index < len(s); index++ {
		s1 := palindrome(s, index, index)
		s2 := palindrome(s, index, index+1)

		if len(res) < len(s1) {
			res = s1
		}

		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

//返回以left和right为中心字符串的回文字符串
func palindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}