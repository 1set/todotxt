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

		{text: "x Download Todo.txt mobile app @Phone",
			segs: []*TaskSegment{
				{
					Type:      TaskSegment_IsCompleted,
					Originals: []string{"x"},
					Display:   "x",
				},
				{
					Type:      TaskSegment_TodoText,
					Originals: []string{"Download Todo.txt mobile app"},
					Display:   "Download Todo.txt mobile app",
				},
				{
					Type:      TaskSegment_Context,
					Originals: []string{"Phone"},
					Display:   "@Phone",
				},
			}},

		{text: "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
			segs: []*TaskSegment{
				{
					Type:      TaskSegment_Priority,
					Originals: []string{"B"},
					Display:   "(B)",
				},
				{
					Type:      TaskSegment_CreatedDate,
					Originals: []string{"2013-12-01"},
					Display:   "2013-12-01",
				},
				{
					Type:      TaskSegment_TodoText,
					Originals: []string{"Outline chapter 5"},
					Display:   "Outline chapter 5",
				},
				{
					Type:      TaskSegment_Context,
					Originals: []string{"Computer"},
					Display:   "@Computer",
				},
				{
					Type:      TaskSegment_Project,
					Originals: []string{"Novel"},
					Display:   "+Novel",
				},
				{
					Type:      TaskSegment_Tag,
					Originals: []string{"Level", "5"},
					Display:   "Level:5",
				},
				{
					Type:      TaskSegment_Tag,
					Originals: []string{"private", "false"},
					Display:   "private:false",
				},
				{
					Type:      TaskSegment_DueDate,
					Originals: []string{"due:2014-02-17"},
					Display:   "due:2014-02-17",
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
