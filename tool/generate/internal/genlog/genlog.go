// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package genlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("generate: ")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lmsgprefix)
}

func InitGenLog(name string, output io.Writer) {
	if name = strings.TrimSpace(name); name != "" {
		log.SetPrefix(fmt.Sprintf("%s: ", name))
	}

	if output != nil {
		log.SetOutput(output)
	}
}

func GenLogf(format string, args ...any) {
	msg := strings.TrimSpace(strings.TrimRight(fmt.Sprintf(format, args...), "\n"))
	log.Printf("%s\n", msg)
}

func GenLog(msg string) {
	log.Println(strings.TrimSpace(strings.TrimRight(msg, "\n")))
}
