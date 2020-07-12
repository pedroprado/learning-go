package sort

import "sort"

func BubbleSort(arr []int) bool {
	var swapped bool
	var sorted = false
	for i := 0; i < len(arr)-1; i++ {
		swapped = false
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

		if !swapped {
			if i == 0 {
				sorted = true
			}
			break
		}
	}
	return sorted
}

func NormalSort(arr []int) {
	sort.Ints(arr)
}
