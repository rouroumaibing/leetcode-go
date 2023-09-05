func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		tmp := numbers[left] + numbers[right]
		if tmp == target {
			return []int{left + 1, right + 1}
		} else if tmp < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}