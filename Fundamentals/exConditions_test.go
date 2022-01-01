package fundamentals

import "testing"

func TestIsOdd(t *testing.T) {
	given := 2
	want := false

	get := isOdd(given)
	if want != get {
		t.Errorf("given %d want %t but %t", given, want, get)
	}
}
