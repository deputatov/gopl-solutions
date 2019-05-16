package deduplicate

import (
	"reflect"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	testCases := []struct {
		input, want []string
	}{
		{[]string{"o", "o", "tw", "th", "th"}, []string{"o", "tw", "th"}},
		{[]string{"a", "a", "b", "b", "c"}, []string{"a", "b", "c"}},
	}
	for _, test := range testCases {
		output := deduplicate(test.input)
		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("deduplicate(%q) = %q", output, test.want)
		}
	}
}
