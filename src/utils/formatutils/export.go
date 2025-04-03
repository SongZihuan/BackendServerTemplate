// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package formatutils

import "strings"

const NormalConsoleWidth = 80

func FormatTextToWidth(text string, width int) string {
	return FormatTextToWidthAndPrefix(text, 0, width)
}

func FormatTextToWidthAndPrefix(text string, prefixWidth int, overallWidth int) string {
	var result strings.Builder

	width := overallWidth - prefixWidth
	if width <= 0 {
		panic("bad width")
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")

	for _, line := range strings.Split(text, "\n") {
		result.WriteString(strings.Repeat(" ", prefixWidth))

		if line == "" {
			result.WriteString("\n")
			continue
		}

		spaceCount := CountSpaceInStringPrefix(line) % width
		newLineLength := 0
		if spaceCount < 80 {
			result.WriteString(strings.Repeat(" ", spaceCount))
			newLineLength = spaceCount
		}

		for _, word := range strings.Fields(line) {
			if newLineLength+len(word) >= width {
				result.WriteString("\n")
				result.WriteString(strings.Repeat(" ", prefixWidth))
				newLineLength = 0
			}

			// 不是第一个词时，添加空格
			if newLineLength != 0 {
				result.WriteString(" ")
				newLineLength += 1
			}

			result.WriteString(word)
			newLineLength += len(word)
		}

		if newLineLength != 0 {
			result.WriteString("\n")
			newLineLength = 0
		}
	}

	return strings.TrimRight(result.String(), "\n")
}

func CountSpaceInStringPrefix(str string) int {
	var res int
	for _, r := range str {
		if r == ' ' {
			res += 1
		} else {
			break
		}
	}

	return res
}
