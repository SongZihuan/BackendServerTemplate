// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package executils

import (
	"bytes"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"os"
	"os/exec"
)

func Run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)

	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}

func RunOnline(name string, args ...string) (string, error) {
	res, err := Run(name, args...)
	if err != nil {
		return "", err
	}

	return cleanstringutils.GetStringOneLine(res), nil
}

func RunBytes(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)

	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
