func spiralOrder(matrix [][]int) []int {
	column, row := len(matrix), len(matrix[0])
	upper, lower, left, right := 0, column-1, 0, row-1
	cap := column * row
	result := make([]int, 0, cap)
	for len(result) < cap {

		//从左往右检索，检索完upper下移一行，upper不超过lower
		if upper <= lower {
			for tmp := left; tmp <= right; tmp++ {
				result = append(result, matrix[upper][tmp])
			}

			upper++
		}
		//从上往下检索，检索完right左移一列，right不超过left
		if left <= right {
			for tmp := upper; tmp <= lower; tmp++ {
				result = append(result, matrix[tmp][right])
			}

			right--
		}
		//从右往左检索，检索完lower上移一行, lower不超过upper
		if upper <= lower {
			for tmp := right; tmp >= left; tmp-- {
				result = append(result, matrix[lower][tmp])
			}

			lower--
		}
		//从下往上检索，检索完left右移一列, left不超过right
		if left <= right {
			for tmp := lower; tmp >= upper; tmp-- {
				result = append(result, matrix[tmp][left])
			}

			left++
		}
	}
	return result
}