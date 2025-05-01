// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package combiningwriter

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/logformat"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/write"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/sliceutils"
)

type CombiningWriter struct {
	writer []write.Writer
	closer []write.WriteCloser
	close  bool
}

func (c *CombiningWriter) Write(data *logformat.LogData) (n int, err error) {
	if c.close {
		return 0, fmt.Errorf("combining writer has been close")
	}

	n = 0
	errMsg := ""
	for _, w := range c.writer {
		if w == nil {
			continue
		}

		nTmp, errTmp := w.Write(data)
		if errTmp != nil {
			errMsg += errTmp.Error() + ";"
		}

		n = min(n, nTmp)
	}

	if errMsg == "" {
		return n, nil
	}

	return n, fmt.Errorf(errMsg)
}

func (c *CombiningWriter) Close() error {
	defer func() {
		c.close = true
	}()

	errMsg := ""
	for _, w := range c.closer {
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

func NewCombiningWriter(w ...write.Writer) *CombiningWriter {
	var res = new(CombiningWriter)

	res.writer = sliceutils.CopySlice(w)
	res.closer = make([]write.WriteCloser, 0, len(w))
	for _, i := range w {
		if wc, ok := i.(write.WriteCloser); ok {
			res.closer = append(res.closer, wc)
		}
	}
	res.close = false

	return res
}
