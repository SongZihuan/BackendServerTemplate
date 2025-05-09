// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package combiningwriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"github.com/SongZihuan/BackendServerTemplate/utils/sliceutils"
	"sync"
)

type CombiningWriter struct {
	writer []logwriter.Writer
	close  bool
	mutex  sync.Mutex
}

func (c *CombiningWriter) Write(data *logformat.LogData) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.close {
		return
	}

	for _, w := range c.writer {
		if w == nil {
			continue
		}
		go w.Write(data)
	}
}

func (c *CombiningWriter) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	defer func() {
		c.close = true
	}()

	errMsg := ""
	for _, w := range c.writer {
		if w == nil {
			continue
		}

		errTmp := w.Close()
		if errTmp != nil {
			errMsg += errTmp.Error() + ";"
		}
	}

	if errMsg == "" {
		return nil
	}

	return fmt.Errorf(errMsg)
}

func NewCombiningWriter(w ...logwriter.Writer) *CombiningWriter {
	var res = new(CombiningWriter)

	res.writer = sliceutils.CopySlice(w)
	res.close = false

	return res
}
