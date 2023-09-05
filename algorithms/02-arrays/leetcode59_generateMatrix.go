func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	upper, lower, left, right := 0, n-1, 0, n-1
	num := 1

	for num <= n*n {
		//从左往右检索，检索完upper下移一行，upper不超过lower
		if upper <= lower {
			for tmp := left; tmp <= right; tmp++ {
				matrix[upper][tmp] = num
				num++
			}

			upper++
		}
		//从上往下检索，检索完right左移一列，right不超过left
		if left <= right {
			for tmp := upper; tmp <= lower; tmp++ {
				matrix[tmp][right] = num
				num++
			}

			right--
		}
		//从右往左检索，检索完lower上移一行, lower不超过upper
		if upper <= lower {
			for tmp := right; tmp >= left; tmp-- {
				matrix[lower][tmp] = num
				num++
			}

			lower--
		}
		//从下往上检索，检索完left右移一列, left不超过right
		if left <= right {
			for tmp := lower; tmp >= upper; tmp-- {
				matrix[tmp][left] = num
				num++
			}

			left++
		}
	}
	return matrix
}