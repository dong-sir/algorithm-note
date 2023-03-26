/*
Copyright 2023 dong-sir.

selectionSort 选择排序

*/

package sorting

// SelectionSort 可没有返回值，为了方便比较前后变化增加返回值
func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			temp := nums[i]
			nums[i] = nums[minIndex]
			nums[minIndex] = temp
		}
	}
	return nums
}
