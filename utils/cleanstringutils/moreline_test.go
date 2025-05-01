// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cleanstringutils

import "testing"

func TestGetStringMoreLine(t *testing.T) {
	t.Run("Text-OnlyLine", func(t *testing.T) {
		text := "Hello"
		if GetString(text) != "Hello" {
			t.Errorf("test GetString error")
		}
	})

	t.Run("Text-OnlyLine-WithSpace", func(t *testing.T) {
		text := "Hello    "
		if GetString(text) != "Hello" {
			t.Errorf("test GetString error")
		}
	})

	t.Run("Text-With-CRLF", func(t *testing.T) {
		text := "Hello\r\n"
		if GetString(text) != "Hello" {
			t.Errorf("test GetString error")
		}
	})

	t.Run("Text-With-More-CRLF", func(t *testing.T) {
		text := "Hello\r\n\r\n\r\n"
		if GetString(text) != "Hello" {
			t.Errorf("test GetString error")
		}
	})

	t.Run("Text-With-CRLF-WithSpace", func(t *testing.T) {
		text := "Hello    \r\n"
		if GetString(text) != "Hello" {
			t.Errorf("test GetString error")
		}
	})

	t.Run("Text-MoreLine", func(t *testing.T) {
		text := "Hello\r\nWorld"
		if GetString(text) != "Hello\nWorld" {
			t.Errorf("test GetString error")
		}
	})
}
