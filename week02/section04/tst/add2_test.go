package main

import "testing"

func TestAdd2(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 1, b: 2, want: 3},
		{3, -2, 1},
		{0, 6, 6},
	}

	for _, test := range tests {
		got := Add2(test.a, test.b)
		if got != test.want {
			t.Errorf("Add2(%d, %d) = %d but wanted %d\n", test.a, test.b, got, test.want)
		}
	}
}
