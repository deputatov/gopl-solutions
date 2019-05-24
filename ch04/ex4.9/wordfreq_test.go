package wordfreq

import (
	"reflect"
	"testing"
)

func TestWordfreq(t *testing.T) {
	input := wordfreq()
	want := map[string]int{
		"word":  2,
		"word1": 1,
		"word2": 2,
	}
	if !reflect.DeepEqual(input, want) {
		t.Errorf("wordfreq(%v) = %v", input, want)
	}
}
