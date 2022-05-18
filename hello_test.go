// Copyright (c) 2022 Orlov Boris onlycodergod@gmail.com.
// This file hello_test.go is subject to the terms and
// conditions defined in file 'LICENSE', which is part of this project source code.
package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 4
	actual := add(2, 2)
	if actual != expected {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}
