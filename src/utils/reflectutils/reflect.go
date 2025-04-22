// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package reflectutils

import "reflect"

func HasFieldByReflect(typ reflect.Type, fieldName string) bool {
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Name == fieldName {
			return true
		}
	}
	return false
}
