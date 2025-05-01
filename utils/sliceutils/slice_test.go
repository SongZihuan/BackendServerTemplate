// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sliceutils

import "testing"

func TestCopySlice(t *testing.T) {
	src := []int{1, 2}
	dest := CopySlice(src)

	if len(dest) != len(src) || dest[0] != src[0] || dest[1] != src[1] {
		t.Errorf("copy slice failed: data does not match")
	}

	src[1] = 3
	if dest[1] == 3 {
		t.Errorf("copy slice failed: shared underlying array")
	}
}

func TestSliceHasItem(t *testing.T) {
	src := []int{1, 2}

	if !SliceHasItem(src, 1) {
		t.Errorf("SliceHasItem([]int{1, 2}, 1) -> true: false")
	}

	if SliceHasItem(src, 3) {
		t.Errorf("SliceHasItem([]int{1, 2}, 3) -> false: true")
	}
}
