// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	_ "github.com/SongZihuan/BackendServerTemplate/src/global"
	lionv1 "github.com/SongZihuan/BackendServerTemplate/src/mainfunc/lion/v1"
	"os"
)

func main() {
	os.Exit(lionv1.MainV1())
}
