package leetcode

import (
	"bytes"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	repeat := 1
	buf := bytes.NewBufferString(A)
	for ; buf.Len() < len(B); repeat++ {
		buf.WriteString(A)
	}
	if strings.Contains(buf.String(), B) {
		return repeat
	}
	buf.WriteString(A)
	if strings.Contains(buf.String(), B) {
		return repeat + 1
	}
	return -1
}
