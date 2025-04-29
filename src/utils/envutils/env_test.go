// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package envutils

import (
	"os"
	"testing"
)

func TestToEnvName(t *testing.T) {
	if res := ToEnvName("AC"); res != "AC" {
		t.Errorf("to env name error: AC -> AC: %s", res)
	}

	if res := ToEnvName("abc"); res != "ABC" {
		t.Errorf("to env name error: abc -> ABC: %s", res)
	}

	if res := ToEnvName("A.C"); res != "A_C" {
		t.Errorf("to env name error: A.C -> A_C: %s", res)
	}

	if res := ToEnvName("a.c"); res != "A_C" {
		t.Errorf("to env name error: a.c -> A_C: %s", res)
	}
}

func TestGetSysEnv(t *testing.T) {
	var err error

	err = os.Setenv("TEST_A", "1")
	if err != nil {
		t.Fatalf("set env TEST_A error: %s", err.Error())
	}

	err = os.Setenv("test_b", "2")
	if err != nil {
		t.Fatalf("set env TEST_A error: %s", err.Error())
	}

	if res := GetSysEnv("TEST_A"); res != "1" {
		t.Errorf("get sys env error: TEST_A -> 1: %s", res)
	}

	if res := GetSysEnv("test_b"); res != "2" {
		t.Errorf("get sys env error: test_a -> 2: %s", res)
	}
}

func TestGetEnv(t *testing.T) {
	var err error

	err = os.Setenv("P_R_TEST_C", "3")
	if err != nil {
		t.Fatalf("set env TEST_A error: %s", err.Error())
	}

	if res := GetEnv("P_R", "TEST_C"); res != "3" {
		t.Errorf("get env error: P_R TEST_C -> 1: %s", res)
	}

	if res := GetEnv("P_R", "test.c"); res != "3" {
		t.Errorf("get env error: P_R TEST_C -> 1: %s", res)
	}
}
