package tdd

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{1, "1"},
	{3, "Fizz"},
	{5, "Buzz"},
	{12, "Fizz"},
	{15, "FizzBuzz"},
}

func TestFizzbuzz(t *testing.T) {
	for i, test := range tests {
		v := FizzBuzz(test.in)
		if v != test.out {
			t.Error("%d: Fizzbuzz(%d)=%s; want %s", i, test.in, v, test.out)
		}
	}
}
