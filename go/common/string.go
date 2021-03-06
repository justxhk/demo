package common

import (
	"bytes"
)

func StringBuilder(strings ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range strings {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func contains(strA string, strB string) bool {
	n := len(strA)
	m := len(strB)
	for i := 0; i < n-m+1; i++ {

	}
	return false
}

/**
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
*/
func ReverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
