package src

/*
925. 长按键入

你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。

你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。
*/
func isLongPressedName(name string, typed string) bool {
	j := 0
	for i := 0; i < len(name); i++ {
		if j == len(typed) {
			return false
		}
		if name[i] == typed[j] {
			j++
		} else {
			for j > 0 && j < len(typed) && typed[j] == typed[j-1] {
				j++
			}
			if j == len(typed) {
				return false
			} else {
				if typed[j] == name[i] {
					j++
				} else {
					return false
				}
			}
		}
	}
	for ; j < len(typed); j++ {
		if typed[j] != name[len(name)-1] {
			return false
		}
	}
	return true
}
