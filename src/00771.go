package src

/*
771. 宝石与石头
给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

*/

func numJewelsInStones(J string, S string) int {
	hash := make(map[rune]bool)
	for _, i2 := range J {
		hash[i2] = true
	}
	sum := 0
	for _, i2 := range S {
		if _, ok := hash[i2]; ok {
			sum += 1
		}
	}
	return sum
}
