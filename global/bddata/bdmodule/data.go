// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bdmodule

import (
	"time"
)

type GlobalData struct {
	LongVersion          string
	ShortVersion         string
	LongSemanticVersion  string
	ShortSemanticVersion string

	License   string
	Report    string
	BuildDate time.Time

	ModuleMame string
	CommitHash string

	ConfigSet BuildConfigSet
}
