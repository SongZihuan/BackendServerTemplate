// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cleanstringutils

import "testing"

func TestCheckAndRemoveBOM(t *testing.T) {
	hasBOM := string([]byte{0xEF, 0xBB, 0xBF}) + "Hello"
	noBOM := "Hello"

	if CheckAndRemoveBOM(noBOM) != noBOM {
		t.Errorf("No BOM check error")
	}

	if CheckAndRemoveBOM(hasBOM) == hasBOM {
		t.Errorf("Has BOM check error")
	}

	if CheckAndRemoveBOM(hasBOM) != noBOM {
		t.Errorf("Has BOM remove error")
	}
}
