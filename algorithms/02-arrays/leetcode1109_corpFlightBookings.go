/*
capacity 代表容量。超过车在路上当个点的最大容量，超过就失败，即上下车后（加减）人数不能超。
maxtoi 代表当前组最远距离
*/

func carPooling(trips [][]int, capacity int) bool {
	//找出汽车行驶最长距离
	var maxtoi int
	for _, v := range trips {
		if v[2] > maxtoi {
			maxtoi = v[2]
		}
	}

	//构造等差数组，刚出站的车上没有人，均为0
	diff := make([]int, maxtoi)
	//等差数组加减值，人员上下车
	for _, v := range trips {
		//上车人数最大多少。左+最多上车人数
		diff[v[1]] += v[0]
		//到站即下车，人数最大多少。右-最多下车人数
		if v[2] < maxtoi {
			diff[v[2]] -= v[0]
		}
	}

	//汽车全程最多人数数组
	answer := make([]int, maxtoi)
	answer[0] = diff[0]
	for k := 0; k < maxtoi-1; k++ {
		answer[k+1] = answer[k] + diff[k+1]
	}

	for _, v := range answer {
		//存在超过最大容量人数即false
		if v > capacity {
			return false
		}
	}
	return true
}
