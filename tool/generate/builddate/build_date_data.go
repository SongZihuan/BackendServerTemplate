// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package builddate

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/tool/generate/basefile"
	"github.com/SongZihuan/BackendServerTemplate/utils/fileutils"
	"log"
	"time"
)

var buildTime = time.Now()

func WriteBuildDateData() error {
	val := fmt.Sprint(buildTime.Unix())

	log.Printf("generate: write build data: %s UTC (timestaamp: %s)\n", buildTime.UTC().Format(time.DateTime), val)
	defer log.Println("generate: write build data finish")

	log.Printf("generate: write %s to file %s\n", val, basefile.FileBuildDateTxt)
	return fileutils.Write(basefile.FileBuildDateTxt, val)
}
