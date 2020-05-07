package merge

import (
	//"fmt"
	"testing"

	"github.com/ippoippo/go-core-lib/equality/slice"
)

func TestTwoSortedIntArrays(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected []int
		desc     string // Description of testcase
	}{
		{[]int{1}, []int{2}, []int{1, 2}, "Length 1 and 1"},
		{[]int{1, 3, 5, 7}, []int{2, 4}, []int{1, 2, 3, 4, 5, 7}, "Length 4 and 2"},
		{[]int{1, 3, 5, 7}, []int{2, 4, 8, 9, 10, 11}, []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, "Length 4 and 6"},
	}

	for _, test := range tests {
		actual := TwoSortedIntArrays(test.a, test.b)
		if !slice.EqualInt(actual, test.expected) {
			t.Errorf("[%s]:String() = Actual:[%q], Expected:[%q]", test.desc, actual, test.expected)
		}
	}
}
