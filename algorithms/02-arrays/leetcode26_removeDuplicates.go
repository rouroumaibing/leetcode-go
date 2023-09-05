func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	fast, slow := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}