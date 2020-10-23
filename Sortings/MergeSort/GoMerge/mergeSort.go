package GoMerge

import (
	"fmt"
	"math/rand"
	"time"
)

func TestSort() []int  {
	arr := GetRandomArray(100)
	fmt.Println("before sort : ")
	fmt.Println(arr)
	fmt.Println("after sort : ")
	fmt.Println(MergeSort(arr))
	return arr
}
func GetRandomArray(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func MergeSort(arr []int) []int {
	sorted := []int{}
	if len(arr) == 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			sorted = append(sorted, arr[1])
			sorted = append(sorted, arr[0])
			return sorted
		}
		return arr
	}
	mid := len(arr) / 2
	first := arr[:len(arr)/2]
	second := arr[mid:]
	first = MergeSort(first)
	second = MergeSort(second)
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			sorted = append(sorted, first[i])
			i++
		} else {
			sorted = append(sorted, second[j])
			j++
		}
	}
	sorted = append(sorted, first[i:]...)
	sorted = append(sorted, second[j:]...)
	return sorted
}
