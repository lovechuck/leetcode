package src

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

Area = Max(min(height[i], height[j]) * (j-i)) {其中0 <= i < j < height,size()}

上面是一种双指针算法，指针 i 指向数组的开头，满足条件后只向后移动；指针 j 指向数组的末尾，满足条件后只向前移动。两指针相遇，算法结束。能使用这种算法解的题还有很多，在最后我还会举个例子。
那么指针移动的条件是什么呢？
在本题中指针移动的条件是，（首先假设数组为 a）如果 a[i] <= a[j]，那么指针 i 则向后移动；如果 a[i] > a[j]，那么指针 j 则向前移动。
我们可以发现，虽然 i，j 指针合起来能遍历整个数组，但并没有穷举所有的容器情况。例如，如果 a[i] <= a[j]，则 i++（移动 j 的情况同理证得），排除了 a[i] 与 a[i+1,...,j-1] 之间的某个数构成最大盛水容器的可能性。为什么可以这样做呢？下面给出证明：
假设存在 k，k 满足 i+1 <= k <= j-1，a[i] 与 a[k] 构成了最大盛水容器，容量为 C。

如果 a[i] <= a[k]，则 C = a[i]*(k - i)，那么我们现在来看一下，a[i] 和 a[j] 构成的盛水容器，容量如何，其容量为 a[i]*(j - i) > C，此时假设不成立！
同理，如果 a[i] > a[k]，则 C = a[k]*(k - i)，注意此时，由不等式的传递性可得a[j] 和 a[k] 的大小关系为 a[j] > a[k]。所以依然有 a[i] 和 a[j] 构成的盛水容器，容量为 a[i]*(j - i) > C，此时假设依然不成立！
综上可得，a[i] 不可能与 a[i+1,...,j-1] 之间的某个数构成最大盛水容器。证明结束！
*/

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	s := (right - left) * min(height[left], height[right])
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		tmp := (right - left) * min(height[left], height[right])
		if s < tmp {
			s = tmp
		}
	}
	return s
}
