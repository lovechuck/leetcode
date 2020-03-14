package src

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。


说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	var nums []int
	nums = append(nums, nums1...)
	i := 0
	j := 0
	k := 0
	for i < m && j < n {
		if nums[i] <= nums2[j] {
			nums1[k] = nums[i]
			i++
			k++
		} else {
			nums1[k] = nums2[j]
			k++
			j++
		}
	}
	if i != m {
		for i < m {
			nums1[k] = nums[i]
			i++
			k++
		}
	}
	if j != n {
		for j < n {
			nums1[k] = nums2[j]
			k++
			j++
		}
	}
}

func Test_merge() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
