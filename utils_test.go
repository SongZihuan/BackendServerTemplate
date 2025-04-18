// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package resource

import "testing"

func TestSemanticVersion(t *testing.T) {
	if !utilsIsSemanticVersion("0.0.0") {
		t.Errorf("SemanticVersion test failed: 0.0.0 (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.0.0") {
		t.Errorf("SemanticVersion test failed: 1.0.0 (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.2.3") {
		t.Errorf("SemanticVersion test failed: 1.2.3 (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.0.0+dev") {
		t.Errorf("SemanticVersion test failed: 1.0.0+dev (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.0.0+dev-123") {
		t.Errorf("SemanticVersion test failed: 1.0.0+dev-123 (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.0.0+dev.123") {
		t.Errorf("SemanticVersion test failed: 1.0.0+dev-123 (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.0.0+dev-123.abc") {
		t.Errorf("SemanticVersion test failed: 1.0.0+dev-123-456 (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.0.0-123.456") {
		t.Errorf("SemanticVersion test failed: 1.0.0-123-456 (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.0.0-123-456+dev") {
		t.Errorf("SemanticVersion test failed: 1.0.0-123-456+dev (must be true, but return false)")
	}

	if !utilsIsSemanticVersion("1.0.0-123-456+dev-127") {
		t.Errorf("SemanticVersion test failed: 1.0.0-123-456+dev-127 (must be true, but return false)")
	}

	if utilsIsSemanticVersion("v0.0.0") {
		t.Errorf("SemanticVersion test failed: v0.0.0 (must be false, but return true)")
	}

	if utilsIsSemanticVersion("1.0.0.0") {
		t.Errorf("SemanticVersion test failed: 1.0.0.0 (must be false, but return true)")
	}

	if utilsIsSemanticVersion("1.0.0-123+dev-234+prod") {
		t.Errorf("SemanticVersion test failed: 1.0.0-123+dev-234+prod (must be false, but return true)")
	}
}

func TestCheckAndRemoveBOM(t *testing.T) {
	hasBOM := string([]byte{0xEF, 0xBB, 0xBF}) + "Hello"
	noBOM := "Hello"

	if utilsCheckAndRemoveBOM(noBOM) != noBOM {
		t.Errorf("No BOM check error")
	}

	if utilsCheckAndRemoveBOM(hasBOM) == hasBOM {
		t.Errorf("Has BOM check error")
	}

	if utilsCheckAndRemoveBOM(hasBOM) != noBOM {
		t.Errorf("Has BOM remove error")
	}
}

func TestClenFileData(t *testing.T) {
	t.Run("Text-OnlyLine", func(t *testing.T) {
		text := "Hello"
		if utilsClenFileData(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-OnlyLine-WithSpace", func(t *testing.T) {
		text := "Hello    "
		if utilsClenFileData(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-CRLF", func(t *testing.T) {
		text := "Hello\r\n"
		if utilsClenFileData(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-More-CRLF", func(t *testing.T) {
		text := "Hello\r\n\r\n\r\n"
		if utilsClenFileData(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-CRLF-WithSpace", func(t *testing.T) {
		text := "Hello    \r\n"
		if utilsClenFileData(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-MoreLine", func(t *testing.T) {
		text := "Hello\r\nWorld"
		if utilsClenFileData(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})
}

func TestClenFileDataMoreLine(t *testing.T) {
	t.Run("Text-OnlyLine", func(t *testing.T) {
		text := "Hello"
		if utilsClenFileDataMoreLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-OnlyLine-WithSpace", func(t *testing.T) {
		text := " Hello    "
		if utilsClenFileDataMoreLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-CRLF", func(t *testing.T) {
		text := "Hello\r\n"
		if utilsClenFileDataMoreLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-More-CRLF", func(t *testing.T) {
		text := "Hello\r\n\r\n\r\n"
		if utilsClenFileDataMoreLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-With-CRLF-WithSpace", func(t *testing.T) {
		text := "Hello    \r\n"
		if utilsClenFileDataMoreLine(text) != "Hello" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})

	t.Run("Text-MoreLine", func(t *testing.T) {
		text := "Hello\r\nWorld"
		if utilsClenFileDataMoreLine(text) != "Hello\nWorld" {
			t.Errorf("ClenFileData OnlyLine error")
		}
	})
}
