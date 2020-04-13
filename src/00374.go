package src

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return 1
}

func guessNumber(n int) int {
	left := 1
	right := n
	for left < right {
		mid := (left + right) / 2
		ges := guess(mid)
		if ges == 0 {
			return mid
		} else if ges < 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
