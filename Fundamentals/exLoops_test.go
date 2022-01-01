package fundamentals

import "testing"

func TestIsPrime(t *testing.T) {
	given := 5
	want := true

	get := isPrime(given)
	if want != get {
		t.Errorf("given %d want %t but %t", given, want, get)
	}
}
