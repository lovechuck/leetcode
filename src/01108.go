package src

import "strings"

/*
1108. IP 地址无效化
*/

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
