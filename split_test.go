package todotxt

import (
	"testing"
)

func isSameTaskSegmentList(s1, s2 []*TaskSegment) bool {
	return len(s1) == len(s2)
}

func TestTaskSegments(t *testing.T) {
	cases := []struct {
		text string
		segs []*TaskSegment
	}{
		{text: "2013-02-22 Pick up milk @GroceryStore",
			segs: []*TaskSegment{}},
	}

	for _, c := range cases {
		task, err := ParseTask(c.text)
		if err != nil {
			t.Errorf("Expected Task %q can be parsed, but got error: %v", c.text, err)
		}
		segs := task.Segments()
		if !isSameTaskSegmentList(segs, c.segs) {
			t.Errorf("Expected Task %q to be [%v], but got error: %v", c.text, c.segs, segs)
		}
	}
}
