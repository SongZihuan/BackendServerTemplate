// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	tigerv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/tiger/v1"
	"os"
)

func main() {
	os.Exit(tigerv1.MainV1())
}
