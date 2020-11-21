package todotxt

import (
	"testing"
)

func BenchmarkTask_Segments(b *testing.B) {
	s := "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt test:benchmark due:2014-01-12   "
	task, _ := ParseTask(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = task.Segments()
	}
}

func TestTaskTaskSegmentType(t *testing.T) {
	names := map[TaskSegmentType]string{
		SegmentIsCompleted:   "IsCompleted",
		SegmentCompletedDate: "CompletedDate",
		SegmentPriority:      "Priority",
		SegmentCreatedDate:   "CreatedDate",
		SegmentTodoText:      "TodoText",
		SegmentContext:       "Context",
		SegmentProject:       "Project",
		SegmentTag:           "Tag",
		SegmentDueDate:       "DueDate",
		0:                    "TaskSegmentType(0)",
	}
	for n, s := range names {
		if ss := n.String(); ss != s {
			t.Errorf("Expected Segment %v is %q, but got: %q", n, s, ss)
		}
	}
}

func TestTaskSegments(t *testing.T) {
	cases := []struct {
		text string
		segs []*TaskSegment
	}{
		{text: "2013-02-22 Pick up milk @GroceryStore",
			segs: []*TaskSegment{
				{
					Type:      SegmentCreatedDate,
					Originals: []string{"2013-02-22"},
					Display:   "2013-02-22",
				},
				{
					Type:      SegmentTodoText,
					Originals: []string{"Pick up milk"},
					Display:   "Pick up milk",
				},
				{
					Type:      SegmentContext,
					Originals: []string{"GroceryStore"},
					Display:   "@GroceryStore",
				},
			}},
		{text: "x Download Todo.txt mobile app @Phone",
			segs: []*TaskSegment{
				{
					Type:      SegmentIsCompleted,
					Originals: []string{"x"},
					Display:   "x",
				},
				{
					Type:      SegmentTodoText,
					Originals: []string{"Download Todo.txt mobile app"},
					Display:   "Download Todo.txt mobile app",
				},
				{
					Type:      SegmentContext,
					Originals: []string{"Phone"},
					Display:   "@Phone",
				},
			}},
		{text: "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
			segs: []*TaskSegment{
				{
					Type:      SegmentPriority,
					Originals: []string{"B"},
					Display:   "(B)",
				},
				{
					Type:      SegmentCreatedDate,
					Originals: []string{"2013-12-01"},
					Display:   "2013-12-01",
				},
				{
					Type:      SegmentTodoText,
					Originals: []string{"Outline chapter 5"},
					Display:   "Outline chapter 5",
				},
				{
					Type:      SegmentContext,
					Originals: []string{"Computer"},
					Display:   "@Computer",
				},
				{
					Type:      SegmentProject,
					Originals: []string{"Novel"},
					Display:   "+Novel",
				},
				{
					Type:      SegmentTag,
					Originals: []string{"Level", "5"},
					Display:   "Level:5",
				},
				{
					Type:      SegmentTag,
					Originals: []string{"private", "false"},
					Display:   "private:false",
				},
				{
					Type:      SegmentDueDate,
					Originals: []string{"due:2014-02-17"},
					Display:   "due:2014-02-17",
				},
			}},
		{text: "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
			segs: []*TaskSegment{
				{
					Type:      SegmentIsCompleted,
					Originals: []string{"x"},
					Display:   "x",
				},
				{
					Type:      SegmentCompletedDate,
					Originals: []string{"2014-01-02"},
					Display:   "2014-01-02",
				},
				{
					Type:      SegmentPriority,
					Originals: []string{"B"},
					Display:   "(B)",
				},
				{
					Type:      SegmentCreatedDate,
					Originals: []string{"2013-12-30"},
					Display:   "2013-12-30",
				},
				{
					Type:      SegmentTodoText,
					Originals: []string{"Create golang library test cases"},
					Display:   "Create golang library test cases",
				},
				{
					Type:      SegmentContext,
					Originals: []string{"Go"},
					Display:   "@Go",
				},
				{
					Type:      SegmentProject,
					Originals: []string{"go-todotxt"},
					Display:   "+go-todotxt",
				},
			}},
		{text: "x 2014-01-03 2014-01-01 Create some more golang library test cases @Go +go-todotxt",
			segs: []*TaskSegment{
				{
					Type:      SegmentIsCompleted,
					Originals: []string{"x"},
					Display:   "x",
				},
				{
					Type:      SegmentCompletedDate,
					Originals: []string{"2014-01-03"},
					Display:   "2014-01-03",
				},
				{
					Type:      SegmentCreatedDate,
					Originals: []string{"2014-01-01"},
					Display:   "2014-01-01",
				},
				{
					Type:      SegmentTodoText,
					Originals: []string{"Create some more golang library test cases"},
					Display:   "Create some more golang library test cases",
				},
				{
					Type:      SegmentContext,
					Originals: []string{"Go"},
					Display:   "@Go",
				},
				{
					Type:      SegmentProject,
					Originals: []string{"go-todotxt"},
					Display:   "+go-todotxt",
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
