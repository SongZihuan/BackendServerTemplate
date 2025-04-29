// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package reutils

import "testing"

func TestSemanticVersion(t *testing.T) {
	if !IsSemanticVersion("0.0.0") {
		t.Errorf("SemanticVersion test failed: 0.0.0 (must be true, but return false)")
	}

	if !IsSemanticVersion("1.0.0") {
		t.Errorf("SemanticVersion test failed: 1.0.0 (must be true, but return false)")
	}

	if !IsSemanticVersion("1.2.3") {
		t.Errorf("SemanticVersion test failed: 1.2.3 (must be true, but return false)")
	}

	if !IsSemanticVersion("1.0.0+dev") {
		t.Errorf("SemanticVersion test failed: 1.0.0+dev (must be true, but return false)")
	}

	if !IsSemanticVersion("1.0.0+dev-123") {
		t.Errorf("SemanticVersion test failed: 1.0.0+dev-123 (must be true, but return false)")
	}

	if !IsSemanticVersion("1.0.0+dev.123") {
		t.Errorf("SemanticVersion test failed: 1.0.0+dev-123 (must be true, but return false)")
	}

	if !IsSemanticVersion("1.0.0+dev-123.abc") {
		t.Errorf("SemanticVersion test failed: 1.0.0+dev-123-456 (must be true, but return false)")
	}

	if !IsSemanticVersion("1.0.0-123.456") {
		t.Errorf("SemanticVersion test failed: 1.0.0-123-456 (must be true, but return false)")
	}

	if !IsSemanticVersion("1.0.0-123-456+dev") {
		t.Errorf("SemanticVersion test failed: 1.0.0-123-456+dev (must be true, but return false)")
	}

	if !IsSemanticVersion("1.0.0-123-456+dev-127") {
		t.Errorf("SemanticVersion test failed: 1.0.0-123-456+dev-127 (must be true, but return false)")
	}

	if IsSemanticVersion("v0.0.0") {
		t.Errorf("SemanticVersion test failed: v0.0.0 (must be false, but return true)")
	}

	if IsSemanticVersion("1.0.0.0") {
		t.Errorf("SemanticVersion test failed: 1.0.0.0 (must be false, but return true)")
	}

	if IsSemanticVersion("1.0.0-123+dev-234+prod") {
		t.Errorf("SemanticVersion test failed: 1.0.0-123+dev-234+prod (must be false, but return true)")
	}
}
