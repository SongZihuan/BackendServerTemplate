// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package typeutils

import "testing"

func TestStringBool(t *testing.T) {
	t.Run("enable default", func(t *testing.T) {
		var res StringBool
		if res.IsEnable() {
			t.Errorf("res.IsEnable() -> false: true")
		}
	})

	t.Run("enable default", func(t *testing.T) {
		var res StringBool
		if res.IsDisable() {
			t.Errorf("res.IsDisable() -> false: true")
		}
	})

	t.Run("enable default true", func(t *testing.T) {
		var res StringBool
		if !res.IsEnable(true) {
			t.Errorf("res.IsEnable(true) -> true: false")
		}
	})

	t.Run("enable default false", func(t *testing.T) {
		var res StringBool
		if res.IsEnable(false) {
			t.Errorf("res.IsEnable(false) -> false: true")
		}
	})

	t.Run("disable default true", func(t *testing.T) {
		var res StringBool
		if !res.IsDisable(true) {
			t.Errorf("res.IsDisable(true) -> true: false")
		}
	})

	t.Run("disable default false", func(t *testing.T) {
		var res StringBool
		if res.IsDisable(false) {
			t.Errorf("res.IsDisable(false) -> false: true")
		}
	})

	t.Run("set default disable", func(t *testing.T) {
		var res StringBool
		res.SetDefaultDisable()

		if !res.IsDisable() {
			t.Errorf("res.IsDisable() -> true: false")
		}
	})

	t.Run("set default enable", func(t *testing.T) {
		var res StringBool
		res.SetDefaultEnable()

		if !res.IsEnable() {
			t.Errorf("res.IsEnable() -> true: false")
		}
	})

	t.Run("set default double", func(t *testing.T) {
		var res StringBool
		res.SetDefaultEnable()
		res.SetDefaultDisable()

		if !res.IsEnable() {
			t.Errorf("res.IsEnable() -> true: false")
		}
	})

	t.Run("to bool", func(t *testing.T) {
		var res StringBool
		res.SetDefaultEnable()

		if res.ToBool() != true {
			t.Errorf("res.ToBool() -> true: false")
		}
	})
}
