package main

import (
	"os"
	"testing"
)

func TestPassArgVerify(t *testing.T) {

	os.Args = []string{"cmd", "yeezy350"}

	actual := passArgVerify()

	if actual != true {
		t.Errorf("Test failed, expected: '%t', got:  '%t'", true, actual)
	}
}
