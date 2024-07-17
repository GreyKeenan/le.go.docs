
package myLibMod

import (
	"testing"
)

func Test_Call(t *testing.T) {
	var _ string
	var err error
	_, err = Call("word")
	if (err != nil) {
		t.Fatalf("Calling Call() with a standard word resulted in an error: %v", err)
	}
}

func Test_CallEmpty(t *testing.T) {
	var message string
	var err error

	message, err = Call("")
	if (message != "" || err == nil) {
		t.Fatalf("Call(\"\") == %v, %v. Expected \"\", !nil.", message, err)
	}
}
