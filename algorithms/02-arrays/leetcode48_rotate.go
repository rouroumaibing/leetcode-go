func rotate(matrix [][]int) {
	row := len(matrix[0])
	colum := len(matrix)
	//以左上角到右下角为中轴线，对称翻转
	for i := 0; i < colum; i++ {
		for j := i; j < row; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//对折翻转
	for i := 0; i < colum; i++ {
		revese(matrix[i])
	}
}

func revese(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}