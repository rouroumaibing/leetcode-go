/*
func corpFlightBookings(bookings [][]int, n int) []int {
    answer := make([]int, n)
    for _, v := range bookings {
        answer = difference(answer, v[0]-1, v[1]-1, v[2])
    }
    return answer
}

func difference(bookinfo []int, i, j, val int) []int {
    //构建等差数组
    long := len(bookinfo)
    diff := make([]int, long)
    diff[0] = bookinfo[0]
    for k:=0;k<long -1 ;k++{
        diff[k+1] = bookinfo[k+1] - bookinfo[k]
    }
    //按要求增加数值差
    diff[i] += val

    if j+1 < long {
        diff[j+1] -= val
    }
    //应用到传入的数组
    answer := make([]int, long)
    answer[0] = diff[0]
    for k:=0;k<long -1 ;k++{
        answer[k+1] = answer[k] + diff[k+1]
    }
    return answer
}
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	answer := make([]int, n)

	//初始化等差数组
	//当前原始值均为0，等差数组均为0
	/*
	   long := n
	   diff := make([]int, long)
	   diff[0] = bookinfo[0]
	   for k:=0;k<long -1 ;k++{
	       diff[k+1] = bookinfo[k+1] - bookinfo[k]
	   }
	*/
	diff := make([]int, n)

	//按要求增加数值差
	for _, v := range bookings {
		//左边第一个需要加val
		diff[v[0]-1] += v[2]
		//右边需要在第右+1个减val
		if v[1] < n {
			diff[v[1]] -= v[2]
		}
	}

	//还原数组
	answer[0] = diff[0]
	for k := 0; k < n-1; k++ {
		answer[k+1] = answer[k] + diff[k+1]
	}
	return answer
}

