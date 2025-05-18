// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build !dev && !prod

package resource

import _ "embed"

//go:embed BUILD.yaml
var BuildConfig []byte
