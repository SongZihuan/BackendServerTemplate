// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import "log"

const exitCodeSuccess = 0
const exitCodeFailed = 1

func ReturnError(err error) int {
	log.Printf("Error: %s\n", err.Error())
	return exitCodeFailed
}

func ReturnSuccess() int {
	log.Println("Success!")
	return exitCodeSuccess
}
