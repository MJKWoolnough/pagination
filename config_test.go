package pagination

import (
	"reflect"
	"testing"
)

func newSections(n ...uint) [3]section {
	var sections [3]section
	l := len(n)
	if l > 1 {
		sections[0].First = n[0]
		sections[0].Last = n[1]
	}
	if l > 3 {
		sections[1].First = n[2]
		sections[1].Last = n[3]
	}
	if l > 5 {
		sections[2].First = n[4]
		sections[2].Last = n[5]
	}
	return sections
}

func TestConfigGet(t *testing.T) {

	noEnds := Config{0, 3}
	noMiddle := Config{3, 0}
	normal := Config{3, 3}

	tests := []struct {
		Config
		currPage, lastPage uint
		sections           [3]section
	}{
		{normal, 0, 9, newSections(0, 3, 7, 9)},
		{normal, 1, 9, newSections(0, 4, 7, 9)},
		{normal, 2, 9, newSections(0, 9)},
		{normal, 3, 9, newSections(0, 9)},
		{normal, 4, 9, newSections(0, 9)},
		{normal, 5, 9, newSections(0, 9)},
		{normal, 6, 9, newSections(0, 9)},
		{normal, 7, 9, newSections(0, 9)},
		{normal, 8, 9, newSections(0, 2, 5, 9)},
		{normal, 9, 9, newSections(0, 2, 6, 9)},
		{normal, 50, 100, newSections(0, 2, 47, 53, 98, 100)},
		{noMiddle, 0, 9, newSections(0, 2, 7, 9)},
		{noMiddle, 1, 9, newSections(0, 2, 7, 9)},
		{noMiddle, 2, 9, newSections(0, 2, 7, 9)},
		{noMiddle, 3, 9, newSections(0, 3, 7, 9)},
		{noMiddle, 4, 9, newSections(0, 4, 7, 9)},
		{noMiddle, 5, 9, newSections(0, 2, 5, 9)},
		{noMiddle, 6, 9, newSections(0, 2, 6, 9)},
		{noMiddle, 7, 9, newSections(0, 2, 7, 9)},
		{noMiddle, 8, 9, newSections(0, 2, 7, 9)},
		{noMiddle, 9, 9, newSections(0, 2, 7, 9)},
		{noEnds, 0, 9, newSections(0, 3)},
		{noEnds, 1, 9, newSections(0, 4)},
		{noEnds, 2, 9, newSections(0, 5)},
		{noEnds, 3, 9, newSections(0, 6)},
		{noEnds, 4, 9, newSections(1, 7)},
		{noEnds, 5, 9, newSections(2, 8)},
		{noEnds, 6, 9, newSections(3, 9)},
		{noEnds, 7, 9, newSections(4, 9)},
		{noEnds, 8, 9, newSections(5, 9)},
		{noEnds, 9, 9, newSections(6, 9)},
	}

	for n, test := range tests {
		if s := test.Get(test.currPage, test.lastPage); !reflect.DeepEqual(s.sections, test.sections) {
			t.Errorf("test %d: expecting %v, got %v", n+1, test.sections, s.sections)
		}
	}
}
