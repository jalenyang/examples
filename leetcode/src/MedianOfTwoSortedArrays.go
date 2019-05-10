package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := nums1
	for i, _ := range nums2 {
		arr = append(arr, nums2[i])
	}
	n3 := len(arr)
	for i := 0; i < n3; i++ {
		for j := i + 1; j < n3; j++ {
			temp := arr[i]
			if arr[i] < arr[j] {
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	if len(arr)%2 == 0 {
		return float64(arr[n3/2]+arr[n3/2-1]) / 2
	} else {
		return float64(arr[n3/2])
	}
}

func main() {
	num1 := []int{1, 2, 3}
	num2 := []int{1}
	f := findMedianSortedArrays(num1, num2)
	fmt.Println(f)
}
