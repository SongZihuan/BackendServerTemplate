// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package copyutils

import (
	"bytes"
	"encoding/gob"
)

func DeepCopy[T any](src *T) (*T, error) {
	if src == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	var dest T

	err := gob.NewEncoder(&buf).Encode(src)
	if err != nil {
		return nil, err
	}

	err = gob.NewDecoder(&buf).Decode(&dest)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}
