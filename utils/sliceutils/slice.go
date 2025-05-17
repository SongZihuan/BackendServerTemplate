// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sliceutils

import "fmt"

func CopySlice[T any](src []T) []T {
	dest := make([]T, len(src))
	copy(dest, src)
	return dest
}

func SliceHasItem[T comparable](src []T, item T) bool {
	for _, i := range src {
		if i == item {
			return true
		}
	}
	return false
}

func MapToSlice[K comparable, V any](src map[K]V) []string {
	res := make([]string, 0, len(src))

	for k, v := range src {
		res = append(res, fmt.Sprintf("%v=%v", k, v))
	}

	return res
}
