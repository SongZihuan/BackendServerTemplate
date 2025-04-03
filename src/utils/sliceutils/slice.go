// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sliceutils

func CopySlice[T any](src []T) []T {
	dest := make([]T, len(src))
	copy(dest, src)
	return dest
}
