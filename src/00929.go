package src

import "strings"

/*
929. 独特的电子邮件地址

每封电子邮件都由一个本地名称和一个域名组成，以 @ 符号分隔。

例如，在 alice@leetcode.com中， alice 是本地名称，而 leetcode.com 是域名。

除了小写字母，这些电子邮件还可能包含 '.' 或 '+'。

如果在电子邮件地址的本地名称部分中的某些字符之间添加句点（'.'），则发往那里的邮件将会转发到本地名称中没有点的同一地址。例如，"alice.z@leetcode.com” 和 “alicez@leetcode.com” 会转发到同一电子邮件地址。 （请注意，此规则不适用于域名。）

如果在本地名称中添加加号（'+'），则会忽略第一个加号后面的所有内容。这允许过滤某些电子邮件，例如 m.y+name@email.com 将转发到 my@email.com。 （同样，此规则不适用于域名。）

可以同时使用这两个规则。

给定电子邮件列表 emails，我们会向列表中的每个地址发送一封电子邮件。实际收到邮件的不同地址有多少？
*/
func numUniqueEmails(emails []string) int {
	hash := make(map[string]int)
	for _, email := range emails {
		arr := strings.Split(email, "@")
		arr[0] = strings.ReplaceAll(arr[0], ".", "")
		ind := strings.Index(arr[0], "+")
		if ind > -1 {
			arr[0] = arr[0][:ind]
		}
		key := arr[0] + "@" + arr[1]
		if _, ok := hash[key]; !ok {
			hash[key] = 1
		}
	}
	return len(hash)
}

func Test_numUniqueEmails() {
	numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"})
}
