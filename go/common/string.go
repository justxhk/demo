package common

import "bytes"

func StringBuilder(strings ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range strings {
		buffer.WriteString(str)
	}
	return buffer.String()
}
