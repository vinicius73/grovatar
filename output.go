package main

import (
	"bytes"
	"fmt"
)

// OutputStdout ...
func OutputStdout(avatar AvatarData) {
	for _, line := range avatar.Table {
		fmt.Println(lineToString(line))
	}
}

func lineToString(line [5]bool) string {
	var buffer bytes.Buffer

	for _, val := range line {
		if val == true {
			buffer.WriteString("⬛")
		} else {
			buffer.WriteString("⬜")
		}
	}

	return buffer.String()
}
