// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package reutils

import "regexp"

const (
	semVerRegexpStr      = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\+[0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*)?$`
	shortSemVerRegexpStr = `^((0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*))`
)

var semVerRegexp = regexp.MustCompile(semVerRegexpStr)
var shortSemVerRegexp = regexp.MustCompile(shortSemVerRegexpStr)

// IsSemanticVersion checks if the given string is a valid semantic version.
func IsSemanticVersion(version string) bool {
	return semVerRegexp.MatchString(version)
}

func GetShortSemanticVersion(version string) string {
	return shortSemVerRegexp.FindString(version)
}

func CheckAndGetShortSemanticVersion(version string) string {
	if !IsSemanticVersion(version) {
		return ""
	}
	return GetShortSemanticVersion(version)
}
