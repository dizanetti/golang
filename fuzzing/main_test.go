package main

import (
	"testing"
)

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if Compare(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func FuzzAdd(f *testing.F) {
	f.Fuzz(
		func(t *testing.T, a int, b int) {
			got := Add(a, b)
			want := a + b

			if Compare(got, want) {
				t.Errorf("got %q, wanted %q", got, want)
			}
		})
}

func Compare(got, want int) bool {
	return got != want
}
