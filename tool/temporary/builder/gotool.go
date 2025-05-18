// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/utils/executils"
	"os"
	"os/exec"
)

func goGenerate() error {
	_, err := executils.Run("go", "generate", "./...")
	if err != nil {
		return err
	}

	return nil
}

func goBuild(goos string, goarch string, target string, output string) error {
	packagePath, ok := packageMap[target]
	if !ok {
		return fmt.Errorf("target [%s] is invalid", target)
	}

	cmd := exec.Command(
		"go",
		"build",
		"-o",
		output,
		"-trimpath",
		`-ldflags`,
		`-s -w -extldflags "-static"`,
		`-gcflags`,
		`all=-l=4`,
		packagePath,
	)

	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env,
		"GOOS="+goos,
		"GOARCH="+goarch,
		"CGO=1",
		"GO111MODULE=on",
	)

	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
