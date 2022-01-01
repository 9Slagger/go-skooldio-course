package fundamentals

import "testing"

func TestDouble(t *testing.T) {
	given := 2
	want := 4
	get := given

	double(&get)
	if want != get {
		t.Errorf("given %d want %d but %d", given, want, get)
	}
}

func TestAppendGreeting(t *testing.T) {
	given := "Bob"
	want := "Hi, Bob"
	get := given

	appendGreeting(&get)
	if want != get {
		t.Errorf("given %s want %s but %s", given, want, get)
	}
}
