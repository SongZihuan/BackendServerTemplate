// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package reutils

import "regexp"

const semVerRegexStr = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\+[0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*)?$`

var semVerRegex = regexp.MustCompile(semVerRegexStr)

// IsSemanticVersion checks if the given string is a valid semantic version.
func IsSemanticVersion(version string) bool {
	return semVerRegex.MatchString(version)
}
