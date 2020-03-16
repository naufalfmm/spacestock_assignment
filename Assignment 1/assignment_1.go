package assignmentone

func sum(dataList []int) int {
	if len(dataList) == 1 {
		return dataList[0]
	} else if len(dataList) == 0 {
		return 0
	}

	return sum(dataList[1:]) + dataList[0]
}
