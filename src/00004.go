package src

import (
	"fmt"
)

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength&1 == 1 {
		return float64(findKthElem(nums1, nums2, totalLength/2+1))
	}
	return 0.5 * float64(findKthElem(nums1, nums2, totalLength/2)+findKthElem(nums1, nums2, totalLength/2+1))

}

//1<=k<=len(nums1)+len(nums2)
func findKthElem(nums1, nums2 []int, k int) int {
	var i, j, l1 int
	for {
		if len(nums1) > len(nums2) {
			nums1, nums2 = nums2, nums1
		}
		l1 = len(nums1)
		if l1 == 0 {
			return nums2[k-1]
		}
		if k == 1 {
			return min(nums1[0], nums2[0])
		}
		if k/2 > l1 {
			i = l1
		} else {
			i = k / 2
		}
		j = k - i
		// i,j <= k/2 <= l1 <= l2
		if nums1[i-1] > nums2[j-1] {
			nums2 = nums2[j:]
			k -= j
		} else if nums2[j-1] > nums1[i-1] {
			nums1 = nums1[i:]
			k -= i
		} else {
			return nums1[i-1]
		}
	}
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func Test_findMedianSortedArrays() {
	nums1 := []int{1, 3, 5}
	nums2 := []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
