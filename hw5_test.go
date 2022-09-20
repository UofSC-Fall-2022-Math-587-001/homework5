package hw5

import (
	"testing"
	"reflect"
)

func Test1(t *testing.T) {
	got := "hi"
	want := "bye"

	if !reflect.DeepEqual(got,want) {
		t.Errorf("%q vs %q", got, want)
	}
}
