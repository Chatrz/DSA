package selectionsort

func Sort(arr []int) {
	arrLen := len(arr)
	var minIndex int
	for i := 0; i < arrLen; i++ {
		minIndex = i
		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}

}
