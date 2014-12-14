package pagination

import "testing"

func TestItoa(t *testing.T) {
	tests := []struct {
		num uint
		str string
	}{
		{0, "0"},
		{1, "1"},
		{2, "2"},
		{9, "9"},
		{10, "10"},
		{11, "11"},
		{20, "20"},
		{99, "99"},
		{100, "100"},
		{101, "101"},
		{999999999, "999999999"},
		{3999999999, "3999999999"},
		{4000000000, "4000000000"},
		{4294967295, "4294967295"},
	}

	for n, test := range tests {
		if numStr := itoa(test.num); numStr != test.str {
			t.Errorf("test %d: expecting %q, got %q", n+1, test.str, numStr)
		}
	}
}
