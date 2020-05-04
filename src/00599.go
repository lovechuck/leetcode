package src

import "math"

/*
599. 两个列表的最小索引总和
*/
func findRestaurant(list1 []string, list2 []string) []string {
	hash := make(map[string]int)
	var res []string
	for i, s := range list1 {
		hash[s] = i
	}
	min := math.MaxInt32
	for i, s := range list2 {
		if _, ok := hash[s]; ok {
			if min > i+hash[s] {
				min = i + hash[s]
				res = []string{
					s,
				}
			} else if min == i+hash[s] {
				res = append(res, s)
			}
		}
	}

	return res
}
