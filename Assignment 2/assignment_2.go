package assignmenttwo

func contains(dataList []string, searchedData string) bool {
	if len(dataList) == 0 {
		return false
	} else if dataList[0] == searchedData {
		return true
	}

	return contains(dataList[1:], searchedData)
}
