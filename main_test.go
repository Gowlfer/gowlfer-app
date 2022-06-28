package main_test

import "testing"

func TestTrue(t *testing.T) {
	if !true {
		t.Error("true is not true")
	}
}
