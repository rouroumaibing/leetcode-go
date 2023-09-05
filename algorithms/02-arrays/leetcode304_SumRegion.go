type NumMatrix struct {
	matrix [][]int
}

/*
 * matrix[0,0,4,4]
 * 3,0,1,4,2
 * 5,6,3,2,1
 * 1,2,0,1,5
 * 4,1,0,1,7
 * 1,0,3,0,5
 *
 * sums[0,0,i-1,j-1]
 * 0,0,0,0,0,0
 * 0,3,3,4,8,10
 * 0,8,14,18,24,27
 * 0,9,17,21,28,36
 * 0,13,22,26,34,49
 * 0,14,23,30,38,58
 *
 */

func Constructor(matrix [][]int) NumMatrix {
	rowlong := len(matrix[0])
	collong := len(matrix)
	if rowlong == 0 || collong == 0 {
		return NumMatrix{matrix}
	}

	sums := make([][]int, collong+1)
	sums[0] = make([]int, rowlong+1)
	for i := 0; i < collong; i++ {
		sums[i+1] = make([]int, rowlong+1)
		for j := 0; j < rowlong; j++ {
			sums[i+1][j+1] = sums[i+1][j] + sums[i][j+1] - sums[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	//21-9-4+3=11
	return this.matrix[row2+1][col2+1] - this.matrix[row1][col2+1] - this.matrix[row2+1][col1] + this.matrix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */