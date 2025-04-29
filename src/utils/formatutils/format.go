// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package formatutils

import (
	"strings"
)

const NormalConsoleWidth = 80

// FormatTextToWidth 把文本控制在固定长度内（直接指定长度 width）
func FormatTextToWidth(text string, width int) string {
	return FormatTextToWidthAndPrefix(text, 0, width)
}

// FormatTextToWidthAndPrefix 把文本控制在固定的长度
// prefixWidth 每行开头的空格
// overallWidth 每行总长度
// 实际每行字符数：width = overallWidth - prefixWidth
func FormatTextToWidthAndPrefix(text string, prefixWidth int, overallWidth int) string {
	var result strings.Builder

	width := overallWidth - prefixWidth
	if width <= 0 {
		panic("bad width")
	}

	text = strings.TrimRight(strings.Replace(text, "\r", "", -1), "\n")

LineCycle:
	for _, line := range strings.Split(text, "\n") { // 逐行遍历
		result.WriteString(strings.Repeat(" ", prefixWidth)) // 输出当前行的 prefix 空格

		if line == "" { // 如果当前行为空则直接换行返回
			result.WriteString("\n")
			continue LineCycle
		}

		newLineLength := 0

		spaceCount := countSpaceInStringPrefix(line) % width // 获取当前行的开头空格，但空格数不超过单行字符总长度（有时很首行空格起到一定语法作用，因此需要保留，而下面的 for 循环使用 strings.Fields 分割字符串会导致空格被忽略，因此在此处要提前处理）
		if spaceCount != 0 {
			result.WriteString(strings.Repeat(" ", spaceCount))
			newLineLength = spaceCount
		}

		line = strings.TrimSpace(line) // 空格已在上面处理，此处可以把空格删除

		for _, word := range strings.Fields(line) { // 使用 strings.Fields 遍历每一行的每一个单词
			if newLineLength+len(word) >= width { // 若写入新单词后超过单行总长度，则换行。
				result.WriteString("\n") // 输出 "\n"，并输出当前行的 prefix 空格。（从第二行开始、本循环中每增加一行，就要在这里写入 prefix 空格。而第一行的 prefix 空格则 LineCycle 一开始的时候写入）
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
			result.WriteString("\n") // 写入最后的换行符
			newLineLength = 0
		}
	}

	return strings.TrimRight(result.String(), "\n")
}

// countSpaceInStringPrefix 计算字符串开头的空格数
func countSpaceInStringPrefix(str string) int {
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
