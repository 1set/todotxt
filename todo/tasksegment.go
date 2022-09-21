package todo

// ----------------------------------------------------------------------------
//  Type: TaskSegment
// ----------------------------------------------------------------------------

// TaskSegment represents a segment in task string.
type TaskSegment struct {
	Display   string
	Originals []string
	Type      TaskSegmentType
}
