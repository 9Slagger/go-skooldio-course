package fundamentals

import "testing"

func TestBook_String(t *testing.T) {
	given := Book{"Harry Potter", "J. K. Rowling"}
	want := "Harry Potter by J. K. Rowling"

	if get := given.String(); get != want {
		t.Errorf("given %v want %v but %v", given, want, get)
	}
}
