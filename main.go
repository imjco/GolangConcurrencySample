package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(time.Since(now))

	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

}
func merge(nums1 []int, m int, nums2 []int, n int) {

	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	for i := 0; i < m+n; i++ {
		for j := i + 1; j < m+n; j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}

	fmt.Println(nums1)
}
