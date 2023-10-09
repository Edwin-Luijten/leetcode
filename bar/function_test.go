package challenge_pkg

import (
	"reflect"
	"testing"
)

type TestCase struct {
	expected int
}

var tests []TestCase = []TestCase{
	{expected: 0},
}

var benchmark = TestCase{expected: 0}

func TestChallenge(t *testing.T) {
	for _, test := range tests {
		actual := challenge()

		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("challenge() = %v; expected %v", actual, test.expected)
		}
	}
}

func BenchmarkChallenge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		challenge()
	}
}
