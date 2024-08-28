package main

import (
	"testing"
)

func TestAdd2(t *testing.T) {
	tests := []struct {
		a, b, want int
		foo        string
	}{
		{a: 1, b: 2, want: 3},
		{3, 4, 7, "qwe"},
		{1, -1, 0, "asd"},
	}

	for _, test := range tests {
		got := Add2(test.a, test.b)
		if got != test.want {
			t.Errorf("Add(%d, %d) = %d but wanted %d", test.a, test.b, got, test.want)
		}
	}
}
