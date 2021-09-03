package main

import (
	"fmt"
	"testing"
)

// TestSprintf START OMIT
var fmtTests = []struct {
	fmt string
	val interface{}
	out string
}{
	{"%d", 12345, "12345"},
	{"%v", 12345, "12345"},
	{"%t", true, "true"},
}

func TestSprintf(t *testing.T) {
	for _, tt := range fmtTests {
		if s := fmt.Sprintf(tt.fmt, tt.val); s != tt.out {
			t.Errorf("...")
		}
	}
}

// TestSprintf END OMIT

func main() {
	fmt.Print(`% go test -v
=== RUN   TestSprintf
--- PASS: TestSprintf (0.00s)
PASS
ok  	gotalks/fmt	0.165s`)
}
