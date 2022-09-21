package todo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:paralleltest // do not parallel to avoid race conditions
func TestTaskSegmentType(t *testing.T) {
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

	for name, expect := range names {
		actual := name.String()

		require.Equal(t, expect, actual,
			"the TaskSegmentType(%d).String() did not return the expected value", name)
	}
}
