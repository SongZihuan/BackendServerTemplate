// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package resource

import "testing"

func TestVersion(t *testing.T) {
	t.Run("Version-Variable", func(t *testing.T) {
		if Version == "" {
			t.Fatalf("Version is empty")
		} else if SemanticVersioning == "" {
			t.Fatalf("SemanticVersioning is empty")
		} else if "v"+SemanticVersioning != Version {
			t.Fatalf("SemanticVersioning and Version do not match")
		} else if !utilsIsSemanticVersion(SemanticVersioning) {
			t.Fatalf("Non-semantic version")
		}
	})

	t.Run("Default-Version", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Fatalf("Get default version panic: %v", err)
			}
		}()
		_ = getDefaultVersion()
	})

	t.Run("Git-Version", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Fatalf("Get git version panic: %v", err)
			}
		}()
		_ = getGitTagVersion()
	})

	t.Run("Random-Version", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Fatalf("Get random version panic: %v", err)
			}
		}()

		ver := getRandomVersion()
		if ver == "" {
			t.Fatalf("Random Version is empty")
		}
	})
}

func TestLicense(t *testing.T) {
	if License == "" {
		t.Fatalf("License is empty")
	}
}

func TestReport(t *testing.T) {
	if Report == "" {
		t.Fatalf("Report is empty")
	}
}

func TestName(t *testing.T) {
	if Name == "" {
		t.Fatalf("Name is empty")
	}
}

func TestBuildDate(t *testing.T) {
	if buildDateTxt == "" {
		t.Fatalf("Build Date is empty")
	}
}

func TestRandomData(t *testing.T) {
	if randomData == "" {
		t.Fatalf("Randome Data is empty")
	}
}
