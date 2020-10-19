package insertionsort

func Sort(arr []int) {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			} else {
				break
			}
		}
	}
}
