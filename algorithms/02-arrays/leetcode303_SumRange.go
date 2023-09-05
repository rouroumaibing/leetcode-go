type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	snums := make([]int, len(nums)+1)
	for k, v := range nums {
		snums[k+1] = snums[k] + v
	}
	return NumArray{snums}
}

/*
 nums: -2,0,3,-5,2,-1
 snums: 0,-2,-2,1,-4,-2,-3
*/
func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
