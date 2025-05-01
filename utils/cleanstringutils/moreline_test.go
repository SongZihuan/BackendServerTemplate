// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cleanstringutils

import "testing"

func TestGetStringMoreLine(t *testing.T) {
	t.Run("Text-OnlyLine", func(t *testing.T) {
		text := "Hello"
		if GetStringOneLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-OnlyLine-WithSpace", func(t *testing.T) {
		text := "Hello    "
		if GetStringOneLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-CRLF", func(t *testing.T) {
		text := "Hello\r\n"
		if GetStringOneLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-More-CRLF", func(t *testing.T) {
		text := "Hello\r\n\r\n\r\n"
		if GetStringOneLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-CRLF-WithSpace", func(t *testing.T) {
		text := "Hello    \r\n"
		if GetStringOneLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-MoreLine", func(t *testing.T) {
		text := "Hello\r\nWorld"
		if GetStringOneLine(text) != "Hello\nWorld" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})
}
