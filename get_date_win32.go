// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build windows

//go:generate powershell -ExecutionPolicy RemoteSigned ./get_date.ps1

//go:generate powershell -ExecutionPolicy RemoteSigned ./get_git.ps1

package resource
