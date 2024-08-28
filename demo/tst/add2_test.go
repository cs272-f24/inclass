package main

import "testing"

func TestAdd2(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 1, b: 2, want: 3},
		{3, 4, 7},
		{5, 6, 11},
	}

	for _, test := range tests {
		got := Add2(test.a, test.b)
		if got != test.want {
			t.Errorf("Add2 returned %d but expected %d\n", got, test.want)
		}
	}
}
