// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rtdata

import "time"

func GetName() string {
	if rtdata.Name != "" {
		return rtdata.Name
	}

	return "program"
}

func GetUTC() *time.Location {
	return rtdata.UTC
}

func GetLocalLocation() *time.Location {
	return rtdata.LocalLocation
}

func GetLocation() *time.Location {
	return rtdata.Location
}
