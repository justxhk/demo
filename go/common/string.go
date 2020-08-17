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
