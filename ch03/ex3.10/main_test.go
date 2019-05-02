package main

import "testing"

func TestCommonIterative(t *testing.T) {
	testCases := []struct {
		input, want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},
	}
	for _, test := range testCases {
		if got := commaIterative(test.input); got != test.want {
			t.Errorf("commaIterative(%q) = %q", test.input, test.want)
		}
	}
}
