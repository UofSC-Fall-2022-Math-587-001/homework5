package hw5

import (
	"testing"
	"github.com/UofSC-Fall-2022-Math-587-001/homework5/library"
)

func Test1(t *testing.T) {
	val := library.FastPower(11,2,3)
	got := babyGiantStep(11,2,val)
	want := 3

	if got != want {
		t.Errorf("%d vs %d", got, want)
	}
}

func Test2(t *testing.T) {
	val := library.FastPower(13,2,8)
	got := babyGiantStep(13,2,val)
	want := 8

	if got != want {
		t.Errorf("%d vs %d", got, want)
	}
}

func Test3(t *testing.T) {
	val := library.FastPower(23,15,8)
	got := babyGiantStep(23,15,val)
	want := 8

	if got != want {
		t.Errorf("%d vs %d", got, want)
	}
}

func Test4(t *testing.T) {
	val := library.FastPower(29,27,2)
	got := babyGiantStep(29,27,val)
	want := 2

	if got != want {
		t.Errorf("%d vs %d", got, want)
	}
}

func Test5(t *testing.T) {
	val := library.FastPower(2,1,0)
	got := babyGiantStep(2,1,val)
	want := 0

	if got != want {
		t.Errorf("%d vs %d", got, want)
	}
}
