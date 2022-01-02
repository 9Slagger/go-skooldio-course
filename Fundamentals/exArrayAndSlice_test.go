package fundamentals

import (
	"reflect"
	"testing"
)

func TestAbc(t *testing.T) {
	given := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	want := []string{"apple", "banana", "coconut"}

	get := abc(given)
	if equal := reflect.DeepEqual(want, get); !equal {
		t.Errorf("given %v want %v but %v", given, want, get)
	}
}

func TestEfg(t *testing.T) {
	given := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	want := []string{"elderberries", "figs", "grapes"}

	get := efg(given)
	if equal := reflect.DeepEqual(want, get); !equal {
		t.Errorf("given %v want %v but %v", given, want, get)
	}
}

func TestCde(t *testing.T) {
	given := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	want := []string{"coconut", "durian", "elderberries"}

	get := cde(given)
	if equal := reflect.DeepEqual(want, get); !equal {
		t.Errorf("given %v want %v but %v", given, want, get)
	}
}
