package todo

// ----------------------------------------------------------------------------
//  Type: TaskSegmentType
// ----------------------------------------------------------------------------

// TaskSegmentType represents type of segment in task string.
//
// The stringer implementation `String()` is defined in task_segment_type.go.
// See doc.go as well.
type TaskSegmentType uint8

// ----------------------------------------------------------------------------
//  Enums of TaskSegmentType
// ----------------------------------------------------------------------------

// Flags for indicating type of segment in task string.
const (
	SegmentIsCompleted TaskSegmentType = iota + 1
	SegmentCompletedDate
	SegmentPriority
	SegmentCreatedDate
	SegmentTodoText
	SegmentContext
	SegmentProject
	SegmentTag
	SegmentDueDate
)
