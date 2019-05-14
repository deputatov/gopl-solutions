package reverse

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		input, want [3]int
	}{
		{[3]int{1, 2, 3}, [3]int{3, 2, 1}},
	}
	for _, test := range testCases {
		reverse(&test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("reverse(%q) = %q", test.input, test.want)
		}
	}
}
