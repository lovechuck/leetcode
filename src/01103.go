package src

/*
排排坐，分糖果。

我们买了一些糖果 candies，打算把它们分给排好队的 n = num_people 个小朋友。

给第一个小朋友 1 颗糖果，第二个小朋友 2 颗，依此类推，直到给最后一个小朋友 n 颗糖果。

然后，我们再回到队伍的起点，给第一个小朋友 n + 1 颗糖果，第二个小朋友 n + 2 颗，依此类推，直到给最后一个小朋友 2 * n 颗糖果。
*/

func distributeCandies(candies int, num_people int) []int {
	n := 0
	ans := make([]int, num_people)
	for candies > 0 {
		for i := 0; i < num_people; i++ {
			if candies < n*num_people+i+1 {
				ans[i] += candies
				candies = 0
			} else {
				ans[i] += n*num_people + i + 1
				candies -= n*num_people + i + 1
			}
			if candies <= 0 {
				break
			}
		}
		n++
	}
	return ans
}

func Test_distributeCandies() {
	distributeCandies(7, 4)
}
