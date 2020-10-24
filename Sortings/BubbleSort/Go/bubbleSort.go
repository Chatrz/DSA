/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package bubblesort

func Sort(arr []int) {
	arrLen := len(arr)
	var tmp int
	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

func OptimizedSort(arr []int) {
	arrLen := len(arr)
	var tmp int
	var flag bool
	for i := 0; i < arrLen && !flag; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
				flag = false
			}
		}
	}
}
