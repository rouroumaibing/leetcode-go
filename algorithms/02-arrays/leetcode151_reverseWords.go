func reverseWords(s string) string {
	//清理字符串前后空格、字符中间空格
	tmp := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != ' ' {
			result = append(result, tmp[i])
		} else if len(result) != 0 && result[len(result)-1] != ' ' {
			result = append(result, ' ')
		}
	}

	if result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	//整体翻转
	tmplong := len(result)
	reverse(result, 0, tmplong-1)

	//单词翻转
	long := len(result)
	fast, slow := 0, 0
	for fast < long {
		if fast+1 == long || result[fast+1] == ' ' {
			reverse(result, slow, fast)
			fast++
			slow = fast + 1
		}
		fast++
	}

	return string(result)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}