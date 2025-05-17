// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package combiningwriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logwriter"
	"github.com/SongZihuan/BackendServerTemplate/utils/sliceutils"
	"sync"
)

type CombiningWriter struct {
	level  loglevel.LoggerLevel
	tag    bool
	writer []logwriter.Writer
	close  bool
	mutex  sync.Mutex
}

func (w *CombiningWriter) Write(data *logformat.LogData) chan any {
	res := make(chan any)

	// 此处 w.level 是只读的，因此可以不上锁操作
	if (w.level.Int() > data.Level.Int()) || (data.Level == loglevel.PseudoLevelTag && !w.tag) {
		close(res)
		return res
	}

	go func() {
		w.write(data)
		close(res)
	}()

	return res
}

func (c *CombiningWriter) write(data *logformat.LogData) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.close {
		return
	}

	var wg sync.WaitGroup

	for _, w := range c.writer {
		if w == nil {
			continue
		}
		go func() {
			wg.Add(1)
			res := w.Write(data)
			<-res
			wg.Done()
		}()
	}

	wg.Wait()
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

	res.level = loglevel.LevelDebug
	res.tag = true
	res.writer = sliceutils.CopySlice(w)
	res.close = false

	return res
}

func NewCombiningWriterWithLevel(level loglevel.LoggerLevel, tag bool, w ...logwriter.Writer) *CombiningWriter {
	var res = new(CombiningWriter)

	res.level = level
	res.tag = tag
	res.writer = sliceutils.CopySlice(w)
	res.close = false

	return res
}
