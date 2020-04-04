package src

import (
	"flag"
	"math"
)

/*
最短单词距离

Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.

Example:
Assume that words = ["practice", "makes", "perfect", "coding", "makes"].

Input: word1 = “coding”, word2 = “practice”
Output: 3
Input: word1 = "makes", word2 = "coding"
Output: 1
Note:
You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.
*/

func shortestDistance(words []string, word1 string, word2 string) int {
	i1 := -1
	i2 := -1
	result := math.MaxInt64
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			i1 = i
		} else if words[i] == word2 {
			i2 = i
		}
		if i1 != -1 && i2 != -1 {
			result = min(result, mabs(i1, i2))
		}
	}
}
