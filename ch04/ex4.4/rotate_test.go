package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		input, want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}},
		{[]int{2, 3, 4, 5, 1}, []int{3, 4, 5, 1, 2}},
	}
	for _, test := range testCases {
		rotate(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("rotate(%q) = %q", test.input, test.want)
		}
	}
}
