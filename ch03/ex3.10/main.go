package main

import (
	"bytes"
	"fmt"
	"os"
)

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func commaIterative(s string) string {
	var buf bytes.Buffer
	rem := len(s) % 3
	if rem == 0 {
		rem = 3
	}
	buf.WriteString(s[:rem])
	for i := rem; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", commaIterative(os.Args[i]))
	}
}
