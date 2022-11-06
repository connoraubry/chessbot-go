package main

import "testing"

func TestIsNumberEvent(t *testing.T) {

	var testVal int64 = 1
	msg := isNumberEven(testVal)
	if msg == true {
		t.Fatalf(`isNumberEven(1) = %v`, msg)
	}

}
