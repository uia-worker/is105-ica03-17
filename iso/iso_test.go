package iso

import (
	"testing"
)

func TestGretingsExtendedASCII(t *testing.T) {
	ASCII := GreetingExtendedASCII2()
	if !(isASCII(ASCII)) {
		t.Fail()
	}
}

func isASCII(s string) bool {
	for _, C := range s {
		if C > 127 && C < 256 {
			return true
		}
	}
	return false
}
