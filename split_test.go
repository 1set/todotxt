package todotxt

import (
	"fmt"
	"strings"
	"testing"
)

func isSameTaskSegmentList(s1, s2 []*TaskSegment) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		a, b := s1[i], s2[i]
		if a.Type != b.Type {
			return false
		}
		if a.Display != b.Display {
			return false
		}
		if len(a.Originals) != len(b.Originals) {
			return false
		}
		for j := 0; j < len(a.Originals); j++ {
			if a.Originals[j] != b.Originals[j] {
				return false
			}
		}
	}
	return true
}

func strTaskSegmentList(l []*TaskSegment) string {
	var parts []string
	for _, s := range l {
		parts = append(parts, fmt.Sprintf("%v", *s))
	}
	return strings.Join(parts, ", ")
}

func TestTaskSegments(t *testing.T) {
	cases := []struct {
		text string
		segs []*TaskSegment
	}{
		{text: "2013-02-22 Pick up milk @GroceryStore",
			segs: []*TaskSegment{
				{
					Type:      TaskSegment_CreatedDate,
					Originals: []string{"2013-02-22"},
					Display:   "2013-02-22",
				},
				{
					Type:      TaskSegment_TodoText,
					Originals: []string{"Pick up milk"},
					Display:   "Pick up milk",
				},
				{
					Type:      TaskSegment_Context,
					Originals: []string{"GroceryStore"},
					Display:   "@GroceryStore",
				},
			}},
	}

	for _, c := range cases {
		task, err := ParseTask(c.text)
		if err != nil {
			t.Errorf("Expected Task %q can be parsed, but got error: %v", c.text, err)
		}
		segs := task.Segments()
		if !isSameTaskSegmentList(segs, c.segs) {
			t.Errorf("Expected segments to be [%s], but got [%s]", strTaskSegmentList(c.segs), strTaskSegmentList(segs))
		}
	}
}
